package models

import (
    "encoding/json"
    "testing"
)

func TestTopicModel(t *testing.T) {
    model := Topic{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "events",        EmailTotal: 100,        SmsTotal: 100,        PushTotal: 100,        Subscribe: []string{"test"},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Topic
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.EmailTotal != model.EmailTotal {
        t.Errorf("Expected EmailTotal %v, got %v", model.EmailTotal, result.EmailTotal)
    }
    if result.SmsTotal != model.SmsTotal {
        t.Errorf("Expected SmsTotal %v, got %v", model.SmsTotal, result.SmsTotal)
    }
    if result.PushTotal != model.PushTotal {
        t.Errorf("Expected PushTotal %v, got %v", model.PushTotal, result.PushTotal)
    }}
