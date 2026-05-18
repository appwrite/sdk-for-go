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

response, error := service.UpdateOAuth2Disqus(
    project.WithUpdateOAuth2DisqusPublicKey("<PUBLIC_KEY>"),
    project.WithUpdateOAuth2DisqusSecretKey("<SECRET_KEY>"),
    project.WithUpdateOAuth2DisqusEnabled(false),
)
```
