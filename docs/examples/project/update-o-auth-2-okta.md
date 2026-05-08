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

response, error := service.UpdateOAuth2Okta(
    project.WithUpdateOAuth2OktaClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2OktaClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2OktaDomain(""),
    project.WithUpdateOAuth2OktaAuthorizationServerId("<AUTHORIZATION_SERVER_ID>"),
    project.WithUpdateOAuth2OktaEnabled(false),
)
```
