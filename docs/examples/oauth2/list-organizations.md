```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/oauth2"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithSession(""),
		client.WithProject("<YOUR_PROJECT_ID>"),
	)

	service := oauth2.New(client)

	response, err := service.ListOrganizations(
		oauth2.WithListOrganizationsLimit(1),
		oauth2.WithListOrganizationsOffset(0),
		oauth2.WithListOrganizationsSearch("<SEARCH>"),
	)
	fmt.Println(response, err)
}
```
