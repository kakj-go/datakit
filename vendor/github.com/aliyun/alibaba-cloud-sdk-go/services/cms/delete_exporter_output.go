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

// DeleteExporterOutput invokes the cms.DeleteExporterOutput API synchronously
// api document: https://help.aliyun.com/api/cms/deleteexporteroutput.html
func (client *Client) DeleteExporterOutput(request *DeleteExporterOutputRequest) (response *DeleteExporterOutputResponse, err error) {
	response = CreateDeleteExporterOutputResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteExporterOutputWithChan invokes the cms.DeleteExporterOutput API asynchronously
// api document: https://help.aliyun.com/api/cms/deleteexporteroutput.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteExporterOutputWithChan(request *DeleteExporterOutputRequest) (<-chan *DeleteExporterOutputResponse, <-chan error) {
	responseChan := make(chan *DeleteExporterOutputResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteExporterOutput(request)
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

// DeleteExporterOutputWithCallback invokes the cms.DeleteExporterOutput API asynchronously
// api document: https://help.aliyun.com/api/cms/deleteexporteroutput.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteExporterOutputWithCallback(request *DeleteExporterOutputRequest, callback func(response *DeleteExporterOutputResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteExporterOutputResponse
		var err error
		defer close(result)
		response, err = client.DeleteExporterOutput(request)
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

// DeleteExporterOutputRequest is the request struct for api DeleteExporterOutput
type DeleteExporterOutputRequest struct {
	*requests.RpcRequest
	DestName string `position:"Query" name:"DestName"`
}

// DeleteExporterOutputResponse is the response struct for api DeleteExporterOutput
type DeleteExporterOutputResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteExporterOutputRequest creates a request to invoke DeleteExporterOutput API
func CreateDeleteExporterOutputRequest() (request *DeleteExporterOutputRequest) {
	request = &DeleteExporterOutputRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteExporterOutput", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteExporterOutputResponse creates a response to parse from DeleteExporterOutput response
func CreateDeleteExporterOutputResponse() (response *DeleteExporterOutputResponse) {
	response = &DeleteExporterOutputResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
