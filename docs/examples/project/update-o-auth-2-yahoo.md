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

response, error := service.UpdateOAuth2Yahoo(
    project.WithUpdateOAuth2YahooClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2YahooClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2YahooEnabled(false),
)
```
