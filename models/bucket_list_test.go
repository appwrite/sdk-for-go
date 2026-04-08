package models

import (
    "encoding/json"
    "testing"
)

func TestBucketListModel(t *testing.T) {
    model := BucketList{        Total: 5,        Buckets: []Bucket{Bucket{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        FileSecurity: true,        Name: "Documents",        Enabled: true,        MaximumFileSize: 100,        AllowedFileExtensions: []string{"test"},        Compression: "gzip",        Encryption: true,        Antivirus: true,        Transformations: true,        TotalSize: 128,    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BucketList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
