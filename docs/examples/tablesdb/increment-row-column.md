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

	response, err := service.IncrementRowColumn(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<ROW_ID>",
		"<COLUMN>",
		service.WithIncrementRowColumnValue(1),
		service.WithIncrementRowColumnMax(100),
		service.WithIncrementRowColumnTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
