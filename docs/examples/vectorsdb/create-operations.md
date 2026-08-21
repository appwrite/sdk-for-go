```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/vectorsdb"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := vectorsdb.New(client)

	response, err := service.CreateOperations(
		"<TRANSACTION_ID>",
		service.WithCreateOperationsOperations([]interface{}{map[string]interface{}{"action": "create", "databaseId": "<DATABASE_ID>", "collectionId": "<COLLECTION_ID>", "documentId": "<DOCUMENT_ID>", "data": map[string]interface{}{"name": "Walter O'Brien"}}}),
	)
	fmt.Println(response, err)
}
```
