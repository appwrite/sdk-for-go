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
		appwrite.WithSession(""),
	)

	service := tablesdb.New(client)

	response, err := service.ListRows(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		service.WithListRowsQueries([]string{"example"}),
		service.WithListRowsTransactionId("<TRANSACTION_ID>"),
		service.WithListRowsTotal(false),
		service.WithListRowsTtl(0),
	)
	fmt.Println(response, err)
}
```
