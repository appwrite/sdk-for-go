package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFunctionModel(t *testing.T) {
	model := Function{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Execute: []string{"test"}, Name: "My Function", Enabled: true, Live: true, Logging: true, Runtime: "python-3.8", DeploymentRetention: 7, DeploymentId: "5e5ea5c16897e", DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentId: "5e5ea5c16897e", LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentStatus: "ready", Scopes: []string{"test"}, Vars: []Variable{Variable{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Key: "API_KEY", Value: "myPa$$word1", Secret: true, ResourceType: "function", ResourceId: "myAwesomeFunction"}}, Events: []string{"test"}, Schedule: "5 4 * * *", Timeout: 300, Entrypoint: "index.js", Commands: "npm install", Version: "v2", InstallationId: "6m40at4ejk5h2u9s1hboo", ProviderRepositoryId: "appwrite", ProviderBranch: "main", ProviderRootDirectory: "functions/helloWorld", ProviderSilentMode: true, ProviderBranches: []string{"test"}, ProviderPaths: []string{"test"}, BuildSpecification: "s-1vcpu-512mb", RuntimeSpecification: "s-1vcpu-512mb", EmailEnabled: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Function
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.Enabled != model.Enabled {
		t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
	}
	if result.Live != model.Live {
		t.Errorf("Expected Live %v, got %v", model.Live, result.Live)
	}
	if result.Logging != model.Logging {
		t.Errorf("Expected Logging %v, got %v", model.Logging, result.Logging)
	}
	if result.Runtime != model.Runtime {
		t.Errorf("Expected Runtime %v, got %v", model.Runtime, result.Runtime)
	}
	if result.DeploymentRetention != model.DeploymentRetention {
		t.Errorf("Expected DeploymentRetention %v, got %v", model.DeploymentRetention, result.DeploymentRetention)
	}
	if result.DeploymentId != model.DeploymentId {
		t.Errorf("Expected DeploymentId %v, got %v", model.DeploymentId, result.DeploymentId)
	}
	if result.DeploymentCreatedAt != model.DeploymentCreatedAt {
		t.Errorf("Expected DeploymentCreatedAt %v, got %v", model.DeploymentCreatedAt, result.DeploymentCreatedAt)
	}
	if result.LatestDeploymentId != model.LatestDeploymentId {
		t.Errorf("Expected LatestDeploymentId %v, got %v", model.LatestDeploymentId, result.LatestDeploymentId)
	}
	if result.LatestDeploymentCreatedAt != model.LatestDeploymentCreatedAt {
		t.Errorf("Expected LatestDeploymentCreatedAt %v, got %v", model.LatestDeploymentCreatedAt, result.LatestDeploymentCreatedAt)
	}
	if result.LatestDeploymentStatus != model.LatestDeploymentStatus {
		t.Errorf("Expected LatestDeploymentStatus %v, got %v", model.LatestDeploymentStatus, result.LatestDeploymentStatus)
	}
	if result.Schedule != model.Schedule {
		t.Errorf("Expected Schedule %v, got %v", model.Schedule, result.Schedule)
	}
	if result.Timeout != model.Timeout {
		t.Errorf("Expected Timeout %v, got %v", model.Timeout, result.Timeout)
	}
	if result.Entrypoint != model.Entrypoint {
		t.Errorf("Expected Entrypoint %v, got %v", model.Entrypoint, result.Entrypoint)
	}
	if result.Commands != model.Commands {
		t.Errorf("Expected Commands %v, got %v", model.Commands, result.Commands)
	}
	if result.Version != model.Version {
		t.Errorf("Expected Version %v, got %v", model.Version, result.Version)
	}
	if result.InstallationId != model.InstallationId {
		t.Errorf("Expected InstallationId %v, got %v", model.InstallationId, result.InstallationId)
	}
	if result.ProviderRepositoryId != model.ProviderRepositoryId {
		t.Errorf("Expected ProviderRepositoryId %v, got %v", model.ProviderRepositoryId, result.ProviderRepositoryId)
	}
	if result.ProviderBranch != model.ProviderBranch {
		t.Errorf("Expected ProviderBranch %v, got %v", model.ProviderBranch, result.ProviderBranch)
	}
	if result.ProviderRootDirectory != model.ProviderRootDirectory {
		t.Errorf("Expected ProviderRootDirectory %v, got %v", model.ProviderRootDirectory, result.ProviderRootDirectory)
	}
	if result.ProviderSilentMode != model.ProviderSilentMode {
		t.Errorf("Expected ProviderSilentMode %v, got %v", model.ProviderSilentMode, result.ProviderSilentMode)
	}
	if result.BuildSpecification != model.BuildSpecification {
		t.Errorf("Expected BuildSpecification %v, got %v", model.BuildSpecification, result.BuildSpecification)
	}
	if result.RuntimeSpecification != model.RuntimeSpecification {
		t.Errorf("Expected RuntimeSpecification %v, got %v", model.RuntimeSpecification, result.RuntimeSpecification)
	}
	if result.EmailEnabled != model.EmailEnabled {
		t.Errorf("Expected EmailEnabled %v, got %v", model.EmailEnabled, result.EmailEnabled)
	}
}

func TestFunctionModelDefaults(t *testing.T) {
	var fresh Function
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Id)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $id to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CreatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $createdAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.UpdatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $updatedAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Execute)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing execute to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Name)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing name to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Enabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing enabled to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Live)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing live to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Logging)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing logging to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Runtime)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing runtime to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DeploymentRetention)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing deploymentRetention to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DeploymentId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing deploymentId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DeploymentCreatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing deploymentCreatedAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LatestDeploymentId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing latestDeploymentId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LatestDeploymentCreatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing latestDeploymentCreatedAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LatestDeploymentStatus)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing latestDeploymentStatus to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Scopes)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing scopes to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Events)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing events to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Schedule)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing schedule to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("15"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Timeout)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing timeout to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Entrypoint)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing entrypoint to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Commands)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing commands to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"v5\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Version)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing version to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.InstallationId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing installationId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ProviderRepositoryId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing providerRepositoryId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ProviderBranch)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing providerBranch to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ProviderRootDirectory)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing providerRootDirectory to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ProviderSilentMode)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing providerSilentMode to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ProviderBranches)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing providerBranches to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ProviderPaths)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing providerPaths to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"s-1vcpu-512mb\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BuildSpecification)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing buildSpecification to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"s-1vcpu-512mb\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RuntimeSpecification)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing runtimeSpecification to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.EmailEnabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing emailEnabled to default to %v, got %v", want, got)
		}
	}

	reused := Function{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Execute: []string{"test"}, Name: "My Function", Enabled: true, Live: true, Logging: true, Runtime: "python-3.8", DeploymentRetention: 7, DeploymentId: "5e5ea5c16897e", DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentId: "5e5ea5c16897e", LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentStatus: "ready", Scopes: []string{"test"}, Vars: []Variable{Variable{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Key: "API_KEY", Value: "myPa$$word1", Secret: true, ResourceType: "function", ResourceId: "myAwesomeFunction"}}, Events: []string{"test"}, Schedule: "5 4 * * *", Timeout: 300, Entrypoint: "index.js", Commands: "npm install", Version: "v2", InstallationId: "6m40at4ejk5h2u9s1hboo", ProviderRepositoryId: "appwrite", ProviderBranch: "main", ProviderRootDirectory: "functions/helloWorld", ProviderSilentMode: true, ProviderBranches: []string{"test"}, ProviderPaths: []string{"test"}, BuildSpecification: "s-1vcpu-512mb", RuntimeSpecification: "s-1vcpu-512mb", EmailEnabled: true}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
