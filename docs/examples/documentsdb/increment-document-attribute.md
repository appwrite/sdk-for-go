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

response, error := service.IncrementDocumentAttribute(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<DOCUMENT_ID>",
    "",
    documentsdb.WithIncrementDocumentAttributeValue(0),
    documentsdb.WithIncrementDocumentAttributeMax(0),
    documentsdb.WithIncrementDocumentAttributeTransactionId("<TRANSACTION_ID>"),
)
```
