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