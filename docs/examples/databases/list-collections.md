```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v3/client"
    "github.com/appwrite/sdk-for-go/v3/databases"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := databases.New(client)

response, error := service.ListCollections(
    "<DATABASE_ID>",
    databases.WithListCollectionsQueries([]interface{}{}),
    databases.WithListCollectionsSearch("<SEARCH>"),
    databases.WithListCollectionsTotal(false),
)
```
