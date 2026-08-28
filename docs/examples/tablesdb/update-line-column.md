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

	response, err := service.UpdateLineColumn(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<KEY>",
		false,
		service.WithUpdateLineColumnDefault([][]interface{}{[]interface{}{1, 2}, []interface{}{3, 4}, []interface{}{5, 6}}),
		service.WithUpdateLineColumnNewKey("<NEW_KEY>"),
	)
	fmt.Println(response, err)
}
```
