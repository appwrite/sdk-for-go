```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/mongo"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := mongo.New(client)

response, error := service.ListOperations(
    "<DATABASE_ID>",
    mongo.WithListOperationsStatus("running"),
    mongo.WithListOperationsLimit(1),
    mongo.WithListOperationsOffset(0),
)
```
