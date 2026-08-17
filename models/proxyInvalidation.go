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
