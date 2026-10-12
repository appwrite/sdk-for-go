package models

import (
	"encoding/json"
	"errors"
)

// AnalyticsMetricList Model
type AnalyticsMetricList struct {
	// Total number of metric rows returned.
	Total int `json:"total"`
	// Metric rows: one per time bucket, one per dimension value, or a single row
	// for the flat aggregate.
	Metrics []AnalyticsMetric `json:"metrics"`

	// Used by Decode() method
	data []byte
}

func (model AnalyticsMetricList) New(data []byte) *AnalyticsMetricList {
	model.data = data
	return &model
}

func (model *AnalyticsMetricList) UnmarshalJSON(data []byte) error {
	type alias AnalyticsMetricList
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

func (model *AnalyticsMetricList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
