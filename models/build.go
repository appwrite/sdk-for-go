package models


// Build Model
type Build struct {
    // Build ID.
    Id string `json:"$id"`
    // The deployment that created this build.
    DeploymentId string `json:"deploymentId"`
    // The build status. There are a few different types and each one means
    // something different. \nFailed - The deployment build has failed. More
    // details can usually be found in buildStderr\nReady - The deployment build
    // was successful and the deployment is ready to be deployed\nProcessing - The
    // deployment is currently waiting to have a build triggered\nBuilding - The
    // deployment is currently being built
    Status string `json:"status"`
    // The stdout of the build.
    Stdout string `json:"stdout"`
    // The stderr of the build.
    Stderr string `json:"stderr"`
    // The deployment creation date in ISO 8601 format.
    StartTime string `json:"startTime"`
    // The time the build was finished in ISO 8601 format.
    EndTime string `json:"endTime"`
    // The build duration in seconds.
    Duration int `json:"duration"`
    // The code size in bytes.
    Size int `json:"size"`

}
