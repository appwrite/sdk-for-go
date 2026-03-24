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

response, error := service.CreateDocuments(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    []interface{}{},
)
```
