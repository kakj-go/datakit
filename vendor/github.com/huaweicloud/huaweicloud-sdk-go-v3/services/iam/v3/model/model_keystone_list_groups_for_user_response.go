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

// Response Object
type KeystoneListGroupsForUserResponse struct {
	// 用户组信息列表。
	Groups         *[]KeystoneGroupResult `json:"groups,omitempty"`
	Links          *Links                 `json:"links,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o KeystoneListGroupsForUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListGroupsForUserResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListGroupsForUserResponse", string(data)}, " ")
}
