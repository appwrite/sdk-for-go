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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AnalyticsMetricList) UnmarshalJSON(data []byte) error {
	type alias AnalyticsMetricList
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

	*model = AnalyticsMetricList(decoded)
	return nil
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
