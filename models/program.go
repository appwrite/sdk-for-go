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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Program) UnmarshalJSON(data []byte) error {
	type alias Program
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id            json.RawMessage `json:"$id"`
		Title         json.RawMessage `json:"title"`
		Description   json.RawMessage `json:"description"`
		Tag           json.RawMessage `json:"tag"`
		Icon          json.RawMessage `json:"icon"`
		Url           json.RawMessage `json:"url"`
		Active        json.RawMessage `json:"active"`
		External      json.RawMessage `json:"external"`
		BillingPlanId json.RawMessage `json:"billingPlanId"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Title == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Title); err != nil {
			return err
		}
	}
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.Tag == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Tag); err != nil {
			return err
		}
	}
	if present.Icon == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Icon); err != nil {
			return err
		}
	}
	if present.Url == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Url); err != nil {
			return err
		}
	}
	if present.Active == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Active); err != nil {
			return err
		}
	}
	if present.External == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.External); err != nil {
			return err
		}
	}
	if present.BillingPlanId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BillingPlanId); err != nil {
			return err
		}
	}

	*model = Program(decoded)
	return nil
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
