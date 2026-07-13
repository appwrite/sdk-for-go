```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v6/client"
    "github.com/appwrite/sdk-for-go/v6/storage"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := storage.New(client)

response, error := service.ListBuckets(
    storage.WithListBucketsQueries([]string{}),
    storage.WithListBucketsSearch("<SEARCH>"),
    storage.WithListBucketsTotal(false),
)
```
