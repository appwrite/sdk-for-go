package models

import (
	"encoding/json"
	"errors"
)

// Embedding Model
type Embedding struct {
	// Embedding model used to generate embeddings.
	Model string `json:"model"`
	// Number of dimensions for each embedding vector.
	Dimension int `json:"dimension"`
	// Embedding vector values. If an error occurs, this will be an empty array.
	Embedding []float64 `json:"embedding"`
	// Error message if embedding generation fails. Empty string if no error.
	Error string `json:"error"`

	// Used by Decode() method
	data []byte
}

func (model Embedding) New(data []byte) *Embedding {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Embedding) UnmarshalJSON(data []byte) error {
	type alias Embedding
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Model     json.RawMessage `json:"model"`
		Dimension json.RawMessage `json:"dimension"`
		Embedding json.RawMessage `json:"embedding"`
		Error     json.RawMessage `json:"error"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Model == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Model); err != nil {
			return err
		}
	}
	if present.Dimension == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Dimension); err != nil {
			return err
		}
	}
	if present.Embedding == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Embedding); err != nil {
			return err
		}
	}
	if present.Error == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Error); err != nil {
			return err
		}
	}

	*model = Embedding(decoded)
	return nil
}

func (model *Embedding) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
