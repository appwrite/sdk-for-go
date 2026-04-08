package models

import (
    "encoding/json"
    "testing"
)

func TestDatabaseModel(t *testing.T) {
    model := Database{        Id: "5e5ea5c16897e",        Name: "My Database",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Enabled: true,        Type: "legacy",        Policies: []Index{Index{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "index1",        Type: "primary",        Status: "available",        Error: "string",        Attributes: []string{"test"},        Lengths: []int{1},    },
            },        Archives: []Collection{Collection{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        DatabaseId: "5e5ea5c16897e",        Name: "My Collection",        Enabled: true,        DocumentSecurity: true,        Attributes: []map[string]any{},        Indexes: []Index{Index{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "index1",        Type: "primary",        Status: "available",        Error: "string",        Attributes: []string{"test"},        Lengths: []int{1},    },
            },        BytesMax: 65535,        BytesUsed: 1500,    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Database
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }}
