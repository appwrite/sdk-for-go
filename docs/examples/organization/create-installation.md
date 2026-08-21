```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/organization"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := organization.New(client)

	response, err := service.CreateInstallation(
		"<APP_ID>",
		organization.WithCreateInstallationAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
	)
	fmt.Println(response, err)
}
```
