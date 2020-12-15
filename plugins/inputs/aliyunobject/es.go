package aliyunobject

import (
	"encoding/json"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/elasticsearch"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
)

const (
	elasticsearchSampleConfig = `
#[inputs.aliyunobject.elasticsearch]

# ## @param - [list of elasticsearch instanceid] - optional
#instanceids = []

# ## @param - [list of excluded elasticsearch instanceid] - optional
#exclude_instanceids = []

# ## @param - custom tags for ecs object - [list of key:value element] - optional
#[inputs.aliyunobject.elasticsearch.tags]
# key1 = 'val1'
`
)

type Elasticsearch struct {
	Tags               map[string]string `toml:"tags,omitempty"`
	InstancesIDs       []string          `toml:"instanceids,omitempty"`
	ExcludeInstanceIDs []string          `toml:"exclude_instanceids,omitempty"`
}

func (e *Elasticsearch) run(ag *objectAgent) {
	var cli *elasticsearch.Client
	var err error

	for {
		select {
		case <-ag.ctx.Done():
			return
		default:
		}

		cli, err = elasticsearch.NewClientWithAccessKey(ag.RegionID, ag.AccessKeyID, ag.AccessKeySecret)
		if err == nil {
			break
		}
		moduleLogger.Errorf("%s", err)
		datakit.SleepContext(ag.ctx, time.Second*3)
	}

	for {

		select {
		case <-ag.ctx.Done():
			return
		default:
		}

		page := 1
		size := 100
		req := elasticsearch.CreateListInstanceRequest()
		for {
			moduleLogger.Infof("pageNume %v, pagesize %v", page, size)
			if len(e.InstancesIDs) > 0 {
				if page <= len(e.InstancesIDs) {
					req.InstanceId = e.InstancesIDs[page-1]
				} else {
					break
				}
			} else {
				req.Page = requests.NewInteger(page)
				req.Size = requests.NewInteger(size)
			}
			resp, err := cli.ListInstance(req)

			select {
			case <-ag.ctx.Done():
				return
			default:
			}

			if err == nil {
				e.handleResponse(resp, ag)
			} else {
				moduleLogger.Errorf("%s", err)
				if len(e.InstancesIDs) > 0 {
					page++
					continue
				}
				break
			}

			if len(e.InstancesIDs) <= 0 && resp.Headers.XTotalCount < page*size {
				break
			}

			page++
			if len(e.InstancesIDs) <= 0 {
				req.Page = requests.NewInteger(page)
			}
		}

		datakit.SleepContext(ag.ctx, ag.Interval.Duration)
	}
}

func (e *Elasticsearch) handleResponse(resp *elasticsearch.ListInstanceResponse, ag *objectAgent) {
	var objs []map[string]interface{}

	for _, inst := range resp.Result {

		if obj, err := datakit.CloudObject2Json(inst.Description, `aliyun_elasticsearch`, inst, inst.InstanceId, e.ExcludeInstanceIDs, e.InstancesIDs); obj != nil {
			objs = append(objs, obj)
		} else {
			if err != nil {
				moduleLogger.Errorf("%s", err)
			}
		}
	}

	if len(objs) <= 0 {
		return
	}

	data, err := json.Marshal(&objs)
	if err != nil {
		moduleLogger.Errorf("%s", err)
		return
	}
	io.NamedFeed(data, io.Object, inputName)
}