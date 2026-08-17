```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/mysql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := mysql.New(client)

response, error := service.UpdatePooler(
    "<DATABASE_ID>",
    mysql.WithUpdatePoolerMode("transaction"),
    mysql.WithUpdatePoolerMaxConnections(10),
    mysql.WithUpdatePoolerDefaultPoolSize(1),
    mysql.WithUpdatePoolerReadWriteSplitting(false),
    mysql.WithUpdatePoolerPoolerCpuRequest("<POOLER_CPU_REQUEST>"),
    mysql.WithUpdatePoolerPoolerCpuLimit("<POOLER_CPU_LIMIT>"),
    mysql.WithUpdatePoolerPoolerMemoryRequest("<POOLER_MEMORY_REQUEST>"),
    mysql.WithUpdatePoolerPoolerMemoryLimit("<POOLER_MEMORY_LIMIT>"),
)
```
