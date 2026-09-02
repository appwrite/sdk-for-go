```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/sites"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := sites.New(client)

	response, err := service.Update(
		"<SITE_ID>",
		"<NAME>",
		"analog",
		service.WithUpdateEnabled(false),
		service.WithUpdateLogging(false),
		service.WithUpdateTimeout(1),
		service.WithUpdateInstallCommand("<INSTALL_COMMAND>"),
		service.WithUpdateBuildCommand("<BUILD_COMMAND>"),
		service.WithUpdateStartCommand("<START_COMMAND>"),
		service.WithUpdateOutputDirectory("<OUTPUT_DIRECTORY>"),
		service.WithUpdateBuildRuntime("node-14.5"),
		service.WithUpdateAdapter("static"),
		service.WithUpdateFallbackFile("<FALLBACK_FILE>"),
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
		service.WithUpdateScopes([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
