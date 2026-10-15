package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTopicListModel(t *testing.T) {
	model := TopicList{Total: 5, Topics: []Topic{Topic{Id: "259125845563242502", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "events", EmailTotal: 100, SmsTotal: 100, PushTotal: 100, Subscribe: []string{"test"}}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result TopicList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestTopicListModelDefaults(t *testing.T) {
	var fresh TopicList
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

	reused := TopicList{Total: 5, Topics: []Topic{Topic{Id: "259125845563242502", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "events", EmailTotal: 100, SmsTotal: 100, PushTotal: 100, Subscribe: []string{"test"}}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
