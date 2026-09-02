```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/functions"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := functions.New(client)

	response, err := service.ListDeployments(
		"<FUNCTION_ID>",
		service.WithListDeploymentsQueries([]string{"example"}),
		service.WithListDeploymentsSearch("<SEARCH>"),
		service.WithListDeploymentsTotal(false),
	)
	fmt.Println(response, err)
}
```
