```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/vectorsdb"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := vectorsdb.New(client)

	response, err := service.CreateCollection(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<NAME>",
		1,
		service.WithCreateCollectionPermissions([]string{"read(\"any\")"}),
		service.WithCreateCollectionDocumentSecurity(false),
		service.WithCreateCollectionEnabled(false),
	)
	fmt.Println(response, err)
}
```
