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

	response, err := service.UpdateRows(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		service.WithUpdateRowsData(map[string]interface{}{"username": "walter.obrien", "email": "walter.obrien@example.com", "fullName": "Walter O'Brien", "age": 33, "isAdmin": false}),
		service.WithUpdateRowsQueries([]string{"example"}),
		service.WithUpdateRowsTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
