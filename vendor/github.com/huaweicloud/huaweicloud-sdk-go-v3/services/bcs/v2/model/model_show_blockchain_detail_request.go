/*
 * BCS
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
type ShowBlockchainDetailRequest struct {
	BlockchainId string `json:"blockchain_id"`
}

func (o ShowBlockchainDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlockchainDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainDetailRequest", string(data)}, " ")
}
