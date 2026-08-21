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

	response, err := service.UpsertDocument(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<DOCUMENT_ID>",
		service.WithUpsertDocumentData([]interface{}{}),
		service.WithUpsertDocumentPermissions([]string{"read(\"any\")"}),
		service.WithUpsertDocumentTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
