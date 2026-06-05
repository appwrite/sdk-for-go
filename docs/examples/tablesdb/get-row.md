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
    client.WithSession("")
)

service := tablesdb.New(client)

response, error := service.GetRow(
    "<DATABASE_ID>",
    "<TABLE_ID>",
    "<ROW_ID>",
    tablesdb.WithGetRowQueries([]string{}),
    tablesdb.WithGetRowTransactionId("<TRANSACTION_ID>"),
)
```
