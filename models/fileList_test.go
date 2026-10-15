package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFileListModel(t *testing.T) {
	model := FileList{Total: 5, Files: []File{File{Id: "5e5ea5c16897e", BucketId: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, Name: "Pink.png", Folder: "photos/2026/", Key: "photos/2026/Pink.png", Signature: "5d529fd02b544198ae075bd57c1762bb", MimeType: "image/png", SizeOriginal: 17890, SizeActual: 12345, ChunksTotal: 17890, ChunksUploaded: 17890, Encryption: true, Compression: "gzip"}}}

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
	}
}

func TestFileListModelDefaults(t *testing.T) {
	var fresh FileList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := FileList{Total: 5, Files: []File{File{Id: "5e5ea5c16897e", BucketId: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, Name: "Pink.png", Folder: "photos/2026/", Key: "photos/2026/Pink.png", Signature: "5d529fd02b544198ae075bd57c1762bb", MimeType: "image/png", SizeOriginal: 17890, SizeActual: 12345, ChunksTotal: 17890, ChunksUploaded: 17890, Encryption: true, Compression: "gzip"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
