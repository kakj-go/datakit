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
type KeystoneCreateUserTokenByPasswordRequest struct {
	Nocatalog *string                                       `json:"nocatalog,omitempty"`
	Body      *KeystoneCreateUserTokenByPasswordRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateUserTokenByPasswordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordRequest", string(data)}, " ")
}