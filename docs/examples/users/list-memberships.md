```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v5/client"
    "github.com/appwrite/sdk-for-go/v5/users"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := users.New(client)

response, error := service.ListMemberships(
    "<USER_ID>",
    users.WithListMembershipsQueries([]string{}),
    users.WithListMembershipsSearch("<SEARCH>"),
    users.WithListMembershipsTotal(false),
)
```
