/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

//
type MappingOption struct {
	// 将联邦用户映射为本地用户的规则列表。
	Rules []MappingRules `json:"rules"`
}

func (o MappingOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MappingOption struct{}"
	}

	return strings.Join([]string{"MappingOption", string(data)}, " ")
}