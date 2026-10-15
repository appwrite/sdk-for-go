package models

import (
	"encoding/json"
	"errors"
)

// Function Model
type Function struct {
	// Function ID.
	Id string `json:"$id"`
	// Function creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Function update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Execution permissions.
	Execute []string `json:"execute"`
	// Function name.
	Name string `json:"name"`
	// Function enabled.
	Enabled bool `json:"enabled"`
	// Is the function deployed with the latest configuration? This is set to
	// false if you've changed an environment variables, entrypoint, commands, or
	// other settings that needs redeploy to be applied. When the value is false,
	// redeploy the function to update it with the latest configuration.
	Live bool `json:"live"`
	// When disabled, executions will exclude logs and errors, and will be
	// slightly faster.
	Logging bool `json:"logging"`
	// Function execution and build runtime.
	Runtime string `json:"runtime"`
	// How many days to keep the non-active deployments before they will be
	// automatically deleted.
	DeploymentRetention int `json:"deploymentRetention"`
	// Function's active deployment ID.
	DeploymentId string `json:"deploymentId"`
	// Active deployment creation date in ISO 8601 format.
	DeploymentCreatedAt string `json:"deploymentCreatedAt"`
	// Function's latest deployment ID.
	LatestDeploymentId string `json:"latestDeploymentId"`
	// Latest deployment creation date in ISO 8601 format.
	LatestDeploymentCreatedAt string `json:"latestDeploymentCreatedAt"`
	// Status of latest deployment. Possible values are "waiting", "processing",
	// "building", "ready", and "failed".
	LatestDeploymentStatus string `json:"latestDeploymentStatus"`
	// Allowed permission scopes.
	Scopes []string `json:"scopes"`
	// Function variables.
	Vars []Variable `json:"vars"`
	// Function trigger events.
	Events []string `json:"events"`
	// Function execution schedule in CRON format.
	Schedule string `json:"schedule"`
	// Minutes between scheduled executions. 0 when the function has no interval.
	Interval *int `json:"interval"`
	// Function execution timeout in seconds.
	Timeout int `json:"timeout"`
	// The entrypoint file used to execute the deployment.
	Entrypoint string `json:"entrypoint"`
	// The build command used to build the deployment.
	Commands string `json:"commands"`
	// Version of Open Runtimes used for the function.
	Version string `json:"version"`
	// Function VCS (Version Control System) installation id.
	InstallationId string `json:"installationId"`
	// VCS (Version Control System) Repository ID
	ProviderRepositoryId string `json:"providerRepositoryId"`
	// VCS (Version Control System) branch name
	ProviderBranch string `json:"providerBranch"`
	// Path to function in VCS (Version Control System) repository
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
	// Machine specification for executions.
	RuntimeSpecification string `json:"runtimeSpecification"`
	// Whether inbound email (SMTP) may trigger this function. Opt-in; default is
	// false. When disabled, recipients for this function are rejected.
	EmailEnabled *bool `json:"emailEnabled"`

	// Used by Decode() method
	data []byte
}

func (model Function) New(data []byte) *Function {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Function) UnmarshalJSON(data []byte) error {
	type alias Function
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                        json.RawMessage `json:"$id"`
		CreatedAt                 json.RawMessage `json:"$createdAt"`
		UpdatedAt                 json.RawMessage `json:"$updatedAt"`
		Execute                   json.RawMessage `json:"execute"`
		Name                      json.RawMessage `json:"name"`
		Enabled                   json.RawMessage `json:"enabled"`
		Live                      json.RawMessage `json:"live"`
		Logging                   json.RawMessage `json:"logging"`
		Runtime                   json.RawMessage `json:"runtime"`
		DeploymentRetention       json.RawMessage `json:"deploymentRetention"`
		DeploymentId              json.RawMessage `json:"deploymentId"`
		DeploymentCreatedAt       json.RawMessage `json:"deploymentCreatedAt"`
		LatestDeploymentId        json.RawMessage `json:"latestDeploymentId"`
		LatestDeploymentCreatedAt json.RawMessage `json:"latestDeploymentCreatedAt"`
		LatestDeploymentStatus    json.RawMessage `json:"latestDeploymentStatus"`
		Scopes                    json.RawMessage `json:"scopes"`
		Events                    json.RawMessage `json:"events"`
		Schedule                  json.RawMessage `json:"schedule"`
		Timeout                   json.RawMessage `json:"timeout"`
		Entrypoint                json.RawMessage `json:"entrypoint"`
		Commands                  json.RawMessage `json:"commands"`
		Version                   json.RawMessage `json:"version"`
		InstallationId            json.RawMessage `json:"installationId"`
		ProviderRepositoryId      json.RawMessage `json:"providerRepositoryId"`
		ProviderBranch            json.RawMessage `json:"providerBranch"`
		ProviderRootDirectory     json.RawMessage `json:"providerRootDirectory"`
		ProviderSilentMode        json.RawMessage `json:"providerSilentMode"`
		ProviderBranches          json.RawMessage `json:"providerBranches"`
		ProviderPaths             json.RawMessage `json:"providerPaths"`
		BuildSpecification        json.RawMessage `json:"buildSpecification"`
		RuntimeSpecification      json.RawMessage `json:"runtimeSpecification"`
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
	if present.Execute == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Execute); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
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
	if present.Runtime == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Runtime); err != nil {
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
	if present.Events == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Events); err != nil {
			return err
		}
	}
	if present.Schedule == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Schedule); err != nil {
			return err
		}
	}
	if present.Timeout == nil {
		if err := json.Unmarshal([]byte("15"), &decoded.Timeout); err != nil {
			return err
		}
	}
	if present.Entrypoint == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Entrypoint); err != nil {
			return err
		}
	}
	if present.Commands == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Commands); err != nil {
			return err
		}
	}
	if present.Version == nil {
		if err := json.Unmarshal([]byte("\"v5\""), &decoded.Version); err != nil {
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

	*model = Function(decoded)
	return nil
}

func (model *Function) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
