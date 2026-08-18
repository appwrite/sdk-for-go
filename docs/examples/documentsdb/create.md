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

response, error := service.Create(
    "<DATABASE_ID>",
    "<NAME>",
    documentsdb.WithCreateEnabled(false),
    documentsdb.WithCreateSpecification("serverless"),
    documentsdb.WithCreateReplicas(0),
    documentsdb.WithCreateSyncMode("async"),
)
```
