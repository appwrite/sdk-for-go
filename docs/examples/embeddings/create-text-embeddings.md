```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/embeddings"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := embeddings.New(client)

	response, err := service.CreateTextEmbeddings(
		[]string{},
		embeddings.WithCreateTextEmbeddingsModel("nomic-embed-text"),
	)
	fmt.Println(response, err)
}
```
