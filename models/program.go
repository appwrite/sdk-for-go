package models

import (
	"encoding/json"
	"errors"
)

// Program Model
type Program struct {
	// Program ID
	Id string `json:"$id"`
	// Program title
	Title string `json:"title"`
	// Program description
	Description string `json:"description"`
	// Program tag for highlighting on console
	Tag string `json:"tag"`
	// Program icon for highlighting on console
	Icon string `json:"icon"`
	// URL for more information on this program
	Url string `json:"url"`
	// Whether this program is active
	Active bool `json:"active"`
	// Whether this program is external
	External bool `json:"external"`
	// Billing plan ID that this is program is associated with.
	BillingPlanId string `json:"billingPlanId"`

	// Used by Decode() method
	data []byte
}

func (model Program) New(data []byte) *Program {
	model.data = data
	return &model
}

func (model *Program) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
