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

response, error := service.UpdateOAuth2Authentik(
    project.WithUpdateOAuth2AuthentikClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2AuthentikClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2AuthentikEndpoint("<ENDPOINT>"),
    project.WithUpdateOAuth2AuthentikEnabled(false),
)
```
