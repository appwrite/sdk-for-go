package models

import (
	"encoding/json"
	"errors"
)

// UsageBillingPlan Model
type UsageBillingPlan struct {
	// Bandwidth additional resources
	Bandwidth AdditionalResource `json:"bandwidth"`
	// Executions additional resources
	Executions AdditionalResource `json:"executions"`
	// Member additional resources
	Member AdditionalResource `json:"member"`
	// Realtime additional resources
	Realtime AdditionalResource `json:"realtime"`
	// Realtime messages additional resources
	RealtimeMessages AdditionalResource `json:"realtimeMessages"`
	// Realtime bandwidth additional resources
	RealtimeBandwidth AdditionalResource `json:"realtimeBandwidth"`
	// Storage additional resources
	Storage AdditionalResource `json:"storage"`
	// User additional resources
	Users AdditionalResource `json:"users"`
	// GBHour additional resources
	GBHours AdditionalResource `json:"GBHours"`
	// Image transformation additional resources
	ImageTransformations AdditionalResource `json:"imageTransformations"`
	// Credits additional resources
	Credits AdditionalResource `json:"credits"`

	// Used by Decode() method
	data []byte
}

func (model UsageBillingPlan) New(data []byte) *UsageBillingPlan {
	model.data = data
	return &model
}

func (model *UsageBillingPlan) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
