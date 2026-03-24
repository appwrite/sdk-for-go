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

response, error := service.ListDocuments(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    vectorsdb.WithListDocumentsQueries([]interface{}{}),
    vectorsdb.WithListDocumentsTransactionId("<TRANSACTION_ID>"),
    vectorsdb.WithListDocumentsTotal(false),
    vectorsdb.WithListDocumentsTtl(0),
)
```
