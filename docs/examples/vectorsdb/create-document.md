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
		appwrite.WithSession(""),
	)

	service := vectorsdb.New(client)

	response, err := service.CreateDocument(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<DOCUMENT_ID>",
		map[string]interface{}{"embeddings": []interface{}{0.12, -0.55, 0.88, 1.02}, "metadata": map[string]interface{}{"key": "value"}},
		service.WithCreateDocumentPermissions([]string{"read(\"any\")"}),
		service.WithCreateDocumentTransactionId("<TRANSACTION_ID>"),
	)
	fmt.Println(response, err)
}
```
