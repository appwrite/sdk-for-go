```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v3/client"
    "github.com/appwrite/sdk-for-go/v3/avatars"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithSession("")
)

service := avatars.New(client)

response, error := service.GetFlag(
    "af",
    avatars.WithGetFlagWidth(0),
    avatars.WithGetFlagHeight(0),
    avatars.WithGetFlagQuality(-1),
)
```
