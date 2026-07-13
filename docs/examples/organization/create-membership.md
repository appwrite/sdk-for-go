```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v6/client"
    "github.com/appwrite/sdk-for-go/v6/organization"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := organization.New(client)

response, error := service.CreateMembership(
    []string{},
    organization.WithCreateMembershipEmail("email@example.com"),
    organization.WithCreateMembershipUserId("<USER_ID>"),
    organization.WithCreateMembershipPhone("+12065550100"),
    organization.WithCreateMembershipUrl("https://example.com"),
    organization.WithCreateMembershipName("<NAME>"),
)
```
