```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v4/client"
    "github.com/appwrite/sdk-for-go/v4/databases"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := databases.New(client)

response, error := service.UpdateBigIntAttribute(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    "",
    false,
    0,
    databases.WithUpdateBigIntAttributeMin(0),
    databases.WithUpdateBigIntAttributeMax(0),
    databases.WithUpdateBigIntAttributeNewKey(""),
)
```
