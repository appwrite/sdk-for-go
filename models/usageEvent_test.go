package models

import (
    "encoding/json"
    "testing"
)

func TestUsageEventModel(t *testing.T) {
    model := UsageEvent{        Metric: "bandwidth",        Value: 5000,        Time: "2026-04-09T12:00:00.000+00:00",        Path: "/v1/storage/files",        Method: "POST",        Status: "201",        ResourceType: "bucket",        ResourceId: "abc123",        CountryCode: "US",        UserAgent: "AppwriteSDK/1.0",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result UsageEvent
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Metric != model.Metric {
        t.Errorf("Expected Metric %v, got %v", model.Metric, result.Metric)
    }
    if result.Value != model.Value {
        t.Errorf("Expected Value %v, got %v", model.Value, result.Value)
    }
    if result.Time != model.Time {
        t.Errorf("Expected Time %v, got %v", model.Time, result.Time)
    }
    if result.Path != model.Path {
        t.Errorf("Expected Path %v, got %v", model.Path, result.Path)
    }
    if result.Method != model.Method {
        t.Errorf("Expected Method %v, got %v", model.Method, result.Method)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.ResourceType != model.ResourceType {
        t.Errorf("Expected ResourceType %v, got %v", model.ResourceType, result.ResourceType)
    }
    if result.ResourceId != model.ResourceId {
        t.Errorf("Expected ResourceId %v, got %v", model.ResourceId, result.ResourceId)
    }
    if result.CountryCode != model.CountryCode {
        t.Errorf("Expected CountryCode %v, got %v", model.CountryCode, result.CountryCode)
    }
    if result.UserAgent != model.UserAgent {
        t.Errorf("Expected UserAgent %v, got %v", model.UserAgent, result.UserAgent)
    }}
