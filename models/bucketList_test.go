package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBucketListModel(t *testing.T) {
	model := BucketList{Total: 5, Buckets: []Bucket{Bucket{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, FileSecurity: true, Name: "Documents", Description: "Invoices and receipts uploaded by customers.", Enabled: true, MaximumFileSize: 100, AllowedFileExtensions: []string{"test"}, Compression: "gzip", Encryption: true, Antivirus: true, Transformations: true, StripMetadata: true, TotalSize: 128}}}

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
	}
}

func TestBucketListModelDefaults(t *testing.T) {
	var fresh BucketList
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

	reused := BucketList{Total: 5, Buckets: []Bucket{Bucket{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, FileSecurity: true, Name: "Documents", Description: "Invoices and receipts uploaded by customers.", Enabled: true, MaximumFileSize: 100, AllowedFileExtensions: []string{"test"}, Compression: "gzip", Encryption: true, Antivirus: true, Transformations: true, StripMetadata: true, TotalSize: 128}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
