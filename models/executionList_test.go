package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestExecutionListModel(t *testing.T) {
	model := ExecutionList{Total: 5, Executions: []Execution{Execution{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, ResourceId: "5e5ea6g16897e", ResourceType: "functions", DeploymentId: "5e5ea5c16897e", Trigger: "http", Status: "processing", RequestMethod: "GET", RequestPath: "/articles?id=5", RequestHeaders: []Headers{Headers{Name: "Content-Type", Value: "application/json"}}, ResponseStatusCode: 200, ResponseBody: "string", ResponseHeaders: []Headers{Headers{Name: "Content-Type", Value: "application/json"}}, Logs: "string", Errors: "string", Duration: 1.23}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result ExecutionList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestExecutionListModelDefaults(t *testing.T) {
	var fresh ExecutionList
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

	reused := ExecutionList{Total: 5, Executions: []Execution{Execution{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, ResourceId: "5e5ea6g16897e", ResourceType: "functions", DeploymentId: "5e5ea5c16897e", Trigger: "http", Status: "processing", RequestMethod: "GET", RequestPath: "/articles?id=5", RequestHeaders: []Headers{Headers{Name: "Content-Type", Value: "application/json"}}, ResponseStatusCode: 200, ResponseBody: "string", ResponseHeaders: []Headers{Headers{Name: "Content-Type", Value: "application/json"}}, Logs: "string", Errors: "string", Duration: 1.23}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
