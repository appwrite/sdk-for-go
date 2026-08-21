```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/sites"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := sites.New(client)

	response, err := service.ListDeployments(
		"<SITE_ID>",
		sites.WithListDeploymentsQueries([]string{}),
		sites.WithListDeploymentsSearch("<SEARCH>"),
		sites.WithListDeploymentsTotal(false),
	)
	fmt.Println(response, err)
}
```
