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

response, error := service.UpdateOAuth2Yandex(
    project.WithUpdateOAuth2YandexClientId("<CLIENT_ID>"),
    project.WithUpdateOAuth2YandexClientSecret("<CLIENT_SECRET>"),
    project.WithUpdateOAuth2YandexEnabled(false),
)
```
