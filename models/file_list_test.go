package models

import (
    "encoding/json"
    "testing"
)

func TestFileListModel(t *testing.T) {
    model := FileList{        Total: 5,        Files: []File{File{        Id: "5e5ea5c16897e",        BucketId: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        Name: "Pink.png",        Signature: "5d529fd02b544198ae075bd57c1762bb",        MimeType: "image/png",        SizeOriginal: 17890,        ChunksTotal: 17890,        ChunksUploaded: 17890,        Encryption: true,        Compression: "gzip",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result FileList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
