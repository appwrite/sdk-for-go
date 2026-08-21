```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/functions"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := functions.New(client)

	response, err := service.Update(
		"<FUNCTION_ID>",
		"<NAME>",
		service.WithUpdateRuntime("node-14.5"),
		service.WithUpdateExecute([]string{"any"}),
		service.WithUpdateEvents([]string{}),
		service.WithUpdateSchedule(""),
		service.WithUpdateTimeout(1),
		service.WithUpdateEnabled(false),
		service.WithUpdateLogging(false),
		service.WithUpdateEntrypoint("<ENTRYPOINT>"),
		service.WithUpdateCommands("<COMMANDS>"),
		service.WithUpdateScopes([]string{}),
		service.WithUpdateInstallationId("<INSTALLATION_ID>"),
		service.WithUpdateProviderRepositoryId("<PROVIDER_REPOSITORY_ID>"),
		service.WithUpdateProviderBranch("<PROVIDER_BRANCH>"),
		service.WithUpdateProviderSilentMode(false),
		service.WithUpdateProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
		service.WithUpdateProviderBranches([]string{}),
		service.WithUpdateProviderPaths([]string{}),
		service.WithUpdateBuildSpecification(""),
		service.WithUpdateRuntimeSpecification(""),
		service.WithUpdateDeploymentRetention(0),
	)
	fmt.Println(response, err)
}
```
