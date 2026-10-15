package models

import (
	"encoding/json"
	"errors"
)

// Variable Model
type Variable struct {
	// Variable ID.
	Id string `json:"$id"`
	// Variable creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Variable creation date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Variable key.
	Key string `json:"key"`
	// Variable value.
	Value string `json:"value"`
	// Variable secret flag. Secret variables can only be updated or deleted, but
	// never read.
	Secret bool `json:"secret"`
	// Service to which the variable belongs. Possible values are "project",
	// "function"
	ResourceType string `json:"resourceType"`
	// ID of resource to which the variable belongs. If resourceType is "project",
	// it is empty. If resourceType is "function", it is ID of the function.
	ResourceId string `json:"resourceId"`

	// Used by Decode() method
	data []byte
}

func (model Variable) New(data []byte) *Variable {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Variable) UnmarshalJSON(data []byte) error {
	type alias Variable
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
		Key          json.RawMessage `json:"key"`
		Value        json.RawMessage `json:"value"`
		Secret       json.RawMessage `json:"secret"`
		ResourceType json.RawMessage `json:"resourceType"`
		ResourceId   json.RawMessage `json:"resourceId"`
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
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.Key == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Key); err != nil {
			return err
		}
	}
	if present.Value == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Value); err != nil {
			return err
		}
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Secret); err != nil {
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

	*model = Variable(decoded)
	return nil
}

func (model *Variable) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
