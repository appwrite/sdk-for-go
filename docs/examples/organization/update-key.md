```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v5/client"
    "github.com/appwrite/sdk-for-go/v5/organization"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := organization.New(client)

response, error := service.UpdateKey(
    "<KEY_ID>",
    "<NAME>",
    []string{},
    organization.WithUpdateKeyExpire("2020-10-15T06:38:00.000+00:00"),
)
```
