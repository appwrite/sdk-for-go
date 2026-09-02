```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/storage"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := storage.New(client)

	response, err := service.ListFiles(
		"<BUCKET_ID>",
		service.WithListFilesQueries([]string{"example"}),
		service.WithListFilesSearch("<SEARCH>"),
		service.WithListFilesTotal(false),
	)
	fmt.Println(response, err)
}
```
