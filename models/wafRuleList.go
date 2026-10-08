package models

import (
	"encoding/json"
	"errors"
)

// WAFRuleList Model
type WafRuleList struct {
	// Total number of rules that matched your query.
	Total int `json:"total"`
	// List of rules.
	Rules []WafRule `json:"rules"`

	// Used by Decode() method
	data []byte
}

func (model WafRuleList) New(data []byte) *WafRuleList {
	model.data = data
	return &model
}

func (model *WafRuleList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
