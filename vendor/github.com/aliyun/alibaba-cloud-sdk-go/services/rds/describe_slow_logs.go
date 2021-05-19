package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSlowLogs invokes the rds.DescribeSlowLogs API synchronously
// api document: https://help.aliyun.com/api/rds/describeslowlogs.html
func (client *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
	response = CreateDescribeSlowLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSlowLogsWithChan invokes the rds.DescribeSlowLogs API asynchronously
// api document: https://help.aliyun.com/api/rds/describeslowlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowLogsWithChan(request *DescribeSlowLogsRequest) (<-chan *DescribeSlowLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeSlowLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSlowLogs(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeSlowLogsWithCallback invokes the rds.DescribeSlowLogs API asynchronously
// api document: https://help.aliyun.com/api/rds/describeslowlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowLogsWithCallback(request *DescribeSlowLogsRequest, callback func(response *DescribeSlowLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSlowLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSlowLogs(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeSlowLogsRequest is the request struct for api DescribeSlowLogs
type DescribeSlowLogsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SortKey              string           `position:"Query" name:"SortKey"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBName               string           `position:"Query" name:"DBName"`
}

// DescribeSlowLogsResponse is the response struct for api DescribeSlowLogs
type DescribeSlowLogsResponse struct {
	*responses.BaseResponse
	RequestId        string                  `json:"RequestId" xml:"RequestId"`
	DBInstanceId     string                  `json:"DBInstanceId" xml:"DBInstanceId"`
	Engine           string                  `json:"Engine" xml:"Engine"`
	StartTime        string                  `json:"StartTime" xml:"StartTime"`
	EndTime          string                  `json:"EndTime" xml:"EndTime"`
	TotalRecordCount int                     `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                     `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                     `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeSlowLogs `json:"Items" xml:"Items"`
}

// CreateDescribeSlowLogsRequest creates a request to invoke DescribeSlowLogs API
func CreateDescribeSlowLogsRequest() (request *DescribeSlowLogsRequest) {
	request = &DescribeSlowLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSlowLogs", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSlowLogsResponse creates a response to parse from DescribeSlowLogs response
func CreateDescribeSlowLogsResponse() (response *DescribeSlowLogsResponse) {
	response = &DescribeSlowLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
