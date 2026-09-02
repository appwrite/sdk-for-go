```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/databases"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := databases.New(client)

	response, err := service.UpdateDocuments(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		service.WithUpdateDocumentsData(map[string]interface{}{"username": "walter.obrien", "email": "walter.obrien@example.com", "fullName": "Walter O'Brien", "age": 33, "isAdmin": false}),
		service.WithUpdateDocumentsQueries([]string{"example"}),
		service.WithUpdateDocumentsTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
