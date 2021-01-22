/*
 * BCS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// BCS组织实例监控数据详情查询请求结构
type ListInstanceMetricRequestBody struct {
	// 实体类型，可选值如下 org     # 节点组织 plugin  # 插件 默认为org
	Type string `json:"type"`
	// 所属实体的名称
	EntityName string `json:"entity_name"`
	// 具体实例的名称
	InstanceName string `json:"instance_name"`
}

func (o ListInstanceMetricRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstanceMetricRequestBody struct{}"
	}

	return strings.Join([]string{"ListInstanceMetricRequestBody", string(data)}, " ")
}
