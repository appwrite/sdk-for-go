package models

import (
    "encoding/json"
    "testing"
)

func TestExecutionModel(t *testing.T) {
    model := Execution{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        FunctionId: "5e5ea6g16897e",        DeploymentId: "5e5ea5c16897e",        Trigger: "http",        Status: "processing",        RequestMethod: "GET",        RequestPath: "/articles?id=5",        RequestHeaders: []Headers{Headers{        Name: "Content-Type",        Value: "application/json",    },
            },        ResponseStatusCode: 200,        ResponseBody: "string",        ResponseHeaders: []Headers{Headers{        Name: "Content-Type",        Value: "application/json",    },
            },        Logs: "string",        Errors: "string",        Duration: 0.4,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Execution
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
    if result.FunctionId != model.FunctionId {
        t.Errorf("Expected FunctionId %v, got %v", model.FunctionId, result.FunctionId)
    }
    if result.DeploymentId != model.DeploymentId {
        t.Errorf("Expected DeploymentId %v, got %v", model.DeploymentId, result.DeploymentId)
    }
    if result.Trigger != model.Trigger {
        t.Errorf("Expected Trigger %v, got %v", model.Trigger, result.Trigger)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.RequestMethod != model.RequestMethod {
        t.Errorf("Expected RequestMethod %v, got %v", model.RequestMethod, result.RequestMethod)
    }
    if result.RequestPath != model.RequestPath {
        t.Errorf("Expected RequestPath %v, got %v", model.RequestPath, result.RequestPath)
    }
    if result.ResponseStatusCode != model.ResponseStatusCode {
        t.Errorf("Expected ResponseStatusCode %v, got %v", model.ResponseStatusCode, result.ResponseStatusCode)
    }
    if result.ResponseBody != model.ResponseBody {
        t.Errorf("Expected ResponseBody %v, got %v", model.ResponseBody, result.ResponseBody)
    }
    if result.Logs != model.Logs {
        t.Errorf("Expected Logs %v, got %v", model.Logs, result.Logs)
    }
    if result.Errors != model.Errors {
        t.Errorf("Expected Errors %v, got %v", model.Errors, result.Errors)
    }
    if result.Duration != model.Duration {
        t.Errorf("Expected Duration %v, got %v", model.Duration, result.Duration)
    }}
