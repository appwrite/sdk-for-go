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
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := organization.New(client)

	response, err := service.ListProjects(
		service.WithListProjectsQueries([]string{"example"}),
		service.WithListProjectsSearch("<SEARCH>"),
		service.WithListProjectsTotal(false),
	)
	fmt.Println(response, err)
}
```
