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

	response, err := service.Create(
		"<FUNCTION_ID>",
		"<NAME>",
		"node-14.5",
		service.WithCreateExecute([]string{"any"}),
		service.WithCreateEvents([]string{"example"}),
		service.WithCreateSchedule("0 0 * * *"),
		service.WithCreateTimeout(1),
		service.WithCreateEnabled(false),
		service.WithCreateLogging(false),
		service.WithCreateEntrypoint("<ENTRYPOINT>"),
		service.WithCreateCommands("<COMMANDS>"),
		service.WithCreateScopes([]string{"example"}),
		service.WithCreateInstallationId("<INSTALLATION_ID>"),
		service.WithCreateProviderRepositoryId("<PROVIDER_REPOSITORY_ID>"),
		service.WithCreateProviderBranch("<PROVIDER_BRANCH>"),
		service.WithCreateProviderSilentMode(false),
		service.WithCreateProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
		service.WithCreateProviderBranches([]string{"example"}),
		service.WithCreateProviderPaths([]string{"example"}),
		service.WithCreateBuildSpecification("s-1vcpu-512mb"),
		service.WithCreateRuntimeSpecification("s-1vcpu-512mb"),
		service.WithCreateDeploymentRetention(0),
	)
	fmt.Println(response, err)
}
```
