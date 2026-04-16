package models

import (
    "encoding/json"
    "testing"
)

func TestAttributeListModel(t *testing.T) {
    model := AttributeList{        Total: 5,        Attributes: []map[string]any{},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AttributeList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
