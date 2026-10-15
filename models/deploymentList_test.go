package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDeploymentListModel(t *testing.T) {
	model := DeploymentList{Total: 5, Deployments: []Deployment{Deployment{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Type: "vcs", ResourceId: "5e5ea6g16897e", ResourceType: "functions", Entrypoint: "index.js", SourceSize: 128, BuildSize: 128, TotalSize: 128, BuildId: "5e5ea5c16897e", Activate: true, ScreenshotLight: "5e5ea5c16897e", ScreenshotDark: "5e5ea5c16897e", Status: "ready", BuildLogs: "Compiling source files...", BuildDuration: 128, ProviderRepositoryName: "database", ProviderRepositoryOwner: "utopia", ProviderRepositoryUrl: "https://github.com/vermakhushboo/g4-node-function", ProviderCommitHash: "7c3f25d", ProviderCommitAuthorUrl: "https://github.com/vermakhushboo", ProviderCommitAuthor: "Khushboo Verma", ProviderCommitMessage: "Update index.js", ProviderCommitUrl: "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb", ProviderBranch: "0.7.x", ProviderBranchUrl: "https://github.com/vermakhushboo/appwrite/tree/0.7.x"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DeploymentList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestDeploymentListModelDefaults(t *testing.T) {
	var fresh DeploymentList
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

	reused := DeploymentList{Total: 5, Deployments: []Deployment{Deployment{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Type: "vcs", ResourceId: "5e5ea6g16897e", ResourceType: "functions", Entrypoint: "index.js", SourceSize: 128, BuildSize: 128, TotalSize: 128, BuildId: "5e5ea5c16897e", Activate: true, ScreenshotLight: "5e5ea5c16897e", ScreenshotDark: "5e5ea5c16897e", Status: "ready", BuildLogs: "Compiling source files...", BuildDuration: 128, ProviderRepositoryName: "database", ProviderRepositoryOwner: "utopia", ProviderRepositoryUrl: "https://github.com/vermakhushboo/g4-node-function", ProviderCommitHash: "7c3f25d", ProviderCommitAuthorUrl: "https://github.com/vermakhushboo", ProviderCommitAuthor: "Khushboo Verma", ProviderCommitMessage: "Update index.js", ProviderCommitUrl: "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb", ProviderBranch: "0.7.x", ProviderBranchUrl: "https://github.com/vermakhushboo/appwrite/tree/0.7.x"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
