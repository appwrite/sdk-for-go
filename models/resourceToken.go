package models

import (
	"encoding/json"
	"errors"
)

// ResourceToken Model
type ResourceToken struct {
	// Token ID.
	Id string `json:"$id"`
	// Token creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Resource ID.
	ResourceId string `json:"resourceId"`
	// Resource type.
	ResourceType string `json:"resourceType"`
	// Token expiration date in ISO 8601 format.
	Expire string `json:"expire"`
	// JWT encoded string.
	Secret string `json:"secret"`
	// Most recent access date in ISO 8601 format. This attribute is only updated
	// again after 24 hours.
	AccessedAt string `json:"accessedAt"`

	// Used by Decode() method
	data []byte
}

func (model ResourceToken) New(data []byte) *ResourceToken {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *ResourceToken) UnmarshalJSON(data []byte) error {
	type alias ResourceToken
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		ResourceId   json.RawMessage `json:"resourceId"`
		ResourceType json.RawMessage `json:"resourceType"`
		Expire       json.RawMessage `json:"expire"`
		Secret       json.RawMessage `json:"secret"`
		AccessedAt   json.RawMessage `json:"accessedAt"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.ResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceId); err != nil {
			return err
		}
	}
	if present.ResourceType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceType); err != nil {
			return err
		}
	}
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Secret); err != nil {
			return err
		}
	}
	if present.AccessedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AccessedAt); err != nil {
			return err
		}
	}

	*model = ResourceToken(decoded)
	return nil
}

func (model *ResourceToken) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
