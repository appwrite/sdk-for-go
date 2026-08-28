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

	response, err := service.DecrementDocumentAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<DOCUMENT_ID>",
		"<ATTRIBUTE>",
		service.WithDecrementDocumentAttributeValue(1),
		service.WithDecrementDocumentAttributeMin(0),
		service.WithDecrementDocumentAttributeTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
