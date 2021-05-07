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
type UpdateDomainProtectPolicyResponse struct {
	ProtectPolicy  *ProtectPolicyResult `json:"protect_policy,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateDomainProtectPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainProtectPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainProtectPolicyResponse", string(data)}, " ")
}