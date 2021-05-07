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
type CreateLoginTokenRequest struct {
	Body *CreateLoginTokenRequestBody `json:"body,omitempty"`
}

func (o CreateLoginTokenRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoginTokenRequest struct{}"
	}

	return strings.Join([]string{"CreateLoginTokenRequest", string(data)}, " ")
}