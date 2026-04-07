```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v2/client"
    "github.com/appwrite/sdk-for-go/v2/backups"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := backups.New(client)

response, error := service.CreatePolicy(
    "<POLICY_ID>",
    []interface{}{},
    1,
    "",
    backups.WithCreatePolicyName("<NAME>"),
    backups.WithCreatePolicyResourceId("<RESOURCE_ID>"),
    backups.WithCreatePolicyEnabled(false),
)
```
