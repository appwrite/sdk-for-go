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
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := tablesdb.New(client)

	response, err := service.UpdateRows(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		tablesdb.WithUpdateRowsData(map[string]interface{}{
        "username": "walter.obrien",
        "email": "walter.obrien@example.com",
        "fullName": "Walter O'Brien",
        "age": 33,
        "isAdmin": false
    }),
		tablesdb.WithUpdateRowsQueries([]string{}),
		tablesdb.WithUpdateRowsTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
