```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/vectorsdb"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := vectorsdb.New(client)

	response, err := service.UpsertDocument(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<DOCUMENT_ID>",
		vectorsdb.WithUpsertDocumentData(map[string]interface{}{}),
		vectorsdb.WithUpsertDocumentPermissions([]string{"read(\"any\")"}),
		vectorsdb.WithUpsertDocumentTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
