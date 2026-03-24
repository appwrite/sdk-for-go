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

response, error := service.CreateCollection(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "<NAME>",
    documentsdb.WithCreateCollectionPermissions(interface{}{"read("any")"}),
    documentsdb.WithCreateCollectionDocumentSecurity(false),
    documentsdb.WithCreateCollectionEnabled(false),
    documentsdb.WithCreateCollectionAttributes([]interface{}{}),
    documentsdb.WithCreateCollectionIndexes([]interface{}{}),
)
```
