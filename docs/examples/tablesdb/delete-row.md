```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/tablesdb"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := tablesdb.New(client)

	response, err := service.DeleteRow(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<ROW_ID>",
		tablesdb.WithDeleteRowTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
