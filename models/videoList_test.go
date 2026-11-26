package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestVideoListModel(t *testing.T) {
	model := VideoList{Total: 5, Videos: []Video{Video{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, BucketId: "5e5ea5c16897e", FileId: "d5fg5ehg1c168g7c", Name: "Product demo", Size: 23647142, Format: "MPEG-4", Duration: 92810, Width: 960, Height: 544, AspectRatio: "16:9", VideoCodec: "h264", VideoFormat: "AVC", VideoFormatProfile: "Baseline@L3.1", VideoBitRate: 564790, VideoFrameRate: 29.97, VideoFrameRateMode: "CFR", AudioCodec: "aac", AudioFormat: "AAC", AudioBitRate: 127999, AudioSampleRate: 48000, TimelineStatus: "ready", TimelineError: "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result VideoList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestVideoListModelDefaults(t *testing.T) {
	var fresh VideoList
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

	reused := VideoList{Total: 5, Videos: []Video{Video{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, BucketId: "5e5ea5c16897e", FileId: "d5fg5ehg1c168g7c", Name: "Product demo", Size: 23647142, Format: "MPEG-4", Duration: 92810, Width: 960, Height: 544, AspectRatio: "16:9", VideoCodec: "h264", VideoFormat: "AVC", VideoFormatProfile: "Baseline@L3.1", VideoBitRate: 564790, VideoFrameRate: 29.97, VideoFrameRateMode: "CFR", AudioCodec: "aac", AudioFormat: "AAC", AudioBitRate: 127999, AudioSampleRate: 48000, TimelineStatus: "ready", TimelineError: "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
