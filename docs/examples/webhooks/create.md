```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v3/client"
    "github.com/appwrite/sdk-for-go/v3/webhooks"
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
    webhooks.WithCreateTls(false),
    webhooks.WithCreateAuthUsername("<AUTH_USERNAME>"),
    webhooks.WithCreateAuthPassword("<AUTH_PASSWORD>"),
    webhooks.WithCreateSecret("<SECRET>"),
)
```
