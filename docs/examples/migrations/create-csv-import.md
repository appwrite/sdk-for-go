```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/migrations"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := migrations.New(client)

	response, err := service.CreateCSVImport(
		"<BUCKET_ID>",
		"<FILE_ID>",
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		service.WithCreateCSVImportInternalFile(false),
		service.WithCreateCSVImportOnDuplicate("fail"),
	)
	fmt.Println(response, err)
}
```
