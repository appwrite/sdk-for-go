```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/vectorsdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithSession("")
)

service := vectorsdb.New(client)

response, error := service.GetDocument(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<DOCUMENT_ID>",
    vectorsdb.WithGetDocumentQueries([]string{}),
    vectorsdb.WithGetDocumentTransactionId("<TRANSACTION_ID>"),
)
```
