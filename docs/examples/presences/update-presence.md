```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v4/client"
    "github.com/appwrite/sdk-for-go/v4/presences"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := presences.New(client)

response, error := service.UpdatePresence(
    "<PRESENCE_ID>",
    "<USER_ID>",
    presences.WithUpdatePresenceStatus("<STATUS>"),
    presences.WithUpdatePresenceExpiresAt("2020-10-15T06:38:00.000+00:00"),
    presences.WithUpdatePresenceMetadata(map[string]interface{}{}),
    presences.WithUpdatePresencePermissions(interface{}{"read("any")"}),
    presences.WithUpdatePresencePurge(false),
)
```
