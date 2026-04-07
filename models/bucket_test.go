package models

import (
    "encoding/json"
    "testing"
)

func TestBucketModel(t *testing.T) {
    model := Bucket{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        FileSecurity: true,        Name: "Documents",        Enabled: true,        MaximumFileSize: 100,        AllowedFileExtensions: []string{"test"},        Compression: "gzip",        Encryption: true,        Antivirus: true,        Transformations: true,        TotalSize: 128,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Bucket
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
    if result.FileSecurity != model.FileSecurity {
        t.Errorf("Expected FileSecurity %v, got %v", model.FileSecurity, result.FileSecurity)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }
    if result.MaximumFileSize != model.MaximumFileSize {
        t.Errorf("Expected MaximumFileSize %v, got %v", model.MaximumFileSize, result.MaximumFileSize)
    }
    if result.Compression != model.Compression {
        t.Errorf("Expected Compression %v, got %v", model.Compression, result.Compression)
    }
    if result.Encryption != model.Encryption {
        t.Errorf("Expected Encryption %v, got %v", model.Encryption, result.Encryption)
    }
    if result.Antivirus != model.Antivirus {
        t.Errorf("Expected Antivirus %v, got %v", model.Antivirus, result.Antivirus)
    }
    if result.Transformations != model.Transformations {
        t.Errorf("Expected Transformations %v, got %v", model.Transformations, result.Transformations)
    }
    if result.TotalSize != model.TotalSize {
        t.Errorf("Expected TotalSize %v, got %v", model.TotalSize, result.TotalSize)
    }}
