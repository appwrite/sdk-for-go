```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v3/client"
    "github.com/appwrite/sdk-for-go/v3/project"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := project.New(client)

response, error := service.UpdateOAuth2Oidc(
    project.WithUpdateOAuth2OidcClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2OidcClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2OidcWellKnownURL("https://example.com"),
    project.WithUpdateOAuth2OidcAuthorizationURL("https://example.com"),
    project.WithUpdateOAuth2OidcTokenURL("https://example.com"),
    project.WithUpdateOAuth2OidcUserInfoURL("https://example.com"),
    project.WithUpdateOAuth2OidcEnabled(false),
)
```
