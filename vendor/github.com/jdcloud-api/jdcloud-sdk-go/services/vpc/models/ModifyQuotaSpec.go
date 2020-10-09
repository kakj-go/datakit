// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type ModifyQuotaSpec struct {

    /* 资源类型，取值范围：vpc、elastic_ip、subnet、security_group、vpcpeering、network_interface（配额只统计辅助网卡）、acl、aclRule、routeTable、route、securityGroupRule  */
    Type string `json:"type"`

    /* type为vpc、elastic_ip、network_interface不设置, type为subnet、security_group、vpcpeering、acl、routeTable设置为vpcId, type为aclRule设置为aclId, type为route设置为routeTableId, type为securityGroupRule为securityGroupId (Optional) */
    ParentResourceId string `json:"parentResourceId"`

    /* 配额大小  */
    MaxLimit int `json:"maxLimit"`
}
