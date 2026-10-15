package models

import (
	"encoding/json"
	"errors"
)

// Runtime Model
type Runtime struct {
	// Runtime ID.
	Id string `json:"$id"`
	// Parent runtime key.
	Key string `json:"key"`
	// Runtime Name.
	Name string `json:"name"`
	// Runtime version.
	Version string `json:"version"`
	// Base Docker image used to build the runtime.
	Base string `json:"base"`
	// Image name of Docker Hub.
	Image string `json:"image"`
	// Name of the logo image.
	Logo string `json:"logo"`
	// List of supported architectures.
	Supports []string `json:"supports"`

	// Used by Decode() method
	data []byte
}

func (model Runtime) New(data []byte) *Runtime {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Runtime) UnmarshalJSON(data []byte) error {
	type alias Runtime
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id      json.RawMessage `json:"$id"`
		Key     json.RawMessage `json:"key"`
		Name    json.RawMessage `json:"name"`
		Version json.RawMessage `json:"version"`
		Base    json.RawMessage `json:"base"`
		Image   json.RawMessage `json:"image"`
		Logo    json.RawMessage `json:"logo"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
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
	if present.Version == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Version); err != nil {
			return err
		}
	}
	if present.Base == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Base); err != nil {
			return err
		}
	}
	if present.Image == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Image); err != nil {
			return err
		}
	}
	if present.Logo == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Logo); err != nil {
			return err
		}
	}

	*model = Runtime(decoded)
	return nil
}

func (model *Runtime) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
