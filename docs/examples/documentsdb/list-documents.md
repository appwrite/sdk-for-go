```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/documentsdb"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := documentsdb.New(client)

	response, err := service.ListDocuments(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		documentsdb.WithListDocumentsQueries([]string{}),
		documentsdb.WithListDocumentsTransactionId("<TRANSACTION_ID>"),
		documentsdb.WithListDocumentsTotal(false),
		documentsdb.WithListDocumentsTtl(0),
	)
	fmt.Println(response, err)
}
```
