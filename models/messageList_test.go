package models

import (
    "encoding/json"
    "testing"
)

func TestMessageListModel(t *testing.T) {
    model := MessageList{        Total: 5,        Messages: []Message{Message{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        ProviderType: "email",        Topics: []string{"test"},        Users: []string{"test"},        Targets: []string{"test"},        DeliveredTotal: 1,        Data: map[string]interface{}{},        Status: "processing",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result MessageList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
