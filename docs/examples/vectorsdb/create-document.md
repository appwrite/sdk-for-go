```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/vectorsdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithSession("")
)

service := vectorsdb.New(client)

response, error := service.CreateDocument(
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
    vectorsdb.WithCreateDocumentPermissions(interface{}{"read("any")"}),
)
```
