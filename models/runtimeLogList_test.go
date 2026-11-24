package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestRuntimeLogListModel(t *testing.T) {
	model := RuntimeLogList{Total: 5, Logs: []RuntimeLog{RuntimeLog{Id: "5e5ea5c16897e", ResourceType: "functions", ResourceId: "610fc2f985ee0", ResourceInternalId: "42", DeploymentId: "610fc2f985ee0", Region: "fra", Timestamp: "2020-10-15T06:38:00.000+00:00", Stream: "stdout", Message: "Listening on port 3000"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result RuntimeLogList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestRuntimeLogListModelDefaults(t *testing.T) {
	var fresh RuntimeLogList
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

	reused := RuntimeLogList{Total: 5, Logs: []RuntimeLog{RuntimeLog{Id: "5e5ea5c16897e", ResourceType: "functions", ResourceId: "610fc2f985ee0", ResourceInternalId: "42", DeploymentId: "610fc2f985ee0", Region: "fra", Timestamp: "2020-10-15T06:38:00.000+00:00", Stream: "stdout", Message: "Listening on port 3000"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
