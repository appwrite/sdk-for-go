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
    // Site variables.
    Vars []Variable `json:"vars"`
    // Site request timeout in seconds.
    Timeout int `json:"timeout"`
    // The install command used to install the site dependencies.
    InstallCommand string `json:"installCommand"`
    // The build command used to build the site.
    BuildCommand string `json:"buildCommand"`
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
    // Machine specification for builds and executions.
    Specification string `json:"specification"`
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