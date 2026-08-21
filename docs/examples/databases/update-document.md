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
		appwrite.WithSession(""),
	)

	service := databases.New(client)

	response, err := service.UpdateDocument(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<DOCUMENT_ID>",
		service.WithUpdateDocumentData(map[string]interface{}{"username": "walter.obrien", "email": "walter.obrien@example.com", "fullName": "Walter O'Brien", "age": 33, "isAdmin": false}),
		service.WithUpdateDocumentPermissions([]string{"read(\"any\")"}),
		service.WithUpdateDocumentTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
