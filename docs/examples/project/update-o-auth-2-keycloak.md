```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v4/client"
    "github.com/appwrite/sdk-for-go/v4/project"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := project.New(client)

response, error := service.UpdateOAuth2Keycloak(
    project.WithUpdateOAuth2KeycloakClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2KeycloakClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2KeycloakEndpoint("<ENDPOINT>"),
    project.WithUpdateOAuth2KeycloakRealmName("<REALM_NAME>"),
    project.WithUpdateOAuth2KeycloakEnabled(false),
)
```
