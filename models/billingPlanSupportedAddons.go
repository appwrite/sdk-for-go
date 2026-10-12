package models

import (
	"encoding/json"
	"errors"
)

// BillingPlanSupportedAddons Model
type BillingPlanSupportedAddons struct {
	// Whether the plan supports BAA (Business Associate Agreement) addon
	Baa bool `json:"baa"`
	// Whether the plan supports Premium Geo DB addon (project-level)
	PremiumGeoDB bool `json:"premiumGeoDB"`
	// Whether the plan supports Premium Geo DB addon (organization-level)
	PremiumGeoDBOrg bool `json:"premiumGeoDBOrg"`

	// Used by Decode() method
	data []byte
}

func (model BillingPlanSupportedAddons) New(data []byte) *BillingPlanSupportedAddons {
	model.data = data
	return &model
}

func (model *BillingPlanSupportedAddons) UnmarshalJSON(data []byte) error {
	type alias BillingPlanSupportedAddons
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"baa\":false,\"premiumGeoDB\":false,\"premiumGeoDBOrg\":false}"), &fields); err != nil {
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

func (model *BillingPlanSupportedAddons) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
