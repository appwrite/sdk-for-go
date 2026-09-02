```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/embeddings"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := embeddings.New(client)

	response, err := service.CreateTextEmbeddings(
		[]string{"example"},
		service.WithCreateTextEmbeddingsModel("nomic-embed-text"),
	)
	fmt.Println(response, err)
}
```
