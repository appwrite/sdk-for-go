package models

import (
	"encoding/json"
	"errors"
)

// Framework Model
type Framework struct {
	// Framework key.
	Key string `json:"key"`
	// Framework Name.
	Name string `json:"name"`
	// Default runtime version.
	BuildRuntime string `json:"buildRuntime"`
	// List of supported runtime versions.
	Runtimes []string `json:"runtimes"`
	// List of supported adapters.
	Adapters []FrameworkAdapter `json:"adapters"`

	// Used by Decode() method
	data []byte
}

func (model Framework) New(data []byte) *Framework {
	model.data = data
	return &model
}

func (model *Framework) UnmarshalJSON(data []byte) error {
	type alias Framework
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"key\":\"\",\"name\":\"\",\"buildRuntime\":\"\"}"), &fields); err != nil {
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

func (model *Framework) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
