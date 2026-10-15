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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Framework) UnmarshalJSON(data []byte) error {
	type alias Framework
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Key          json.RawMessage `json:"key"`
		Name         json.RawMessage `json:"name"`
		BuildRuntime json.RawMessage `json:"buildRuntime"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Key == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Key); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.BuildRuntime == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BuildRuntime); err != nil {
			return err
		}
	}

	*model = Framework(decoded)
	return nil
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
