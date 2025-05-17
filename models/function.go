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
    // Machine specification for builds and executions.
    Specification string `json:"specification"`

    // Used by Decode() method
    data []byte
}

func (model Function) New(data []byte) *Function {
    model.data = data
    return &model
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