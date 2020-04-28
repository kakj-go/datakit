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

// CheckInstanceExist invokes the rds.CheckInstanceExist API synchronously
// api document: https://help.aliyun.com/api/rds/checkinstanceexist.html
func (client *Client) CheckInstanceExist(request *CheckInstanceExistRequest) (response *CheckInstanceExistResponse, err error) {
	response = CreateCheckInstanceExistResponse()
	err = client.DoAction(request, response)
	return
}

// CheckInstanceExistWithChan invokes the rds.CheckInstanceExist API asynchronously
// api document: https://help.aliyun.com/api/rds/checkinstanceexist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckInstanceExistWithChan(request *CheckInstanceExistRequest) (<-chan *CheckInstanceExistResponse, <-chan error) {
	responseChan := make(chan *CheckInstanceExistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckInstanceExist(request)
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

// CheckInstanceExistWithCallback invokes the rds.CheckInstanceExist API asynchronously
// api document: https://help.aliyun.com/api/rds/checkinstanceexist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckInstanceExistWithCallback(request *CheckInstanceExistRequest, callback func(response *CheckInstanceExistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckInstanceExistResponse
		var err error
		defer close(result)
		response, err = client.CheckInstanceExist(request)
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

// CheckInstanceExistRequest is the request struct for api CheckInstanceExist
type CheckInstanceExistRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// CheckInstanceExistResponse is the response struct for api CheckInstanceExist
type CheckInstanceExistResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	IsExistInstance bool   `json:"IsExistInstance" xml:"IsExistInstance"`
}

// CreateCheckInstanceExistRequest creates a request to invoke CheckInstanceExist API
func CreateCheckInstanceExistRequest() (request *CheckInstanceExistRequest) {
	request = &CheckInstanceExistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CheckInstanceExist", "rds", "openAPI")
	return
}

// CreateCheckInstanceExistResponse creates a response to parse from CheckInstanceExist response
func CreateCheckInstanceExistResponse() (response *CheckInstanceExistResponse) {
	response = &CheckInstanceExistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
