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

// DescribeDBInstanceTDE invokes the rds.DescribeDBInstanceTDE API synchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancetde.html
func (client *Client) DescribeDBInstanceTDE(request *DescribeDBInstanceTDERequest) (response *DescribeDBInstanceTDEResponse, err error) {
	response = CreateDescribeDBInstanceTDEResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceTDEWithChan invokes the rds.DescribeDBInstanceTDE API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancetde.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceTDEWithChan(request *DescribeDBInstanceTDERequest) (<-chan *DescribeDBInstanceTDEResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceTDEResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceTDE(request)
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

// DescribeDBInstanceTDEWithCallback invokes the rds.DescribeDBInstanceTDE API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancetde.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceTDEWithCallback(request *DescribeDBInstanceTDERequest, callback func(response *DescribeDBInstanceTDEResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceTDEResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceTDE(request)
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

// DescribeDBInstanceTDERequest is the request struct for api DescribeDBInstanceTDE
type DescribeDBInstanceTDERequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeDBInstanceTDEResponse is the response struct for api DescribeDBInstanceTDE
type DescribeDBInstanceTDEResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	TDEStatus string                           `json:"TDEStatus" xml:"TDEStatus"`
	Databases DatabasesInDescribeDBInstanceTDE `json:"Databases" xml:"Databases"`
}

// CreateDescribeDBInstanceTDERequest creates a request to invoke DescribeDBInstanceTDE API
func CreateDescribeDBInstanceTDERequest() (request *DescribeDBInstanceTDERequest) {
	request = &DescribeDBInstanceTDERequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceTDE", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceTDEResponse creates a response to parse from DescribeDBInstanceTDE response
func CreateDescribeDBInstanceTDEResponse() (response *DescribeDBInstanceTDEResponse) {
	response = &DescribeDBInstanceTDEResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
