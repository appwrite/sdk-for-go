```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v2/client"
    "github.com/appwrite/sdk-for-go/v2/health"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := health.New(client)

response, error := service.GetQueueWebhooks(
    health.WithGetQueueWebhooksThreshold(0),
)
```
