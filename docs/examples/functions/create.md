```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v6/client"
    "github.com/appwrite/sdk-for-go/v6/functions"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := functions.New(client)

response, error := service.Create(
    "<FUNCTION_ID>",
    "<NAME>",
    "node-14.5",
    functions.WithCreateExecute([]string{"any"}),
    functions.WithCreateEvents([]string{}),
    functions.WithCreateSchedule(""),
    functions.WithCreateTimeout(1),
    functions.WithCreateEnabled(false),
    functions.WithCreateLogging(false),
    functions.WithCreateEntrypoint("<ENTRYPOINT>"),
    functions.WithCreateCommands("<COMMANDS>"),
    functions.WithCreateScopes([]string{}),
    functions.WithCreateInstallationId("<INSTALLATION_ID>"),
    functions.WithCreateProviderRepositoryId("<PROVIDER_REPOSITORY_ID>"),
    functions.WithCreateProviderBranch("<PROVIDER_BRANCH>"),
    functions.WithCreateProviderSilentMode(false),
    functions.WithCreateProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
    functions.WithCreateProviderBranches([]string{}),
    functions.WithCreateProviderPaths([]string{}),
    functions.WithCreateBuildSpecification(""),
    functions.WithCreateRuntimeSpecification(""),
    functions.WithCreateDeploymentRetention(0),
)
```
