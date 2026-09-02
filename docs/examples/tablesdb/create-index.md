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

	response, err := service.CreateIndex(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<KEY>",
		"key",
		[]string{"example"},
		service.WithCreateIndexOrders([]string{"example"}),
		service.WithCreateIndexLengths([]int{0}),
	)
	fmt.Println(response, err)
}
```
