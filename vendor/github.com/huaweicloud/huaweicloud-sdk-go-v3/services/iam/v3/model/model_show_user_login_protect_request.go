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
type ShowUserLoginProtectRequest struct {
	UserId string `json:"user_id"`
}

func (o ShowUserLoginProtectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowUserLoginProtectRequest struct{}"
	}

	return strings.Join([]string{"ShowUserLoginProtectRequest", string(data)}, " ")
}