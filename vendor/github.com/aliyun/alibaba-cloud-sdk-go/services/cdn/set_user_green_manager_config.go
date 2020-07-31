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

// SetUserGreenManagerConfig invokes the cdn.SetUserGreenManagerConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setusergreenmanagerconfig.html
func (client *Client) SetUserGreenManagerConfig(request *SetUserGreenManagerConfigRequest) (response *SetUserGreenManagerConfigResponse, err error) {
	response = CreateSetUserGreenManagerConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetUserGreenManagerConfigWithChan invokes the cdn.SetUserGreenManagerConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setusergreenmanagerconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUserGreenManagerConfigWithChan(request *SetUserGreenManagerConfigRequest) (<-chan *SetUserGreenManagerConfigResponse, <-chan error) {
	responseChan := make(chan *SetUserGreenManagerConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetUserGreenManagerConfig(request)
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

// SetUserGreenManagerConfigWithCallback invokes the cdn.SetUserGreenManagerConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setusergreenmanagerconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUserGreenManagerConfigWithCallback(request *SetUserGreenManagerConfigRequest, callback func(response *SetUserGreenManagerConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetUserGreenManagerConfigResponse
		var err error
		defer close(result)
		response, err = client.SetUserGreenManagerConfig(request)
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

// SetUserGreenManagerConfigRequest is the request struct for api SetUserGreenManagerConfig
type SetUserGreenManagerConfigRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	Quota         string           `position:"Query" name:"Quota"`
	Ratio         string           `position:"Query" name:"Ratio"`
}

// SetUserGreenManagerConfigResponse is the response struct for api SetUserGreenManagerConfig
type SetUserGreenManagerConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetUserGreenManagerConfigRequest creates a request to invoke SetUserGreenManagerConfig API
func CreateSetUserGreenManagerConfigRequest() (request *SetUserGreenManagerConfigRequest) {
	request = &SetUserGreenManagerConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "SetUserGreenManagerConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateSetUserGreenManagerConfigResponse creates a response to parse from SetUserGreenManagerConfig response
func CreateSetUserGreenManagerConfigResponse() (response *SetUserGreenManagerConfigResponse) {
	response = &SetUserGreenManagerConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
