package models

import (
	"encoding/json"
	"reflect"
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

func TestEmbeddingListModelDefaults(t *testing.T) {
	var fresh EmbeddingList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := EmbeddingList{Total: 5, Embeddings: []Embedding{Embedding{Model: "nomic-embed-text", Dimension: 768, Embedding: []float64{}, Error: "Error message"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
