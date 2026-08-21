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

	response, err := service.CreateDocument(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<DOCUMENT_ID>",
		map[string]interface{}{
        "embeddings": [
            0.12,
            -0.55,
            0.88,
            1.02
        ],
        "metadata": {
            "key": "value"
        }
    },
		vectorsdb.WithCreateDocumentPermissions([]string{"read(\"any\")"}),
	)
	fmt.Println(response, err)
}
```
