```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/documentsdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := documentsdb.New(client)

response, error := service.UpdateDocuments(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    documentsdb.WithUpdateDocumentsData(map[string]interface{}{}),
    documentsdb.WithUpdateDocumentsQueries([]string{}),
    documentsdb.WithUpdateDocumentsTransactionId("<TRANSACTION_ID>"),
)
```
