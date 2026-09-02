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

	response, err := service.ListDocuments(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		service.WithListDocumentsQueries([]string{"example"}),
		service.WithListDocumentsTransactionId("<TRANSACTION_ID>"),
		service.WithListDocumentsTotal(false),
		service.WithListDocumentsTtl(0),
	)
	fmt.Println(response, err)
}
```
