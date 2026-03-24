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
    client.WithKey("<YOUR_API_KEY>")
)

service := documentsdb.New(client)

response, error := service.DeleteDocuments(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    documentsdb.WithDeleteDocumentsQueries([]interface{}{}),
    documentsdb.WithDeleteDocumentsTransactionId("<TRANSACTION_ID>"),
)
```
