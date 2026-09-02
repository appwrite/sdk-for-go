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
		service.WithUpdateEvents([]string{"example"}),
		service.WithUpdateSchedule("0 0 * * *"),
		service.WithUpdateTimeout(1),
		service.WithUpdateEnabled(false),
		service.WithUpdateLogging(false),
		service.WithUpdateEntrypoint("<ENTRYPOINT>"),
		service.WithUpdateCommands("<COMMANDS>"),
		service.WithUpdateScopes([]string{"example"}),
		service.WithUpdateInstallationId("<INSTALLATION_ID>"),
		service.WithUpdateProviderRepositoryId("<PROVIDER_REPOSITORY_ID>"),
		service.WithUpdateProviderBranch("<PROVIDER_BRANCH>"),
		service.WithUpdateProviderSilentMode(false),
		service.WithUpdateProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
		service.WithUpdateProviderBranches([]string{"example"}),
		service.WithUpdateProviderPaths([]string{"example"}),
		service.WithUpdateBuildSpecification("s-1vcpu-512mb"),
		service.WithUpdateRuntimeSpecification("s-1vcpu-512mb"),
		service.WithUpdateDeploymentRetention(0),
	)
	fmt.Println(response, err)
}
```
