package actiontrail

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

// DescribeRegions invokes the actiontrail.DescribeRegions API synchronously
// api document: https://help.aliyun.com/api/actiontrail/describeregions.html
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
	response = CreateDescribeRegionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRegionsWithChan invokes the actiontrail.DescribeRegions API asynchronously
// api document: https://help.aliyun.com/api/actiontrail/describeregions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRegionsWithChan(request *DescribeRegionsRequest) (<-chan *DescribeRegionsResponse, <-chan error) {
	responseChan := make(chan *DescribeRegionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRegions(request)
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

// DescribeRegionsWithCallback invokes the actiontrail.DescribeRegions API asynchronously
// api document: https://help.aliyun.com/api/actiontrail/describeregions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRegionsWithCallback(request *DescribeRegionsRequest, callback func(response *DescribeRegionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRegionsResponse
		var err error
		defer close(result)
		response, err = client.DescribeRegions(request)
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

// DescribeRegionsRequest is the request struct for api DescribeRegions
type DescribeRegionsRequest struct {
	*requests.RpcRequest
}

// DescribeRegionsResponse is the response struct for api DescribeRegions
type DescribeRegionsResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Regions   Regions `json:"Regions" xml:"Regions"`
}

// CreateDescribeRegionsRequest creates a request to invoke DescribeRegions API
func CreateDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Actiontrail", "2017-12-04", "DescribeRegions", "actiontrail", "openAPI")
	return
}

// CreateDescribeRegionsResponse creates a response to parse from DescribeRegions response
func CreateDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
