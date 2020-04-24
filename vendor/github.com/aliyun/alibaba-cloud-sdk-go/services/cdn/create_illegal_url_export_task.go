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

// CreateIllegalUrlExportTask invokes the cdn.CreateIllegalUrlExportTask API synchronously
// api document: https://help.aliyun.com/api/cdn/createillegalurlexporttask.html
func (client *Client) CreateIllegalUrlExportTask(request *CreateIllegalUrlExportTaskRequest) (response *CreateIllegalUrlExportTaskResponse, err error) {
	response = CreateCreateIllegalUrlExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateIllegalUrlExportTaskWithChan invokes the cdn.CreateIllegalUrlExportTask API asynchronously
// api document: https://help.aliyun.com/api/cdn/createillegalurlexporttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateIllegalUrlExportTaskWithChan(request *CreateIllegalUrlExportTaskRequest) (<-chan *CreateIllegalUrlExportTaskResponse, <-chan error) {
	responseChan := make(chan *CreateIllegalUrlExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateIllegalUrlExportTask(request)
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

// CreateIllegalUrlExportTaskWithCallback invokes the cdn.CreateIllegalUrlExportTask API asynchronously
// api document: https://help.aliyun.com/api/cdn/createillegalurlexporttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateIllegalUrlExportTaskWithCallback(request *CreateIllegalUrlExportTaskRequest, callback func(response *CreateIllegalUrlExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateIllegalUrlExportTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateIllegalUrlExportTask(request)
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

// CreateIllegalUrlExportTaskRequest is the request struct for api CreateIllegalUrlExportTask
type CreateIllegalUrlExportTaskRequest struct {
	*requests.RpcRequest
	TaskName  string           `position:"Query" name:"TaskName"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	TimePoint string           `position:"Query" name:"TimePoint"`
}

// CreateIllegalUrlExportTaskResponse is the response struct for api CreateIllegalUrlExportTask
type CreateIllegalUrlExportTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateIllegalUrlExportTaskRequest creates a request to invoke CreateIllegalUrlExportTask API
func CreateCreateIllegalUrlExportTaskRequest() (request *CreateIllegalUrlExportTaskRequest) {
	request = &CreateIllegalUrlExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "CreateIllegalUrlExportTask", "", "")
	return
}

// CreateCreateIllegalUrlExportTaskResponse creates a response to parse from CreateIllegalUrlExportTask response
func CreateCreateIllegalUrlExportTaskResponse() (response *CreateIllegalUrlExportTaskResponse) {
	response = &CreateIllegalUrlExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
