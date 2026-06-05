```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v5/client"
    "github.com/appwrite/sdk-for-go/v5/presences"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := presences.New(client)

response, error := service.Upsert(
    "<PRESENCE_ID>",
    "<USER_ID>",
    "<STATUS>",
    presences.WithUpsertPermissions([]string{"read("any")"}),
    presences.WithUpsertExpiresAt("2020-10-15T06:38:00.000+00:00"),
    presences.WithUpsertMetadata(map[string]interface{}{}),
)
```
