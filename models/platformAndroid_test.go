package models

import (
    "encoding/json"
    "testing"
)

func TestPlatformAndroidModel(t *testing.T) {
    model := PlatformAndroid{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My Web App",        Type: "web",        ApplicationId: "com.company.appname",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PlatformAndroid
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
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.ApplicationId != model.ApplicationId {
        t.Errorf("Expected ApplicationId %v, got %v", model.ApplicationId, result.ApplicationId)
    }}
