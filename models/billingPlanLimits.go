package models

import (
	"encoding/json"
	"errors"
)

// PlanLimits Model
type BillingPlanLimits struct {
	// Credits limit per billing cycle
	Credits int `json:"credits"`
	// Daily credits limit (if applicable)
	DailyCredits int `json:"dailyCredits"`

	// Used by Decode() method
	data []byte
}

func (model BillingPlanLimits) New(data []byte) *BillingPlanLimits {
	model.data = data
	return &model
}

func (model *BillingPlanLimits) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
