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

	response, err := service.GetDeploymentDownload(
		"<SITE_ID>",
		"<DEPLOYMENT_ID>",
		service.WithGetDeploymentDownloadType("source"),
		service.WithGetDeploymentDownloadToken("<TOKEN>"),
	)
	fmt.Println(response, err)
}
```
