package models

import (
	"encoding/json"
	"errors"
)

// Block Model
type Block struct {
	// Block creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Resource type that is blocked
	ResourceType string `json:"resourceType"`
	// Resource identifier that is blocked
	ResourceId string `json:"resourceId"`
	// Block mode. full blocks reads and writes; readOnly blocks writes only.
	Mode string `json:"mode"`
	// Block expiration date in ISO 8601 format. Can be null if the block does not
	// expire.
	ExpiredAt *string `json:"expiredAt"`

	// Used by Decode() method
	data []byte
}

func (model Block) New(data []byte) *Block {
	model.data = data
	return &model
}

func (model *Block) UnmarshalJSON(data []byte) error {
	type alias Block
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$createdAt\":\"\",\"resourceType\":\"\",\"resourceId\":\"\",\"mode\":\"full\"}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
}

func (model *Block) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
