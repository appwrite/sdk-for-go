```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/vectorsdb"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithSession("")
)

service := vectorsdb.New(client)

response, error := service.CreateQuery(
    "<DATABASE_ID>",
    "<COLLECTION_ID>",
    vectorsdb.WithCreateQueryQueries([]string{}),
    vectorsdb.WithCreateQueryTransactionId("<TRANSACTION_ID>"),
    vectorsdb.WithCreateQueryTotal(false),
    vectorsdb.WithCreateQueryTtl(0),
)
```
