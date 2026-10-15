package models

import (
	"encoding/json"
	"errors"
)

// Deployment Model
type Deployment struct {
	// Deployment ID.
	Id string `json:"$id"`
	// Deployment creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Deployment update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Type of deployment.
	Type string `json:"type"`
	// Resource ID.
	ResourceId string `json:"resourceId"`
	// Resource type.
	ResourceType string `json:"resourceType"`
	// The entrypoint file to use to execute the deployment code.
	Entrypoint string `json:"entrypoint"`
	// The code size in bytes.
	SourceSize int `json:"sourceSize"`
	// The build output size in bytes.
	BuildSize int `json:"buildSize"`
	// The total size in bytes (source and build output).
	TotalSize int `json:"totalSize"`
	// The current build ID.
	BuildId string `json:"buildId"`
	// Whether the deployment should be automatically activated.
	Activate bool `json:"activate"`
	// Screenshot with light theme preference file ID.
	ScreenshotLight string `json:"screenshotLight"`
	// Screenshot with dark theme preference file ID.
	ScreenshotDark string `json:"screenshotDark"`
	// The deployment status. Possible values are "waiting", "processing",
	// "building", "ready", "canceled" and "failed".
	Status string `json:"status"`
	// The build logs.
	BuildLogs string `json:"buildLogs"`
	// The current build time in seconds.
	BuildDuration int `json:"buildDuration"`
	// The name of the vcs provider repository
	ProviderRepositoryName string `json:"providerRepositoryName"`
	// The name of the vcs provider repository owner
	ProviderRepositoryOwner string `json:"providerRepositoryOwner"`
	// The url of the vcs provider repository
	ProviderRepositoryUrl string `json:"providerRepositoryUrl"`
	// The commit hash of the vcs commit
	ProviderCommitHash string `json:"providerCommitHash"`
	// The url of vcs commit author
	ProviderCommitAuthorUrl string `json:"providerCommitAuthorUrl"`
	// The name of vcs commit author
	ProviderCommitAuthor string `json:"providerCommitAuthor"`
	// The commit message
	ProviderCommitMessage string `json:"providerCommitMessage"`
	// The url of the vcs commit
	ProviderCommitUrl string `json:"providerCommitUrl"`
	// The branch of the vcs repository
	ProviderBranch string `json:"providerBranch"`
	// The branch of the vcs repository
	ProviderBranchUrl string `json:"providerBranchUrl"`

	// Used by Decode() method
	data []byte
}

func (model Deployment) New(data []byte) *Deployment {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Deployment) UnmarshalJSON(data []byte) error {
	type alias Deployment
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                      json.RawMessage `json:"$id"`
		CreatedAt               json.RawMessage `json:"$createdAt"`
		UpdatedAt               json.RawMessage `json:"$updatedAt"`
		Type                    json.RawMessage `json:"type"`
		ResourceId              json.RawMessage `json:"resourceId"`
		ResourceType            json.RawMessage `json:"resourceType"`
		Entrypoint              json.RawMessage `json:"entrypoint"`
		SourceSize              json.RawMessage `json:"sourceSize"`
		BuildSize               json.RawMessage `json:"buildSize"`
		TotalSize               json.RawMessage `json:"totalSize"`
		BuildId                 json.RawMessage `json:"buildId"`
		Activate                json.RawMessage `json:"activate"`
		ScreenshotLight         json.RawMessage `json:"screenshotLight"`
		ScreenshotDark          json.RawMessage `json:"screenshotDark"`
		BuildLogs               json.RawMessage `json:"buildLogs"`
		BuildDuration           json.RawMessage `json:"buildDuration"`
		ProviderRepositoryName  json.RawMessage `json:"providerRepositoryName"`
		ProviderRepositoryOwner json.RawMessage `json:"providerRepositoryOwner"`
		ProviderRepositoryUrl   json.RawMessage `json:"providerRepositoryUrl"`
		ProviderCommitHash      json.RawMessage `json:"providerCommitHash"`
		ProviderCommitAuthorUrl json.RawMessage `json:"providerCommitAuthorUrl"`
		ProviderCommitAuthor    json.RawMessage `json:"providerCommitAuthor"`
		ProviderCommitMessage   json.RawMessage `json:"providerCommitMessage"`
		ProviderCommitUrl       json.RawMessage `json:"providerCommitUrl"`
		ProviderBranch          json.RawMessage `json:"providerBranch"`
		ProviderBranchUrl       json.RawMessage `json:"providerBranchUrl"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.ResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceId); err != nil {
			return err
		}
	}
	if present.ResourceType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceType); err != nil {
			return err
		}
	}
	if present.Entrypoint == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Entrypoint); err != nil {
			return err
		}
	}
	if present.SourceSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.SourceSize); err != nil {
			return err
		}
	}
	if present.BuildSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.BuildSize); err != nil {
			return err
		}
	}
	if present.TotalSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.TotalSize); err != nil {
			return err
		}
	}
	if present.BuildId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BuildId); err != nil {
			return err
		}
	}
	if present.Activate == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Activate); err != nil {
			return err
		}
	}
	if present.ScreenshotLight == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ScreenshotLight); err != nil {
			return err
		}
	}
	if present.ScreenshotDark == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ScreenshotDark); err != nil {
			return err
		}
	}
	if present.BuildLogs == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BuildLogs); err != nil {
			return err
		}
	}
	if present.BuildDuration == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.BuildDuration); err != nil {
			return err
		}
	}
	if present.ProviderRepositoryName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderRepositoryName); err != nil {
			return err
		}
	}
	if present.ProviderRepositoryOwner == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderRepositoryOwner); err != nil {
			return err
		}
	}
	if present.ProviderRepositoryUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderRepositoryUrl); err != nil {
			return err
		}
	}
	if present.ProviderCommitHash == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderCommitHash); err != nil {
			return err
		}
	}
	if present.ProviderCommitAuthorUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderCommitAuthorUrl); err != nil {
			return err
		}
	}
	if present.ProviderCommitAuthor == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderCommitAuthor); err != nil {
			return err
		}
	}
	if present.ProviderCommitMessage == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderCommitMessage); err != nil {
			return err
		}
	}
	if present.ProviderCommitUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderCommitUrl); err != nil {
			return err
		}
	}
	if present.ProviderBranch == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderBranch); err != nil {
			return err
		}
	}
	if present.ProviderBranchUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderBranchUrl); err != nil {
			return err
		}
	}

	*model = Deployment(decoded)
	return nil
}

func (model *Deployment) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
