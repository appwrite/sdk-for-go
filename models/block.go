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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Block) UnmarshalJSON(data []byte) error {
	type alias Block
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		CreatedAt    json.RawMessage `json:"$createdAt"`
		ResourceType json.RawMessage `json:"resourceType"`
		ResourceId   json.RawMessage `json:"resourceId"`
		Mode         json.RawMessage `json:"mode"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.ResourceType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceType); err != nil {
			return err
		}
	}
	if present.ResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceId); err != nil {
			return err
		}
	}
	if present.Mode == nil {
		if err := json.Unmarshal([]byte("\"full\""), &decoded.Mode); err != nil {
			return err
		}
	}

	*model = Block(decoded)
	return nil
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
