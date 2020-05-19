package ecs

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

// ValueItem is a nested struct in ecs response
type ValueItem struct {
	Value              string `json:"Value" xml:"Value"`
	ExpiredTime        string `json:"ExpiredTime" xml:"ExpiredTime"`
	ZoneId             string `json:"ZoneId" xml:"ZoneId"`
	InstanceChargeType string `json:"InstanceChargeType" xml:"InstanceChargeType"`
	InstanceType       string `json:"InstanceType" xml:"InstanceType"`
	Count              int    `json:"Count" xml:"Count"`
	DiskCategory       string `json:"DiskCategory" xml:"DiskCategory"`
}
