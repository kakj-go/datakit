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
type TokenUserResult struct {
	// IAM用户名。
	Name string `json:"name"`
	// IAM用户ID。
	Id string `json:"id"`
	// 密码过期时间（UTC时间），“”表示密码不过期。
	PasswordExpiresAt string                 `json:"password_expires_at"`
	Domain            *TokenUserDomainResult `json:"domain"`
}

func (o TokenUserResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TokenUserResult struct{}"
	}

	return strings.Join([]string{"TokenUserResult", string(data)}, " ")
}