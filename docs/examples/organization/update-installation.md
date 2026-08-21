```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/organization"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := organization.New(client)

	response, err := service.UpdateInstallation(
		"<INSTALLATION_ID>",
		service.WithUpdateInstallationAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
	)
	fmt.Println(response, err)
}
```
