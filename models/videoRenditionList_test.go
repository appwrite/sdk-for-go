package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestVideoRenditionListModel(t *testing.T) {
	model := VideoRenditionList{Total: 5, Renditions: []VideoRendition{VideoRendition{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", VideoId: "5e5ea5c16897e", ProfileId: "d5fg5ehg1c168g7c", Name: "1280X720@3679", StartedAt: "2020-10-15T06:38:00.000+00:00", EndedAt: "2020-10-15T06:38:00.000+00:00", Width: 1280, Height: 720, VideoBitRate: 3551000, AudioBitRate: 128000, TargetDuration: 6.006, Status: "ready", Error: "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)", Progress: 88, Output: "hls", Codec: "h264", Size: 18350080}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result VideoRenditionList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestVideoRenditionListModelDefaults(t *testing.T) {
	var fresh VideoRenditionList
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

	reused := VideoRenditionList{Total: 5, Renditions: []VideoRendition{VideoRendition{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", VideoId: "5e5ea5c16897e", ProfileId: "d5fg5ehg1c168g7c", Name: "1280X720@3679", StartedAt: "2020-10-15T06:38:00.000+00:00", EndedAt: "2020-10-15T06:38:00.000+00:00", Width: 1280, Height: 720, VideoBitRate: 3551000, AudioBitRate: 128000, TargetDuration: 6.006, Status: "ready", Error: "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)", Progress: 88, Output: "hls", Codec: "h264", Size: 18350080}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
