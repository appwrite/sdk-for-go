```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/organization"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithSession("")
)

service := organization.New(client)

response, error := service.UpdateInstallation(
    "<INSTALLATION_ID>",
    organization.WithUpdateInstallationAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
)
```
