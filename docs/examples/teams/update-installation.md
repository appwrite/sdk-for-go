```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/teams"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := teams.New(client)

	response, err := service.UpdateInstallation(
		"<TEAM_ID>",
		"<INSTALLATION_ID>",
		teams.WithUpdateInstallationAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
	)
	fmt.Println(response, err)
}
```
