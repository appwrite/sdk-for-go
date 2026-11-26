package models

import (
	"encoding/json"
	"errors"
)

// VideoCodec Model
type VideoCodec struct {
	// Codec ID.
	Id string `json:"$id"`
	// Codec display name.
	Name string `json:"name"`
	// Streaming output formats this codec may be packaged into.
	Outputs []string `json:"outputs"`

	// Used by Decode() method
	data []byte
}

func (model VideoCodec) New(data []byte) *VideoCodec {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *VideoCodec) UnmarshalJSON(data []byte) error {
	type alias VideoCodec
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id      json.RawMessage `json:"$id"`
		Name    json.RawMessage `json:"name"`
		Outputs json.RawMessage `json:"outputs"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Outputs == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Outputs); err != nil {
			return err
		}
	}

	*model = VideoCodec(decoded)
	return nil
}

func (model *VideoCodec) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
