package models

import (
	"encoding/json"
	"testing"
)

func TestEmbeddingListModel(t *testing.T) {
	model := EmbeddingList{Total: 5, Embeddings: []Embedding{Embedding{Model: "nomic-embed-text", Dimension: 768, Embedding: []float64{}, Error: "Error message"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result EmbeddingList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
