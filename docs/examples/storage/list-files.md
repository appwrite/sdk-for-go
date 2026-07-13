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
    client.WithSession("")
)

service := storage.New(client)

response, error := service.ListFiles(
    "<BUCKET_ID>",
    storage.WithListFilesQueries([]string{}),
    storage.WithListFilesSearch("<SEARCH>"),
    storage.WithListFilesTotal(false),
)
```
