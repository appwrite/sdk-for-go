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

	response, err := service.CreateTemplateDeployment(
		"<SITE_ID>",
		"<REPOSITORY>",
		"<OWNER>",
		"<ROOT_DIRECTORY>",
		"branch",
		"<REFERENCE>",
		service.WithCreateTemplateDeploymentActivate(false),
	)
	fmt.Println(response, err)
}
```
