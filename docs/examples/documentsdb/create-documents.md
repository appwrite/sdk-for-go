```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/documentsdb"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := documentsdb.New(client)

	response, err := service.CreateDocuments(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		[]interface{}{map[string]interface{}{"$id": "example1", "username": "walter.obrien", "email": "walter.obrien@example.com", "fullName": "Walter O'Brien", "age": 30, "isAdmin": false}},
		service.WithCreateDocumentsTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
