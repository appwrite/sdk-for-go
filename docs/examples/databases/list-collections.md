```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/databases"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := databases.New(client)

	response, err := service.ListCollections(
		"<DATABASE_ID>",
		databases.WithListCollectionsQueries([]string{}),
		databases.WithListCollectionsSearch("<SEARCH>"),
		databases.WithListCollectionsTotal(false),
	)
	fmt.Println(response, err)
}
```
