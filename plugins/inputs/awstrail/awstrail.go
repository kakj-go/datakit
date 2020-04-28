package awstrail

import (
	"context"
	"encoding/json"
	"time"

	"github.com/influxdata/telegraf"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/time/rate"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/models"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

type (
	AwsTrailAgent struct {
		AwsTrailInstance []*AwsTrailInstance `toml:"aws_cloudtrail"`

		ctx       context.Context
		cancelFun context.CancelFunc

		accumulator telegraf.Accumulator

		logger *models.Logger

		runningInstances []*runningInstance
	}

	runningInstance struct {
		cfg *AwsTrailInstance

		apiClient *cloudtrail.CloudTrail

		agent *AwsTrailAgent

		logger *models.Logger

		metricName string

		rateLimiter *rate.Limiter

		lastTime time.Time
	}
)

func (_ *AwsTrailAgent) SampleConfig() string {
	return sampleConfig
}

func (_ *AwsTrailAgent) Description() string {
	return ""
}

func (_ *AwsTrailAgent) Gather(telegraf.Accumulator) error {
	return nil
}

func (_ *AwsTrailAgent) Init() error {
	return nil
}

func (a *AwsTrailAgent) Start(acc telegraf.Accumulator) error {
	a.logger = &models.Logger{
		Name: inputName,
	}

	if len(a.AwsTrailInstance) == 0 {
		a.logger.Warnf("no configuration found")
		return nil
	}

	a.logger.Infof("starting...")

	a.accumulator = acc

	for _, instCfg := range a.AwsTrailInstance {
		r := &runningInstance{
			cfg:    instCfg,
			agent:  a,
			logger: a.logger,
		}
		r.metricName = instCfg.MetricName
		if r.metricName == "" {
			r.metricName = `aws_cloudtrail`
		}

		if r.cfg.Interval.Duration == 0 {
			r.cfg.Interval.Duration = time.Minute
		}

		limit := rate.Every(600 * time.Millisecond)
		r.rateLimiter = rate.NewLimiter(limit, 1)

		a.runningInstances = append(a.runningInstances, r)

		go r.run(a.ctx)
	}

	return nil
}

func (r *runningInstance) fetchOnce(ctx context.Context, params *cloudtrail.LookupEventsInput, token string) (*cloudtrail.LookupEventsOutput, error) {

	reqUid, _ := uuid.NewV4()

	var tempDelay time.Duration
	var response *cloudtrail.LookupEventsOutput
	var err error
	for i := 0; i < 5; i++ {
		r.rateLimiter.Wait(ctx)
		response, err = r.apiClient.LookupEvents(params)

		if tempDelay == 0 {
			tempDelay = time.Millisecond * 50
		} else {
			tempDelay *= 2
		}

		if max := time.Second; tempDelay > max {
			tempDelay = max
		}

		if err != nil {
			r.logger.Warnf("%s", err)
			time.Sleep(tempDelay)
		} else {
			if i != 0 {
				r.logger.Debugf("retry %s successed, %d", reqUid.String(), i)
			}
			break
		}
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *runningInstance) run(ctx context.Context) error {

	defer func() {
		if e := recover(); e != nil {
			r.logger.Errorf("panic, %v", e)
		}
	}()

	type evResSt struct {
		ResourceName string
		ResourceType string
	}

	if err := r.initClient(); err != nil {
		r.logger.Errorf("fail to init client, %s", err)
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return context.Canceled
		default:
		}

		useage := time.Now()

		end_time := time.Now().UTC().Truncate(time.Minute)
		var start_time time.Time
		if r.lastTime.IsZero() {
			start_time = end_time.Add(-r.cfg.Interval.Duration)
		} else {
			start_time = r.lastTime.Add(10 * time.Second)
		}

		var response []*cloudtrail.LookupEventsOutput
		var tempResp *cloudtrail.LookupEventsOutput
		var err error
		token := ""

		params := &cloudtrail.LookupEventsInput{
			EndTime:    aws.Time(end_time),
			StartTime:  aws.Time(start_time),
			MaxResults: aws.Int64(500),
		}

		for {
			params.NextToken = aws.String(token)

			tempResp, err = r.fetchOnce(ctx, params, token)
			if err != nil {
				break
			}
			response = append(response, tempResp)
			if tempResp.NextToken == nil {
				break
			} else {
				token = *tempResp.NextToken
			}
		}

		if err != nil {
			r.logger.Errorf("fail to LookupEvents, %s", err)
		}

		r.lastTime = end_time

		used := time.Now().Sub(useage)

		for _, res := range response {
			for _, ev := range res.Events {
				tags := map[string]string{}
				tags["EventName"] = *ev.EventName
				tags["ReadOnly"] = *ev.ReadOnly
				tags["Username"] = *ev.Username

				fields := map[string]interface{}{}
				fields["EventId"] = *ev.EventId
				fields["AccessKeyId"] = *ev.AccessKeyId
				fields["EventSource"] = *ev.EventSource
				fields["Detail"] = *ev.CloudTrailEvent
				refRes := []*evResSt{}
				for _, r := range ev.Resources {
					refRes = append(refRes, &evResSt{
						ResourceName: *r.ResourceName,
						ResourceType: *r.ResourceType,
					})
				}
				if resStr, err := json.Marshal(refRes); err == nil {
					fields["Resource"] = resStr
				}

				r.agent.accumulator.AddFields(r.metricName, fields, tags, *ev.EventTime)
			}
		}

		if used < r.cfg.Interval.Duration {
			internal.SleepContext(ctx, r.cfg.Interval.Duration-used)
		}
	}
}

func (r *runningInstance) initClient() error {
	cred := credentials.NewStaticCredentials(r.cfg.AccessKey, r.cfg.AccessSecret, r.cfg.AccessToken)

	cfg := aws.NewConfig()
	cfg.WithCredentials(cred).WithRegion(r.cfg.RegionID) //.WithRegion(`cn-north-1`)

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigDisable,
		Config:            *cfg,
	})

	if err != nil {
		return err
	}

	r.apiClient = cloudtrail.New(sess)

	return nil
}

func (a *AwsTrailAgent) Stop() {
	a.cancelFun()
}

func init() {
	inputs.Add(inputName, func() telegraf.Input {
		ac := &AwsTrailAgent{}
		ac.ctx, ac.cancelFun = context.WithCancel(context.Background())
		return ac
	})
}