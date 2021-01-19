package slb

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

// DescribeLoadBalancerUDPListenerAttribute invokes the slb.DescribeLoadBalancerUDPListenerAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancerudplistenerattribute.html
func (client *Client) DescribeLoadBalancerUDPListenerAttribute(request *DescribeLoadBalancerUDPListenerAttributeRequest) (response *DescribeLoadBalancerUDPListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerUDPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerUDPListenerAttributeWithChan invokes the slb.DescribeLoadBalancerUDPListenerAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancerudplistenerattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoadBalancerUDPListenerAttributeWithChan(request *DescribeLoadBalancerUDPListenerAttributeRequest) (<-chan *DescribeLoadBalancerUDPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerUDPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerUDPListenerAttribute(request)
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

// DescribeLoadBalancerUDPListenerAttributeWithCallback invokes the slb.DescribeLoadBalancerUDPListenerAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancerudplistenerattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoadBalancerUDPListenerAttributeWithCallback(request *DescribeLoadBalancerUDPListenerAttributeRequest, callback func(response *DescribeLoadBalancerUDPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerUDPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerUDPListenerAttribute(request)
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

// DescribeLoadBalancerUDPListenerAttributeRequest is the request struct for api DescribeLoadBalancerUDPListenerAttribute
type DescribeLoadBalancerUDPListenerAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerUDPListenerAttributeResponse is the response struct for api DescribeLoadBalancerUDPListenerAttribute
type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId                 string                                               `json:"RequestId" xml:"RequestId"`
	ListenerPort              int                                                  `json:"ListenerPort" xml:"ListenerPort"`
	BackendServerPort         int                                                  `json:"BackendServerPort" xml:"BackendServerPort"`
	Status                    string                                               `json:"Status" xml:"Status"`
	Bandwidth                 int                                                  `json:"Bandwidth" xml:"Bandwidth"`
	Scheduler                 string                                               `json:"Scheduler" xml:"Scheduler"`
	PersistenceTimeout        int                                                  `json:"PersistenceTimeout" xml:"PersistenceTimeout"`
	HealthCheck               string                                               `json:"HealthCheck" xml:"HealthCheck"`
	HealthyThreshold          int                                                  `json:"HealthyThreshold" xml:"HealthyThreshold"`
	UnhealthyThreshold        int                                                  `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int                                                  `json:"HealthCheckConnectTimeout" xml:"HealthCheckConnectTimeout"`
	HealthCheckConnectPort    int                                                  `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthCheckInterval       int                                                  `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	HealthCheckReq            string                                               `json:"HealthCheckReq" xml:"HealthCheckReq"`
	HealthCheckExp            string                                               `json:"HealthCheckExp" xml:"HealthCheckExp"`
	MaxConnection             int                                                  `json:"MaxConnection" xml:"MaxConnection"`
	VServerGroupId            string                                               `json:"VServerGroupId" xml:"VServerGroupId"`
	MasterSlaveServerGroupId  string                                               `json:"MasterSlaveServerGroupId" xml:"MasterSlaveServerGroupId"`
	AclId                     string                                               `json:"AclId" xml:"AclId"`
	AclType                   string                                               `json:"AclType" xml:"AclType"`
	AclStatus                 string                                               `json:"AclStatus" xml:"AclStatus"`
	VpcIds                    string                                               `json:"VpcIds" xml:"VpcIds"`
	Description               string                                               `json:"Description" xml:"Description"`
	ConnectionDrain           string                                               `json:"ConnectionDrain" xml:"ConnectionDrain"`
	ConnectionDrainTimeout    int                                                  `json:"ConnectionDrainTimeout" xml:"ConnectionDrainTimeout"`
	AclIds                    AclIdsInDescribeLoadBalancerUDPListenerAttribute     `json:"AclIds" xml:"AclIds"`
	PortRanges                PortRangesInDescribeLoadBalancerUDPListenerAttribute `json:"PortRanges" xml:"PortRanges"`
}

// CreateDescribeLoadBalancerUDPListenerAttributeRequest creates a request to invoke DescribeLoadBalancerUDPListenerAttribute API
func CreateDescribeLoadBalancerUDPListenerAttributeRequest() (request *DescribeLoadBalancerUDPListenerAttributeRequest) {
	request = &DescribeLoadBalancerUDPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerUDPListenerAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerUDPListenerAttributeResponse creates a response to parse from DescribeLoadBalancerUDPListenerAttribute response
func CreateDescribeLoadBalancerUDPListenerAttributeResponse() (response *DescribeLoadBalancerUDPListenerAttributeResponse) {
	response = &DescribeLoadBalancerUDPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
