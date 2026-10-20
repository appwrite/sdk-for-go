package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSiteModel(t *testing.T) {
	model := Site{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "My Site", Description: "Marketing site for the spring launch.", Enabled: true, Live: true, Logging: true, Framework: "react", DeploymentRetention: 7, DeploymentId: "5e5ea5c16897e", DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", DeploymentScreenshotLight: "5e5ea5c16897e", DeploymentScreenshotDark: "5e5ea5c16897e", LatestDeploymentId: "5e5ea5c16897e", LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentStatus: "ready", Scopes: []string{"test"}, Vars: []Variable{Variable{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Key: "API_KEY", Description: "Stripe secret key used at checkout.", Value: "myPa$$word1", Secret: true, ResourceType: "function", ResourceId: "myAwesomeFunction"}}, Timeout: 300, InstallCommand: "npm install", BuildCommand: "npm run build", StartCommand: "node custom-server.mjs", OutputDirectory: "build", InstallationId: "6m40at4ejk5h2u9s1hboo", ProviderRepositoryId: "appwrite", ProviderBranch: "main", ProviderRootDirectory: "sites/helloWorld", ProviderSilentMode: true, ProviderBranches: []string{"test"}, ProviderPaths: []string{"test"}, BuildSpecification: "s-1vcpu-512mb", RuntimeSpecification: "s-1vcpu-512mb", BuildRuntime: "node-22", Adapter: "static", FallbackFile: "index.html"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Site
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
	if result.Description != model.Description {
		t.Errorf("Expected Description %v, got %v", model.Description, result.Description)
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
	if result.Framework != model.Framework {
		t.Errorf("Expected Framework %v, got %v", model.Framework, result.Framework)
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
	if result.DeploymentScreenshotLight != model.DeploymentScreenshotLight {
		t.Errorf("Expected DeploymentScreenshotLight %v, got %v", model.DeploymentScreenshotLight, result.DeploymentScreenshotLight)
	}
	if result.DeploymentScreenshotDark != model.DeploymentScreenshotDark {
		t.Errorf("Expected DeploymentScreenshotDark %v, got %v", model.DeploymentScreenshotDark, result.DeploymentScreenshotDark)
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
	if result.Timeout != model.Timeout {
		t.Errorf("Expected Timeout %v, got %v", model.Timeout, result.Timeout)
	}
	if result.InstallCommand != model.InstallCommand {
		t.Errorf("Expected InstallCommand %v, got %v", model.InstallCommand, result.InstallCommand)
	}
	if result.BuildCommand != model.BuildCommand {
		t.Errorf("Expected BuildCommand %v, got %v", model.BuildCommand, result.BuildCommand)
	}
	if result.StartCommand != model.StartCommand {
		t.Errorf("Expected StartCommand %v, got %v", model.StartCommand, result.StartCommand)
	}
	if result.OutputDirectory != model.OutputDirectory {
		t.Errorf("Expected OutputDirectory %v, got %v", model.OutputDirectory, result.OutputDirectory)
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
	if result.BuildRuntime != model.BuildRuntime {
		t.Errorf("Expected BuildRuntime %v, got %v", model.BuildRuntime, result.BuildRuntime)
	}
	if result.Adapter != model.Adapter {
		t.Errorf("Expected Adapter %v, got %v", model.Adapter, result.Adapter)
	}
	if result.FallbackFile != model.FallbackFile {
		t.Errorf("Expected FallbackFile %v, got %v", model.FallbackFile, result.FallbackFile)
	}
}

func TestSiteModelDefaults(t *testing.T) {
	var fresh Site
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
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Description)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing description to default to %v, got %v", want, got)
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
		encoded, err := json.Marshal(fresh.Framework)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing framework to default to %v, got %v", want, got)
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
		encoded, err := json.Marshal(fresh.DeploymentScreenshotLight)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing deploymentScreenshotLight to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DeploymentScreenshotDark)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing deploymentScreenshotDark to default to %v, got %v", want, got)
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
		encoded, err := json.Marshal(fresh.InstallCommand)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing installCommand to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BuildCommand)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing buildCommand to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.StartCommand)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing startCommand to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.OutputDirectory)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing outputDirectory to default to %v, got %v", want, got)
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
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BuildRuntime)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing buildRuntime to default to %v, got %v", want, got)
		}
	}

	reused := Site{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "My Site", Description: "Marketing site for the spring launch.", Enabled: true, Live: true, Logging: true, Framework: "react", DeploymentRetention: 7, DeploymentId: "5e5ea5c16897e", DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", DeploymentScreenshotLight: "5e5ea5c16897e", DeploymentScreenshotDark: "5e5ea5c16897e", LatestDeploymentId: "5e5ea5c16897e", LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00", LatestDeploymentStatus: "ready", Scopes: []string{"test"}, Vars: []Variable{Variable{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Key: "API_KEY", Description: "Stripe secret key used at checkout.", Value: "myPa$$word1", Secret: true, ResourceType: "function", ResourceId: "myAwesomeFunction"}}, Timeout: 300, InstallCommand: "npm install", BuildCommand: "npm run build", StartCommand: "node custom-server.mjs", OutputDirectory: "build", InstallationId: "6m40at4ejk5h2u9s1hboo", ProviderRepositoryId: "appwrite", ProviderBranch: "main", ProviderRootDirectory: "sites/helloWorld", ProviderSilentMode: true, ProviderBranches: []string{"test"}, ProviderPaths: []string{"test"}, BuildSpecification: "s-1vcpu-512mb", RuntimeSpecification: "s-1vcpu-512mb", BuildRuntime: "node-22", Adapter: "static", FallbackFile: "index.html"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
