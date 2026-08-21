```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/backups"
	"github.com/appwrite/sdk-for-go/v7/client"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := backups.New(client)

	response, err := service.CreateRestoration(
		"<ARCHIVE_ID>",
		[]string{},
		backups.WithCreateRestorationNewResourceId("<NEW_RESOURCE_ID>"),
		backups.WithCreateRestorationNewResourceName("<NEW_RESOURCE_NAME>"),
	)
	fmt.Println(response, err)
}
```
