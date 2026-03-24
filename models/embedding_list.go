package models

import (
    "encoding/json"
    "errors"
)

// EmbeddingList Model
type EmbeddingList struct {
    // Total number of embeddings that matched your query.
    Total int `json:"total"`
    // List of embeddings.
    Embeddings []Embedding `json:"embeddings"`

    // Used by Decode() method
    data []byte
}

func (model EmbeddingList) New(data []byte) *EmbeddingList {
    model.data = data
    return &model
}

func (model *EmbeddingList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}