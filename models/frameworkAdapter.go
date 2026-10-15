package models

import (
	"encoding/json"
	"errors"
)

// FrameworkAdapter Model
type FrameworkAdapter struct {
	// Adapter key.
	Key string `json:"key"`
	// Default command to download dependencies.
	InstallCommand string `json:"installCommand"`
	// Default command to build site into output directory.
	BuildCommand string `json:"buildCommand"`
	// Default output directory of build.
	OutputDirectory string `json:"outputDirectory"`
	// Name of fallback file to use instead of 404 page. If null, Appwrite 404
	// page will be displayed.
	FallbackFile *string `json:"fallbackFile"`

	// Used by Decode() method
	data []byte
}

func (model FrameworkAdapter) New(data []byte) *FrameworkAdapter {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *FrameworkAdapter) UnmarshalJSON(data []byte) error {
	type alias FrameworkAdapter
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Key             json.RawMessage `json:"key"`
		InstallCommand  json.RawMessage `json:"installCommand"`
		BuildCommand    json.RawMessage `json:"buildCommand"`
		OutputDirectory json.RawMessage `json:"outputDirectory"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Key == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Key); err != nil {
			return err
		}
	}
	if present.InstallCommand == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.InstallCommand); err != nil {
			return err
		}
	}
	if present.BuildCommand == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BuildCommand); err != nil {
			return err
		}
	}
	if present.OutputDirectory == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.OutputDirectory); err != nil {
			return err
		}
	}

	*model = FrameworkAdapter(decoded)
	return nil
}

func (model *FrameworkAdapter) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
