package models

import (
	"encoding/json"
	"errors"
)

// Addon Model
type BillingPlanAddon struct {
	// Addon seats
	Seats BillingPlanAddonDetails `json:"seats"`
	// Addon projects
	Projects BillingPlanAddonDetails `json:"projects"`

	// Used by Decode() method
	data []byte
}

func (model BillingPlanAddon) New(data []byte) *BillingPlanAddon {
	model.data = data
	return &model
}

func (model *BillingPlanAddon) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
