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

response, error := service.CreateToken(
    "<GRANT_TYPE>",
    oauth2.WithCreateTokenCode("<CODE>"),
    oauth2.WithCreateTokenRefreshToken("<REFRESH_TOKEN>"),
    oauth2.WithCreateTokenDeviceCode("<DEVICE_CODE>"),
    oauth2.WithCreateTokenClientId("<CLIENT_ID>"),
    oauth2.WithCreateTokenClientSecret("<CLIENT_SECRET>"),
    oauth2.WithCreateTokenCodeVerifier("<CODE_VERIFIER>"),
    oauth2.WithCreateTokenRedirectUri("https://example.com"),
    oauth2.WithCreateTokenResource(""),
    oauth2.WithCreateTokenAudience("<AUDIENCE>"),
)
```
