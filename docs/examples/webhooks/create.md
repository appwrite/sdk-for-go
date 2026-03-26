```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/webhooks"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := webhooks.New(client)

response, error := service.Create(
    "<WEBHOOK_ID>",
    "",
    "<NAME>",
    []interface{}{},
    webhooks.WithCreateEnabled(false),
    webhooks.WithCreateSecurity(false),
    webhooks.WithCreateHttpUser("<HTTP_USER>"),
    webhooks.WithCreateHttpPass("<HTTP_PASS>"),
)
```
