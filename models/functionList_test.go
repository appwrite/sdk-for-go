package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFunctionListModel(t *testing.T) {
	model := FunctionList{Total: 5, Functions: []Function{Function{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Execute: []string{"test"}, Name: "My Function", Description: "Sends a welcome email to newly registered users.", Enabled: true, Live: true, Logging: true, Runtime: "python-3.8", DeploymentRetention: 7, DeploymentId: "5e5ea5c16897e", DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentId: "5e5ea5c16897e", LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentStatus: "ready", Scopes: []string{"test"}, Vars: []Variable{Variable{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Key: "API_KEY", Description: "Stripe secret key used at checkout.", Value: "myPa$$word1", Secret: true, ResourceType: "function", ResourceId: "myAwesomeFunction"}}, Events: []string{"test"}, Schedule: "5 4 * * *", Timeout: 300, Entrypoint: "index.js", Commands: "npm install", Version: "v2", InstallationId: "6m40at4ejk5h2u9s1hboo", ProviderRepositoryId: "appwrite", ProviderBranch: "main", ProviderRootDirectory: "functions/helloWorld", ProviderSilentMode: true, ProviderBranches: []string{"test"}, ProviderPaths: []string{"test"}, BuildSpecification: "s-1vcpu-512mb", RuntimeSpecification: "s-1vcpu-512mb", EmailEnabled: true}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result FunctionList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestFunctionListModelDefaults(t *testing.T) {
	var fresh FunctionList
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

	reused := FunctionList{Total: 5, Functions: []Function{Function{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Execute: []string{"test"}, Name: "My Function", Description: "Sends a welcome email to newly registered users.", Enabled: true, Live: true, Logging: true, Runtime: "python-3.8", DeploymentRetention: 7, DeploymentId: "5e5ea5c16897e", DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentId: "5e5ea5c16897e", LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentStatus: "ready", Scopes: []string{"test"}, Vars: []Variable{Variable{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Key: "API_KEY", Description: "Stripe secret key used at checkout.", Value: "myPa$$word1", Secret: true, ResourceType: "function", ResourceId: "myAwesomeFunction"}}, Events: []string{"test"}, Schedule: "5 4 * * *", Timeout: 300, Entrypoint: "index.js", Commands: "npm install", Version: "v2", InstallationId: "6m40at4ejk5h2u9s1hboo", ProviderRepositoryId: "appwrite", ProviderBranch: "main", ProviderRootDirectory: "functions/helloWorld", ProviderSilentMode: true, ProviderBranches: []string{"test"}, ProviderPaths: []string{"test"}, BuildSpecification: "s-1vcpu-512mb", RuntimeSpecification: "s-1vcpu-512mb", EmailEnabled: true}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
