package models

import (
    "encoding/json"
    "testing"
)

func TestEmbeddingModel(t *testing.T) {
    model := Embedding{        Model: "nomic-embed-text",        Dimension: 768,        Embedding: []float64{},        Error: "Error message",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Embedding
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Model != model.Model {
        t.Errorf("Expected Model %v, got %v", model.Model, result.Model)
    }
    if result.Dimension != model.Dimension {
        t.Errorf("Expected Dimension %v, got %v", model.Dimension, result.Dimension)
    }
    if result.Error != model.Error {
        t.Errorf("Expected Error %v, got %v", model.Error, result.Error)
    }}
