package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestVideoCaptionListModel(t *testing.T) {
	model := VideoCaptionList{Total: 5, Captions: []VideoCaption{VideoCaption{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", VideoId: "5e5ea5c16897e", BucketId: "d5fg5ehg1c168g7c", FileId: "c5fg5emg1c168grr", Name: "English", Language: "eng", Default: true, Embedded: true, TargetDuration: 93, Status: "ready", Error: "Caption is not a proper .srt file (Invalid timestamp delimiter at line 2)", Size: 4096}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result VideoCaptionList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestVideoCaptionListModelDefaults(t *testing.T) {
	var fresh VideoCaptionList
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

	reused := VideoCaptionList{Total: 5, Captions: []VideoCaption{VideoCaption{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", VideoId: "5e5ea5c16897e", BucketId: "d5fg5ehg1c168g7c", FileId: "c5fg5emg1c168grr", Name: "English", Language: "eng", Default: true, Embedded: true, TargetDuration: 93, Status: "ready", Error: "Caption is not a proper .srt file (Invalid timestamp delimiter at line 2)", Size: 4096}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
