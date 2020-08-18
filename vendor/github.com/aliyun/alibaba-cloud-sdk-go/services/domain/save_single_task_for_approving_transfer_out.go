package domain

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

// SaveSingleTaskForApprovingTransferOut invokes the domain.SaveSingleTaskForApprovingTransferOut API synchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforapprovingtransferout.html
func (client *Client) SaveSingleTaskForApprovingTransferOut(request *SaveSingleTaskForApprovingTransferOutRequest) (response *SaveSingleTaskForApprovingTransferOutResponse, err error) {
	response = CreateSaveSingleTaskForApprovingTransferOutResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForApprovingTransferOutWithChan invokes the domain.SaveSingleTaskForApprovingTransferOut API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforapprovingtransferout.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForApprovingTransferOutWithChan(request *SaveSingleTaskForApprovingTransferOutRequest) (<-chan *SaveSingleTaskForApprovingTransferOutResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForApprovingTransferOutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForApprovingTransferOut(request)
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

// SaveSingleTaskForApprovingTransferOutWithCallback invokes the domain.SaveSingleTaskForApprovingTransferOut API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforapprovingtransferout.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForApprovingTransferOutWithCallback(request *SaveSingleTaskForApprovingTransferOutRequest, callback func(response *SaveSingleTaskForApprovingTransferOutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForApprovingTransferOutResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForApprovingTransferOut(request)
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

// SaveSingleTaskForApprovingTransferOutRequest is the request struct for api SaveSingleTaskForApprovingTransferOut
type SaveSingleTaskForApprovingTransferOutRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// SaveSingleTaskForApprovingTransferOutResponse is the response struct for api SaveSingleTaskForApprovingTransferOut
type SaveSingleTaskForApprovingTransferOutResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForApprovingTransferOutRequest creates a request to invoke SaveSingleTaskForApprovingTransferOut API
func CreateSaveSingleTaskForApprovingTransferOutRequest() (request *SaveSingleTaskForApprovingTransferOutRequest) {
	request = &SaveSingleTaskForApprovingTransferOutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForApprovingTransferOut", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveSingleTaskForApprovingTransferOutResponse creates a response to parse from SaveSingleTaskForApprovingTransferOut response
func CreateSaveSingleTaskForApprovingTransferOutResponse() (response *SaveSingleTaskForApprovingTransferOutResponse) {
	response = &SaveSingleTaskForApprovingTransferOutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
