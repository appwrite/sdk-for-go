```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v4/client"
    "github.com/appwrite/sdk-for-go/v4/proxy"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := proxy.New(client)

response, error := service.ListRules(
    proxy.WithListRulesQueries([]interface{}{}),
    proxy.WithListRulesTotal(false),
)
```
