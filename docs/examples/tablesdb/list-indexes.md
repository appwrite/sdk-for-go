```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/tablesdb"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := tablesdb.New(client)

	response, err := service.ListIndexes(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		service.WithListIndexesQueries([]string{"example"}),
		service.WithListIndexesTotal(false),
	)
	fmt.Println(response, err)
}
```
