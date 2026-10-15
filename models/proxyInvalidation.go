package models

import (
	"encoding/json"
	"errors"
)

// Invalidation Model
type ProxyInvalidation struct {
	// Domain name.
	Domain string `json:"domain"`
	// Invalidation type. Possible values are "tag", "path", or "all".
	Type string `json:"type"`
	// Invalidated reference. Depending on type this is a cache tag name, a URL
	// path, or empty when type is all.
	Reference string `json:"reference"`
	// Invalidation status.
	Status string `json:"status"`

	// Used by Decode() method
	data []byte
}

func (model ProxyInvalidation) New(data []byte) *ProxyInvalidation {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *ProxyInvalidation) UnmarshalJSON(data []byte) error {
	type alias ProxyInvalidation
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Domain    json.RawMessage `json:"domain"`
		Type      json.RawMessage `json:"type"`
		Reference json.RawMessage `json:"reference"`
		Status    json.RawMessage `json:"status"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Domain == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Domain); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Reference == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Reference); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Status); err != nil {
			return err
		}
	}

	*model = ProxyInvalidation(decoded)
	return nil
}

func (model *ProxyInvalidation) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
