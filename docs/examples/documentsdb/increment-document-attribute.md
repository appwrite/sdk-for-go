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

	response, err := service.IncrementDocumentAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<DOCUMENT_ID>",
		"<ATTRIBUTE>",
		service.WithIncrementDocumentAttributeValue(1),
		service.WithIncrementDocumentAttributeMax(100),
		service.WithIncrementDocumentAttributeTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
