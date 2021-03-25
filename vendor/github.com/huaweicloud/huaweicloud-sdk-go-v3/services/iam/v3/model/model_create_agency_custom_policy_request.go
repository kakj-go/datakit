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

// Request Object
type CreateAgencyCustomPolicyRequest struct {
	Body *CreateAgencyCustomPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateAgencyCustomPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAgencyCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateAgencyCustomPolicyRequest", string(data)}, " ")
}
