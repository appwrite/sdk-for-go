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

response, error := service.UpdateOAuth2Linkedin(
    project.WithUpdateOAuth2LinkedinClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2LinkedinPrimaryClientSecret("<PRIMARY_CLIENT_SECRET>"),
    project.WithUpdateOAuth2LinkedinEnabled(false),
)
```
