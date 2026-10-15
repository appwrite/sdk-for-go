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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *BillingPlanSupportedAddons) UnmarshalJSON(data []byte) error {
	type alias BillingPlanSupportedAddons
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Baa             json.RawMessage `json:"baa"`
		PremiumGeoDB    json.RawMessage `json:"premiumGeoDB"`
		PremiumGeoDBOrg json.RawMessage `json:"premiumGeoDBOrg"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Baa == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Baa); err != nil {
			return err
		}
	}
	if present.PremiumGeoDB == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.PremiumGeoDB); err != nil {
			return err
		}
	}
	if present.PremiumGeoDBOrg == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.PremiumGeoDBOrg); err != nil {
			return err
		}
	}

	*model = BillingPlanSupportedAddons(decoded)
	return nil
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
