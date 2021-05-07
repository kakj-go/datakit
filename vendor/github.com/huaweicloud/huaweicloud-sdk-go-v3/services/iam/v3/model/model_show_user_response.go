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
type ShowUserResponse struct {
	User           *ShowUserResult `json:"user,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowUserResponse struct{}"
	}

	return strings.Join([]string{"ShowUserResponse", string(data)}, " ")
}