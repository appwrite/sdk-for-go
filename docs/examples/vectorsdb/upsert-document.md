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

response, error := service.UpsertDocument(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<DOCUMENT_ID>",
    vectorsdb.WithUpsertDocumentData(map[string]interface{}{}),
    vectorsdb.WithUpsertDocumentPermissions(interface{}{"read("any")"}),
    vectorsdb.WithUpsertDocumentTransactionId("<TRANSACTION_ID>"),
)
```
