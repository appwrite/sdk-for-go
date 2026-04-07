```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v2/client"
    "github.com/appwrite/sdk-for-go/v2/databases"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := databases.New(client)

response, error := service.UpdateCollection(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    databases.WithUpdateCollectionName("<NAME>"),
    databases.WithUpdateCollectionPermissions(interface{}{"read("any")"}),
    databases.WithUpdateCollectionDocumentSecurity(false),
    databases.WithUpdateCollectionEnabled(false),
)
```
