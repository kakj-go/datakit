package aliyuncost

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func staticClient() *bssopenapi.Client {
	//client, err := bssopenapi.NewClientWithAccessKey(`cn-hangzhou`, `LTAI4Fc72xGdZKKr6cTBV72S`, `QXZ4FFCq3yhN5TCGC9rj1kBNZNJksc`)
	client, err := bssopenapi.NewClientWithAccessKey(`cn-hangzhou`, `LTAIlsWpTrg1vUf4`, `dy5lQzWpU17RDNHGCj84LBDhoU9LVU`)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

//https://help.aliyun.com/document_detail/87997.html?spm=a2c4g.11186623.6.621.a5f8392dHi0imZ
func TestAccountBalance(t *testing.T) {

	cli := staticClient()

	req := bssopenapi.CreateQueryAccountBalanceRequest()

	// req := bssopenapi.CreateQueryProductListRequest()
	// req.PageNum = requests.NewInteger(0)
	// req.QueryTotalCount = requests.NewBoolean(true)

	resp, err := cli.QueryAccountBalance(req)
	if err != nil {
		log.Fatalln(err)
	}

	fields := map[string]interface{}{}
	tags := map[string]string{}

	tags[`Currency`] = resp.Data.Currency

	var fv float64
	if fv, err = strconv.ParseFloat(internal.NumberFormat(resp.Data.AvailableAmount), 64); err == nil {
		fields[`AvailableAmount`] = fv
	}
	if fv, err = strconv.ParseFloat(internal.NumberFormat(resp.Data.MybankCreditAmount), 64); err == nil {
		fields[`MybankCreditAmount`] = fv
	}
	if fv, err = strconv.ParseFloat(internal.NumberFormat(resp.Data.AvailableCashAmount), 64); err == nil {
		fields[`AvailableCashAmount`] = fv
	}
	if fv, err = strconv.ParseFloat(internal.NumberFormat(resp.Data.CreditAmount), 64); err == nil {
		fields[`CreditAmount`] = fv
	}

	log.Printf("%s", resp.String())
	log.Printf("tags: %v", tags)
	log.Printf("fields: %v", fields)
}

func TestAccountTransactions(t *testing.T) {

	tm, err := time.Parse("2006-01-02T15:04:05Z", "2020-02-15T00:21:12Z")
	if err == nil {
		log.Printf("tm: %v", tm)
	} else {
		log.Fatalf("%s", err)
	}
	return

	cli := staticClient()

	req := bssopenapi.CreateQueryAccountTransactionsRequest()
	req.PageSize = requests.NewInteger(300)
	now := time.Now().Truncate(time.Minute)
	start := now.Add(-time.Hour * 24).Format(`2006-01-02T15:04:05Z`)
	req.CreateTimeStart = start
	req.CreateTimeEnd = now.Format(`2006-01-02T15:04:05Z`)

	resp, err := cli.QueryAccountTransactions(req)
	if err != nil {
		log.Fatalln(err)
	}

	//og.Printf("%s", resp.String())

	//fmt.Printf("TotalCount=%d, PageSize=%d, PageNum=%d\n", resp.Data.TotalCount, resp.Data.PageSize, resp.Data.PageNum)

	for _, at := range resp.Data.AccountTransactionsList.AccountTransactionsListItem {
		log.Printf("%s - %s - %s, %s", at.TransactionTime, at.TransactionAccount, at.Amount, at.Balance)
	}
}

func TestQueryBill(t *testing.T) {

	cli := staticClient()

	req := bssopenapi.CreateQueryBillRequest()
	req.BillingCycle = fmt.Sprintf("%d-%d", 2019, 12)
	req.PageSize = requests.NewInteger(300)

	var respBill *bssopenapi.QueryBillResponse

	for {
		resp, err := cli.QueryBill(req)
		if err != nil {
			log.Fatalln(err)
		}
		if respBill == nil {
			respBill = resp
		} else {
			respBill.Data.Items.Item = append(respBill.Data.Items.Item, resp.Data.Items.Item...)
		}

		if resp.Data.TotalCount > 0 && resp.Data.PageNum*resp.Data.PageSize < resp.Data.TotalCount {
			req.PageNum = requests.NewInteger(resp.Data.PageNum + 1)
		} else {
			break
		}
	}

	for _, item := range respBill.Data.Items.Item {
		if item.PaymentTime != "" {
			fmt.Printf("%s -, %s, %v\n", item.PaymentTime, item.ProductName, item.PretaxAmount)
		}
	}

}

func TestQueryInstBill(t *testing.T) {

	cli := staticClient()

	req := bssopenapi.CreateQueryInstanceBillRequest()
	//today := time.Now()
	req.PageSize = requests.NewInteger(300)
	req.BillingCycle = "2019-10" // fmt.Sprintf("%d-%d", today.Year(), today.Month()) // `2019-10-01`

	resp, err := cli.QueryInstanceBill(req)
	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range resp.Data.Items.Item {
		if item.PaymentTime != "" {
			fmt.Printf("%s - %s, %v, %s\n", item.PaymentTime, item.ProductName, item.PretaxAmount, item.Tag)
		}
	}

}

func TestQueryOrder(t *testing.T) {

	cli := staticClient()

	req := bssopenapi.CreateQueryOrdersRequest()
	now := time.Now().Truncate(time.Hour)
	start := now.Add(-time.Hour * 24 * 30).Format(time.RFC3339) // Format(`2006-01-02T15:04:05Z`)
	start = strings.Replace(start, "+", "Z", -1)
	log.Printf("start=%s", start)
	req.CreateTimeStart = start
	//req.CreateTimeEnd = now.Format(`2006-01-02T15:04:05Z`)

	resp, err := cli.QueryOrders(req)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("TotalCount=%d, PageNum=%d, PageSize=%d\n", resp.Data.TotalCount, resp.Data.PageNum, resp.Data.PageSize)

	for _, item := range resp.Data.OrderList.Order {
		fmt.Printf("%s - %s, %v, %s\n", item.PaymentTime, item.PaymentStatus, item.PretaxAmount, item.Currency)
	}

}

func TestConfig(t *testing.T) {
	// if err := Cfg.Load(`./demo.toml`); err != nil {
	// 	fmt.Printf("%s", err)
	// }

	// fmt.Printf("%#v", Cfg.Boas[0].BiilInterval)
}

func TestSvr(t *testing.T) {

	// if err := Cfg.Load(`./demo.toml`); err != nil {
	// 	log.Fatalln(err)
	// }

	// logHandler, _ := log.NewStreamHandler(os.Stdout)

	// ll := log.NewDefault(logHandler)
	// ll.SetLevel(log.LevelDebug)

	// ll.Debugf("acckey: %s, accountInterval: %v", Cfg.Boas[0].AccessKeySecret, Cfg.Boas[0].AccountInterval)

	// svr := &AliyunBoaSvr{
	// 	logger: ll,
	// }

	// ctx, cancel := context.WithCancel(context.Background())

	// go func() {
	// 	svr.Start(ctx, nil)
	// }()

	// time.Sleep(10000 * time.Second)

	// cancel()

	// time.Sleep(1 * time.Second)
}
