```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v2/client"
    "github.com/appwrite/sdk-for-go/v2/webhooks"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := webhooks.New(client)

response, error := service.Update(
    "<WEBHOOK_ID>",
    "<NAME>",
    "",
    []interface{}{},
    webhooks.WithUpdateEnabled(false),
    webhooks.WithUpdateSecurity(false),
    webhooks.WithUpdateHttpUser("<HTTP_USER>"),
    webhooks.WithUpdateHttpPass("<HTTP_PASS>"),
)
```
