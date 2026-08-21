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

	response, err := service.UpsertRow(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<ROW_ID>",
		service.WithUpsertRowData(map[string]interface{}{"username": "walter.obrien", "email": "walter.obrien@example.com", "fullName": "Walter O'Brien", "age": 33, "isAdmin": false}),
		service.WithUpsertRowPermissions([]string{"read(\"any\")"}),
		service.WithUpsertRowTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
