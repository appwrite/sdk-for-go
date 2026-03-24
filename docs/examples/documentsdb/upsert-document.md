```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/documentsdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithSession("")
)

service := documentsdb.New(client)

response, error := service.UpsertDocument(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<DOCUMENT_ID>",
    documentsdb.WithUpsertDocumentData(map[string]interface{}{}),
    documentsdb.WithUpsertDocumentPermissions(interface{}{"read("any")"}),
    documentsdb.WithUpsertDocumentTransactionId("<TRANSACTION_ID>"),
)
```
