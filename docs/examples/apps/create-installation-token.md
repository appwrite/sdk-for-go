```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/apps"
	"github.com/appwrite/sdk-for-go/v7/appwrite"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := apps.New(client)

	response, err := service.CreateInstallationToken(
		"<APP_ID>",
		"<INSTALLATION_ID>",
	)
	fmt.Println(response, err)
}
```
