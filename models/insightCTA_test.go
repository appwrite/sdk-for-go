package models

import (
    "encoding/json"
    "testing"
)

func TestInsightCTAModel(t *testing.T) {
    model := InsightCTA{        Label: "Create missing index",        Service: "tablesDB",        Method: "createIndex",        Params: map[string]interface{}{},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result InsightCTA
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Label != model.Label {
        t.Errorf("Expected Label %v, got %v", model.Label, result.Label)
    }
    if result.Service != model.Service {
        t.Errorf("Expected Service %v, got %v", model.Service, result.Service)
    }
    if result.Method != model.Method {
        t.Errorf("Expected Method %v, got %v", model.Method, result.Method)
    }}
