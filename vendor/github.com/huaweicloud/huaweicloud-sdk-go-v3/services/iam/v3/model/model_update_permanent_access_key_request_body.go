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
type UpdatePermanentAccessKeyRequestBody struct {
	Credential *UpdateCredentialOption `json:"credential"`
}

func (o UpdatePermanentAccessKeyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePermanentAccessKeyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePermanentAccessKeyRequestBody", string(data)}, " ")
}