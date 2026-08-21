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

	response, err := service.UpdateCollection(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<NAME>",
		service.WithUpdateCollectionDimension(1),
		service.WithUpdateCollectionPermissions([]string{"read(\"any\")"}),
		service.WithUpdateCollectionDocumentSecurity(false),
		service.WithUpdateCollectionEnabled(false),
	)
	fmt.Println(response, err)
}
```
