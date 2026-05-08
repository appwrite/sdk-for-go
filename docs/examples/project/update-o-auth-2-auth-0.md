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

response, error := service.UpdateOAuth2Auth0(
    project.WithUpdateOAuth2Auth0ClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2Auth0ClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2Auth0Endpoint("<ENDPOINT>"),
    project.WithUpdateOAuth2Auth0Enabled(false),
)
```
