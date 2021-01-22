/*
 * EVS
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
type ShowSnapshotResponse struct {
	Snapshot       *SnapshotDetails `json:"snapshot,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowSnapshotResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ShowSnapshotResponse", string(data)}, " ")
}
