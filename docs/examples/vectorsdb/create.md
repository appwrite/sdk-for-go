```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/vectorsdb"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := vectorsdb.New(client)

	response, err := service.Create(
		"<DATABASE_ID>",
		"<NAME>",
		service.WithCreateEnabled(false),
		service.WithCreateSpecification("serverless"),
		service.WithCreateReplicas(0),
		service.WithCreateSyncMode("async"),
	)
	fmt.Println(response, err)
}
```
