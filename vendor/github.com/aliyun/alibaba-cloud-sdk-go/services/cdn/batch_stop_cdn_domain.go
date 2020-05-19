package cdn

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

// BatchStopCdnDomain invokes the cdn.BatchStopCdnDomain API synchronously
// api document: https://help.aliyun.com/api/cdn/batchstopcdndomain.html
func (client *Client) BatchStopCdnDomain(request *BatchStopCdnDomainRequest) (response *BatchStopCdnDomainResponse, err error) {
	response = CreateBatchStopCdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// BatchStopCdnDomainWithChan invokes the cdn.BatchStopCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/batchstopcdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchStopCdnDomainWithChan(request *BatchStopCdnDomainRequest) (<-chan *BatchStopCdnDomainResponse, <-chan error) {
	responseChan := make(chan *BatchStopCdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchStopCdnDomain(request)
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

// BatchStopCdnDomainWithCallback invokes the cdn.BatchStopCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/batchstopcdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchStopCdnDomainWithCallback(request *BatchStopCdnDomainRequest, callback func(response *BatchStopCdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchStopCdnDomainResponse
		var err error
		defer close(result)
		response, err = client.BatchStopCdnDomain(request)
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

// BatchStopCdnDomainRequest is the request struct for api BatchStopCdnDomain
type BatchStopCdnDomainRequest struct {
	*requests.RpcRequest
	DomainNames   string           `position:"Query" name:"DomainNames"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// BatchStopCdnDomainResponse is the response struct for api BatchStopCdnDomain
type BatchStopCdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchStopCdnDomainRequest creates a request to invoke BatchStopCdnDomain API
func CreateBatchStopCdnDomainRequest() (request *BatchStopCdnDomainRequest) {
	request = &BatchStopCdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "BatchStopCdnDomain", "", "")
	return
}

// CreateBatchStopCdnDomainResponse creates a response to parse from BatchStopCdnDomain response
func CreateBatchStopCdnDomainResponse() (response *BatchStopCdnDomainResponse) {
	response = &BatchStopCdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
