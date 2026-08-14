```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/embeddings"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := embeddings.New(client)

response, error := service.CreateTextEmbeddings(
    []string{},
    embeddings.WithCreateTextEmbeddingsModel("nomic-embed-text"),
)
```
