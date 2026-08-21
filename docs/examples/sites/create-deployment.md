```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/file"
	"github.com/appwrite/sdk-for-go/v7/sites"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := sites.New(client)

	response, err := service.CreateDeployment(
		"<SITE_ID>",
		file.NewInputFile("/path/to/file.png", "file.png"),
		service.WithCreateDeploymentInstallCommand("<INSTALL_COMMAND>"),
		service.WithCreateDeploymentBuildCommand("<BUILD_COMMAND>"),
		service.WithCreateDeploymentOutputDirectory("<OUTPUT_DIRECTORY>"),
		service.WithCreateDeploymentActivate(false),
	)
	fmt.Println(response, err)
}
```
