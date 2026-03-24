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
    client.WithKey("<YOUR_API_KEY>")
)

service := vectorsdb.New(client)

response, error := service.UpdateDocuments(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    vectorsdb.WithUpdateDocumentsData(map[string]interface{}{}),
    vectorsdb.WithUpdateDocumentsQueries([]interface{}{}),
    vectorsdb.WithUpdateDocumentsTransactionId("<TRANSACTION_ID>"),
)
```
