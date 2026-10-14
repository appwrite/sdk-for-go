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
	Rules []interface{} `json:"rules"`

	// Used by Decode() method
	data []byte
}

func (model WafRuleList) New(data []byte) *WafRuleList {
	model.data = data
	return &model
}

func (model *WafRuleList) UnmarshalJSON(data []byte) error {
	type alias WafRuleList
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"total\":0}"), &fields); err != nil {
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
