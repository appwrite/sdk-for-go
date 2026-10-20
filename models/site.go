package models

import (
	"encoding/json"
	"errors"
)

// Site Model
type Site struct {
	// Site ID.
	Id string `json:"$id"`
	// Site creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Site update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Site name.
	Name string `json:"name"`
	// Optional description for the site.
	Description string `json:"description"`
	// Site enabled.
	Enabled bool `json:"enabled"`
	// Is the site deployed with the latest configuration? This is set to false if
	// you've changed an environment variables, entrypoint, commands, or other
	// settings that needs redeploy to be applied. When the value is false,
	// redeploy the site to update it with the latest configuration.
	Live bool `json:"live"`
	// When disabled, request logs will exclude logs and errors, and site
	// responses will be slightly faster.
	Logging bool `json:"logging"`
	// Site framework.
	Framework string `json:"framework"`
	// How many days to keep the non-active deployments before they will be
	// automatically deleted.
	DeploymentRetention int `json:"deploymentRetention"`
	// Site's active deployment ID.
	DeploymentId string `json:"deploymentId"`
	// Active deployment creation date in ISO 8601 format.
	DeploymentCreatedAt string `json:"deploymentCreatedAt"`
	// Screenshot of active deployment with light theme preference file ID.
	DeploymentScreenshotLight string `json:"deploymentScreenshotLight"`
	// Screenshot of active deployment with dark theme preference file ID.
	DeploymentScreenshotDark string `json:"deploymentScreenshotDark"`
	// Site's latest deployment ID.
	LatestDeploymentId string `json:"latestDeploymentId"`
	// Latest deployment creation date in ISO 8601 format.
	LatestDeploymentCreatedAt string `json:"latestDeploymentCreatedAt"`
	// Status of latest deployment. Possible values are "waiting", "processing",
	// "building", "ready", and "failed".
	LatestDeploymentStatus string `json:"latestDeploymentStatus"`
	// Allowed permission scopes.
	Scopes []string `json:"scopes"`
	// Site variables.
	Vars []Variable `json:"vars"`
	// Site request timeout in seconds.
	Timeout int `json:"timeout"`
	// The install command used to install the site dependencies.
	InstallCommand string `json:"installCommand"`
	// The build command used to build the site.
	BuildCommand string `json:"buildCommand"`
	// Custom command to use when starting site runtime.
	StartCommand string `json:"startCommand"`
	// The directory where the site build output is located.
	OutputDirectory string `json:"outputDirectory"`
	// Site VCS (Version Control System) installation id.
	InstallationId string `json:"installationId"`
	// VCS (Version Control System) Repository ID
	ProviderRepositoryId string `json:"providerRepositoryId"`
	// VCS (Version Control System) branch name
	ProviderBranch string `json:"providerBranch"`
	// Path to site in VCS (Version Control System) repository
	ProviderRootDirectory string `json:"providerRootDirectory"`
	// Is VCS (Version Control System) connection is in silent mode? When in
	// silence mode, no comments will be posted on the repository pull or merge
	// requests
	ProviderSilentMode bool `json:"providerSilentMode"`
	// List of branch name patterns that trigger automatic deployments. Supports
	// glob wildcards. Empty list deploys on all branches.
	ProviderBranches []string `json:"providerBranches"`
	// List of file path patterns that trigger automatic deployments. Supports
	// glob wildcards. Empty list deploys on all file changes.
	ProviderPaths []string `json:"providerPaths"`
	// Machine specification for deployment builds.
	BuildSpecification string `json:"buildSpecification"`
	// Machine specification for SSR executions.
	RuntimeSpecification string `json:"runtimeSpecification"`
	// Site build runtime.
	BuildRuntime string `json:"buildRuntime"`
	// Site framework adapter.
	Adapter string `json:"adapter"`
	// Name of fallback file to use instead of 404 page. If null, Appwrite 404
	// page will be displayed.
	FallbackFile string `json:"fallbackFile"`

	// Used by Decode() method
	data []byte
}

func (model Site) New(data []byte) *Site {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Site) UnmarshalJSON(data []byte) error {
	type alias Site
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                        json.RawMessage `json:"$id"`
		CreatedAt                 json.RawMessage `json:"$createdAt"`
		UpdatedAt                 json.RawMessage `json:"$updatedAt"`
		Name                      json.RawMessage `json:"name"`
		Description               json.RawMessage `json:"description"`
		Enabled                   json.RawMessage `json:"enabled"`
		Live                      json.RawMessage `json:"live"`
		Logging                   json.RawMessage `json:"logging"`
		Framework                 json.RawMessage `json:"framework"`
		DeploymentRetention       json.RawMessage `json:"deploymentRetention"`
		DeploymentId              json.RawMessage `json:"deploymentId"`
		DeploymentCreatedAt       json.RawMessage `json:"deploymentCreatedAt"`
		DeploymentScreenshotLight json.RawMessage `json:"deploymentScreenshotLight"`
		DeploymentScreenshotDark  json.RawMessage `json:"deploymentScreenshotDark"`
		LatestDeploymentId        json.RawMessage `json:"latestDeploymentId"`
		LatestDeploymentCreatedAt json.RawMessage `json:"latestDeploymentCreatedAt"`
		LatestDeploymentStatus    json.RawMessage `json:"latestDeploymentStatus"`
		Scopes                    json.RawMessage `json:"scopes"`
		Timeout                   json.RawMessage `json:"timeout"`
		InstallCommand            json.RawMessage `json:"installCommand"`
		BuildCommand              json.RawMessage `json:"buildCommand"`
		StartCommand              json.RawMessage `json:"startCommand"`
		OutputDirectory           json.RawMessage `json:"outputDirectory"`
		InstallationId            json.RawMessage `json:"installationId"`
		ProviderRepositoryId      json.RawMessage `json:"providerRepositoryId"`
		ProviderBranch            json.RawMessage `json:"providerBranch"`
		ProviderRootDirectory     json.RawMessage `json:"providerRootDirectory"`
		ProviderSilentMode        json.RawMessage `json:"providerSilentMode"`
		ProviderBranches          json.RawMessage `json:"providerBranches"`
		ProviderPaths             json.RawMessage `json:"providerPaths"`
		BuildSpecification        json.RawMessage `json:"buildSpecification"`
		RuntimeSpecification      json.RawMessage `json:"runtimeSpecification"`
		BuildRuntime              json.RawMessage `json:"buildRuntime"`
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
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Live == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Live); err != nil {
			return err
		}
	}
	if present.Logging == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Logging); err != nil {
			return err
		}
	}
	if present.Framework == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Framework); err != nil {
			return err
		}
	}
	if present.DeploymentRetention == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.DeploymentRetention); err != nil {
			return err
		}
	}
	if present.DeploymentId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeploymentId); err != nil {
			return err
		}
	}
	if present.DeploymentCreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeploymentCreatedAt); err != nil {
			return err
		}
	}
	if present.DeploymentScreenshotLight == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeploymentScreenshotLight); err != nil {
			return err
		}
	}
	if present.DeploymentScreenshotDark == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeploymentScreenshotDark); err != nil {
			return err
		}
	}
	if present.LatestDeploymentId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.LatestDeploymentId); err != nil {
			return err
		}
	}
	if present.LatestDeploymentCreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.LatestDeploymentCreatedAt); err != nil {
			return err
		}
	}
	if present.LatestDeploymentStatus == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.LatestDeploymentStatus); err != nil {
			return err
		}
	}
	if present.Scopes == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Scopes); err != nil {
			return err
		}
	}
	if present.Timeout == nil {
		if err := json.Unmarshal([]byte("15"), &decoded.Timeout); err != nil {
			return err
		}
	}
	if present.InstallCommand == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.InstallCommand); err != nil {
			return err
		}
	}
	if present.BuildCommand == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BuildCommand); err != nil {
			return err
		}
	}
	if present.StartCommand == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.StartCommand); err != nil {
			return err
		}
	}
	if present.OutputDirectory == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.OutputDirectory); err != nil {
			return err
		}
	}
	if present.InstallationId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.InstallationId); err != nil {
			return err
		}
	}
	if present.ProviderRepositoryId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderRepositoryId); err != nil {
			return err
		}
	}
	if present.ProviderBranch == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderBranch); err != nil {
			return err
		}
	}
	if present.ProviderRootDirectory == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderRootDirectory); err != nil {
			return err
		}
	}
	if present.ProviderSilentMode == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.ProviderSilentMode); err != nil {
			return err
		}
	}
	if present.ProviderBranches == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.ProviderBranches); err != nil {
			return err
		}
	}
	if present.ProviderPaths == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.ProviderPaths); err != nil {
			return err
		}
	}
	if present.BuildSpecification == nil {
		if err := json.Unmarshal([]byte("\"s-1vcpu-512mb\""), &decoded.BuildSpecification); err != nil {
			return err
		}
	}
	if present.RuntimeSpecification == nil {
		if err := json.Unmarshal([]byte("\"s-1vcpu-512mb\""), &decoded.RuntimeSpecification); err != nil {
			return err
		}
	}
	if present.BuildRuntime == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BuildRuntime); err != nil {
			return err
		}
	}

	*model = Site(decoded)
	return nil
}

func (model *Site) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
