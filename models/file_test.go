package models

import (
    "encoding/json"
    "testing"
)

func TestFileModel(t *testing.T) {
    model := File{        Id: "5e5ea5c16897e",        BucketId: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        Name: "Pink.png",        Signature: "5d529fd02b544198ae075bd57c1762bb",        MimeType: "image/png",        SizeOriginal: 17890,        ChunksTotal: 17890,        ChunksUploaded: 17890,        Encryption: true,        Compression: "gzip",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result File
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.BucketId != model.BucketId {
        t.Errorf("Expected BucketId %v, got %v", model.BucketId, result.BucketId)
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
    if result.Signature != model.Signature {
        t.Errorf("Expected Signature %v, got %v", model.Signature, result.Signature)
    }
    if result.MimeType != model.MimeType {
        t.Errorf("Expected MimeType %v, got %v", model.MimeType, result.MimeType)
    }
    if result.SizeOriginal != model.SizeOriginal {
        t.Errorf("Expected SizeOriginal %v, got %v", model.SizeOriginal, result.SizeOriginal)
    }
    if result.ChunksTotal != model.ChunksTotal {
        t.Errorf("Expected ChunksTotal %v, got %v", model.ChunksTotal, result.ChunksTotal)
    }
    if result.ChunksUploaded != model.ChunksUploaded {
        t.Errorf("Expected ChunksUploaded %v, got %v", model.ChunksUploaded, result.ChunksUploaded)
    }
    if result.Encryption != model.Encryption {
        t.Errorf("Expected Encryption %v, got %v", model.Encryption, result.Encryption)
    }
    if result.Compression != model.Compression {
        t.Errorf("Expected Compression %v, got %v", model.Compression, result.Compression)
    }}
