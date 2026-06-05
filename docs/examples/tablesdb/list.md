```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v5/client"
    "github.com/appwrite/sdk-for-go/v5/tablesdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := tablesdb.New(client)

response, error := service.List(
    tablesdb.WithListQueries([]string{}),
    tablesdb.WithListSearch("<SEARCH>"),
    tablesdb.WithListTotal(false),
)
```
