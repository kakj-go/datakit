package ons

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

// OnsWarnCreate invokes the ons.OnsWarnCreate API synchronously
// api document: https://help.aliyun.com/api/ons/onswarncreate.html
func (client *Client) OnsWarnCreate(request *OnsWarnCreateRequest) (response *OnsWarnCreateResponse, err error) {
	response = CreateOnsWarnCreateResponse()
	err = client.DoAction(request, response)
	return
}

// OnsWarnCreateWithChan invokes the ons.OnsWarnCreate API asynchronously
// api document: https://help.aliyun.com/api/ons/onswarncreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsWarnCreateWithChan(request *OnsWarnCreateRequest) (<-chan *OnsWarnCreateResponse, <-chan error) {
	responseChan := make(chan *OnsWarnCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsWarnCreate(request)
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

// OnsWarnCreateWithCallback invokes the ons.OnsWarnCreate API asynchronously
// api document: https://help.aliyun.com/api/ons/onswarncreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsWarnCreateWithCallback(request *OnsWarnCreateRequest, callback func(response *OnsWarnCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsWarnCreateResponse
		var err error
		defer close(result)
		response, err = client.OnsWarnCreate(request)
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

// OnsWarnCreateRequest is the request struct for api OnsWarnCreate
type OnsWarnCreateRequest struct {
	*requests.RpcRequest
	Level      string `position:"Query" name:"Level"`
	GroupId    string `position:"Query" name:"GroupId"`
	DelayTime  string `position:"Query" name:"DelayTime"`
	Threshold  string `position:"Query" name:"Threshold"`
	AlertTime  string `position:"Query" name:"AlertTime"`
	InstanceId string `position:"Query" name:"InstanceId"`
	BlockTime  string `position:"Query" name:"BlockTime"`
	Topic      string `position:"Query" name:"Topic"`
	Contacts   string `position:"Query" name:"Contacts"`
}

// OnsWarnCreateResponse is the response struct for api OnsWarnCreate
type OnsWarnCreateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

// CreateOnsWarnCreateRequest creates a request to invoke OnsWarnCreate API
func CreateOnsWarnCreateRequest() (request *OnsWarnCreateRequest) {
	request = &OnsWarnCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsWarnCreate", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsWarnCreateResponse creates a response to parse from OnsWarnCreate response
func CreateOnsWarnCreateResponse() (response *OnsWarnCreateResponse) {
	response = &OnsWarnCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}