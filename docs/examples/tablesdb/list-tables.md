```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v3/client"
    "github.com/appwrite/sdk-for-go/v3/tablesdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := tablesdb.New(client)

response, error := service.ListTables(
    "<DATABASE_ID>",
    tablesdb.WithListTablesQueries([]interface{}{}),
    tablesdb.WithListTablesSearch("<SEARCH>"),
    tablesdb.WithListTablesTotal(false),
)
```
