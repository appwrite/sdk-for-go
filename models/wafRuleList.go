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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *WafRuleList) UnmarshalJSON(data []byte) error {
	type alias WafRuleList
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Total json.RawMessage `json:"total"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Total == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Total); err != nil {
			return err
		}
	}

	*model = WafRuleList(decoded)
	return nil
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
