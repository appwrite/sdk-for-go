```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/apps"
	"github.com/appwrite/sdk-for-go/v7/client"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := apps.New(client)

	response, err := service.ListSecrets(
		"<APP_ID>",
		apps.WithListSecretsQueries([]string{}),
		apps.WithListSecretsTotal(false),
	)
	fmt.Println(response, err)
}
```
