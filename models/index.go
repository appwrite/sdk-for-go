package models

import (
	"encoding/json"
	"errors"
)

// Index Model
type Index struct {
	// Index ID.
	Id string `json:"$id"`
	// Index creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Index update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Index key.
	Key string `json:"key"`
	// Optional description for the index.
	Description string `json:"description"`
	// Index type.
	Type string `json:"type"`
	// Index status. Possible values: `available`, `processing`, `deleting`,
	// `stuck`, or `failed`
	Status string `json:"status"`
	// Error message. Displays error generated on failure of creating or deleting
	// an index.
	Error string `json:"error"`
	// Index attributes.
	Attributes []string `json:"attributes"`
	// Index attributes length.
	Lengths []int `json:"lengths"`
	// Index orders.
	Orders []string `json:"orders"`

	// Used by Decode() method
	data []byte
}

func (model Index) New(data []byte) *Index {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Index) UnmarshalJSON(data []byte) error {
	type alias Index
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id          json.RawMessage `json:"$id"`
		CreatedAt   json.RawMessage `json:"$createdAt"`
		UpdatedAt   json.RawMessage `json:"$updatedAt"`
		Key         json.RawMessage `json:"key"`
		Description json.RawMessage `json:"description"`
		Type        json.RawMessage `json:"type"`
		Error       json.RawMessage `json:"error"`
		Attributes  json.RawMessage `json:"attributes"`
		Lengths     json.RawMessage `json:"lengths"`
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
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Error == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Error); err != nil {
			return err
		}
	}
	if present.Attributes == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Attributes); err != nil {
			return err
		}
	}
	if present.Lengths == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Lengths); err != nil {
			return err
		}
	}

	*model = Index(decoded)
	return nil
}

func (model *Index) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
