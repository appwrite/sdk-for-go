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
    client.WithKey("<YOUR_API_KEY>")
)

service := vectorsdb.New(client)

response, error := service.UpdateCollection(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<NAME>",
    vectorsdb.WithUpdateCollectionDimension(1),
    vectorsdb.WithUpdateCollectionPermissions([]string{"read("any")"}),
    vectorsdb.WithUpdateCollectionDocumentSecurity(false),
    vectorsdb.WithUpdateCollectionEnabled(false),
)
```
