package aegis

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

// SaveWhiteListStrategy invokes the aegis.SaveWhiteListStrategy API synchronously
// api document: https://help.aliyun.com/api/aegis/savewhiteliststrategy.html
func (client *Client) SaveWhiteListStrategy(request *SaveWhiteListStrategyRequest) (response *SaveWhiteListStrategyResponse, err error) {
	response = CreateSaveWhiteListStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// SaveWhiteListStrategyWithChan invokes the aegis.SaveWhiteListStrategy API asynchronously
// api document: https://help.aliyun.com/api/aegis/savewhiteliststrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveWhiteListStrategyWithChan(request *SaveWhiteListStrategyRequest) (<-chan *SaveWhiteListStrategyResponse, <-chan error) {
	responseChan := make(chan *SaveWhiteListStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveWhiteListStrategy(request)
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

// SaveWhiteListStrategyWithCallback invokes the aegis.SaveWhiteListStrategy API asynchronously
// api document: https://help.aliyun.com/api/aegis/savewhiteliststrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveWhiteListStrategyWithCallback(request *SaveWhiteListStrategyRequest, callback func(response *SaveWhiteListStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveWhiteListStrategyResponse
		var err error
		defer close(result)
		response, err = client.SaveWhiteListStrategy(request)
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

// SaveWhiteListStrategyRequest is the request struct for api SaveWhiteListStrategy
type SaveWhiteListStrategyRequest struct {
	*requests.RpcRequest
	StrategyName string           `position:"Query" name:"StrategyName"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	StudyTime    requests.Integer `position:"Query" name:"StudyTime"`
	StrategyId   requests.Integer `position:"Query" name:"StrategyId"`
	Lang         string           `position:"Query" name:"Lang"`
}

// SaveWhiteListStrategyResponse is the response struct for api SaveWhiteListStrategy
type SaveWhiteListStrategyResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	StrategyId int    `json:"StrategyId" xml:"StrategyId"`
}

// CreateSaveWhiteListStrategyRequest creates a request to invoke SaveWhiteListStrategy API
func CreateSaveWhiteListStrategyRequest() (request *SaveWhiteListStrategyRequest) {
	request = &SaveWhiteListStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "SaveWhiteListStrategy", "vipaegis", "openAPI")
	return
}

// CreateSaveWhiteListStrategyResponse creates a response to parse from SaveWhiteListStrategy response
func CreateSaveWhiteListStrategyResponse() (response *SaveWhiteListStrategyResponse) {
	response = &SaveWhiteListStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
