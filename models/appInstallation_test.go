package models

import (
    "encoding/json"
    "testing"
)

func TestAppInstallationModel(t *testing.T) {
    model := AppInstallation{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        AppId: "5e5ea5c16897e",        TeamId: "5e5ea5c16897e",        Scopes: []string{"test"},        AuthorizationDetails: map[string]interface{}{},        CreatedById: "5e5ea5c16897e",        CreatedByName: "Walter White",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AppInstallation
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
    if result.AppId != model.AppId {
        t.Errorf("Expected AppId %v, got %v", model.AppId, result.AppId)
    }
    if result.TeamId != model.TeamId {
        t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
    }
    if result.CreatedById != model.CreatedById {
        t.Errorf("Expected CreatedById %v, got %v", model.CreatedById, result.CreatedById)
    }
    if result.CreatedByName != model.CreatedByName {
        t.Errorf("Expected CreatedByName %v, got %v", model.CreatedByName, result.CreatedByName)
    }}
