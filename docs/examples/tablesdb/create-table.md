```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v6/client"
    "github.com/appwrite/sdk-for-go/v6/tablesdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := tablesdb.New(client)

response, error := service.CreateTable(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    "<NAME>",
    tablesdb.WithCreateTablePermissions([]string{"read("any")"}),
    tablesdb.WithCreateTableRowSecurity(false),
    tablesdb.WithCreateTableEnabled(false),
    tablesdb.WithCreateTableColumns([]interface{}{}),
    tablesdb.WithCreateTableIndexes([]interface{}{}),
)
```
