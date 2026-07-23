```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v6/client"
    "github.com/appwrite/sdk-for-go/v6/oauth2"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithSession("")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := oauth2.New(client)

response, error := service.Revoke(
    "<TOKEN>",
    oauth2.WithRevokeTokenTypeHint("access_token"),
    oauth2.WithRevokeClientId("<CLIENT_ID>"),
    oauth2.WithRevokeClientSecret("<CLIENT_SECRET>"),
)
```
