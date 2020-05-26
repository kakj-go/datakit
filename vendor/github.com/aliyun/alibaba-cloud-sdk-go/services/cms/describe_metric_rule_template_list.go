package cms

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

// DescribeMetricRuleTemplateList invokes the cms.DescribeMetricRuleTemplateList API synchronously
// api document: https://help.aliyun.com/api/cms/describemetricruletemplatelist.html
func (client *Client) DescribeMetricRuleTemplateList(request *DescribeMetricRuleTemplateListRequest) (response *DescribeMetricRuleTemplateListResponse, err error) {
	response = CreateDescribeMetricRuleTemplateListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMetricRuleTemplateListWithChan invokes the cms.DescribeMetricRuleTemplateList API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricruletemplatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricRuleTemplateListWithChan(request *DescribeMetricRuleTemplateListRequest) (<-chan *DescribeMetricRuleTemplateListResponse, <-chan error) {
	responseChan := make(chan *DescribeMetricRuleTemplateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMetricRuleTemplateList(request)
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

// DescribeMetricRuleTemplateListWithCallback invokes the cms.DescribeMetricRuleTemplateList API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricruletemplatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricRuleTemplateListWithCallback(request *DescribeMetricRuleTemplateListRequest, callback func(response *DescribeMetricRuleTemplateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMetricRuleTemplateListResponse
		var err error
		defer close(result)
		response, err = client.DescribeMetricRuleTemplateList(request)
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

// DescribeMetricRuleTemplateListRequest is the request struct for api DescribeMetricRuleTemplateList
type DescribeMetricRuleTemplateListRequest struct {
	*requests.RpcRequest
	History    requests.Boolean `position:"Query" name:"History"`
	TemplateId requests.Integer `position:"Query" name:"TemplateId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Name       string           `position:"Query" name:"Name"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Keyword    string           `position:"Query" name:"Keyword"`
}

// DescribeMetricRuleTemplateListResponse is the response struct for api DescribeMetricRuleTemplateList
type DescribeMetricRuleTemplateListResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Success   bool      `json:"Success" xml:"Success"`
	Code      int       `json:"Code" xml:"Code"`
	Message   string    `json:"Message" xml:"Message"`
	Total     int64     `json:"Total" xml:"Total"`
	Templates Templates `json:"Templates" xml:"Templates"`
}

// CreateDescribeMetricRuleTemplateListRequest creates a request to invoke DescribeMetricRuleTemplateList API
func CreateDescribeMetricRuleTemplateListRequest() (request *DescribeMetricRuleTemplateListRequest) {
	request = &DescribeMetricRuleTemplateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMetricRuleTemplateList", "cms", "openAPI")
	return
}

// CreateDescribeMetricRuleTemplateListResponse creates a response to parse from DescribeMetricRuleTemplateList response
func CreateDescribeMetricRuleTemplateListResponse() (response *DescribeMetricRuleTemplateListResponse) {
	response = &DescribeMetricRuleTemplateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
