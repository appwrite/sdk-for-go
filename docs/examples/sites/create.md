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

	response, err := service.Create(
		"<SITE_ID>",
		"<NAME>",
		"analog",
		"node-14.5",
		service.WithCreateEnabled(false),
		service.WithCreateLogging(false),
		service.WithCreateTimeout(1),
		service.WithCreateInstallCommand("<INSTALL_COMMAND>"),
		service.WithCreateBuildCommand("<BUILD_COMMAND>"),
		service.WithCreateStartCommand("<START_COMMAND>"),
		service.WithCreateOutputDirectory("<OUTPUT_DIRECTORY>"),
		service.WithCreateAdapter("static"),
		service.WithCreateInstallationId("<INSTALLATION_ID>"),
		service.WithCreateFallbackFile("<FALLBACK_FILE>"),
		service.WithCreateProviderRepositoryId("<PROVIDER_REPOSITORY_ID>"),
		service.WithCreateProviderBranch("<PROVIDER_BRANCH>"),
		service.WithCreateProviderSilentMode(false),
		service.WithCreateProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
		service.WithCreateProviderBranches([]string{}),
		service.WithCreateProviderPaths([]string{}),
		service.WithCreateBuildSpecification(""),
		service.WithCreateRuntimeSpecification(""),
		service.WithCreateDeploymentRetention(0),
		service.WithCreateScopes([]string{}),
	)
	fmt.Println(response, err)
}
```
