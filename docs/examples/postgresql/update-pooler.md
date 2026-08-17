```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/postgresql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := postgresql.New(client)

response, error := service.UpdatePooler(
    "<DATABASE_ID>",
    postgresql.WithUpdatePoolerMode("transaction"),
    postgresql.WithUpdatePoolerMaxConnections(10),
    postgresql.WithUpdatePoolerDefaultPoolSize(1),
    postgresql.WithUpdatePoolerReadWriteSplitting(false),
    postgresql.WithUpdatePoolerPoolerCpuRequest("<POOLER_CPU_REQUEST>"),
    postgresql.WithUpdatePoolerPoolerCpuLimit("<POOLER_CPU_LIMIT>"),
    postgresql.WithUpdatePoolerPoolerMemoryRequest("<POOLER_MEMORY_REQUEST>"),
    postgresql.WithUpdatePoolerPoolerMemoryLimit("<POOLER_MEMORY_LIMIT>"),
)
```
