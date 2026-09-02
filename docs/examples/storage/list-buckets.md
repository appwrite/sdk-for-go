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
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := storage.New(client)

	response, err := service.ListBuckets(
		service.WithListBucketsQueries([]string{"example"}),
		service.WithListBucketsSearch("<SEARCH>"),
		service.WithListBucketsTotal(false),
	)
	fmt.Println(response, err)
}
```
