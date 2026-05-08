package models

import (
    "encoding/json"
    "errors"
)

// RuleList Model
type ProxyRuleList struct {
    // Total number of rules that matched your query.
    Total int `json:"total"`
    // List of rules.
    Rules []ProxyRule `json:"rules"`

    // Used by Decode() method
    data []byte
}

func (model ProxyRuleList) New(data []byte) *ProxyRuleList {
    model.data = data
    return &model
}

func (model *ProxyRuleList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}