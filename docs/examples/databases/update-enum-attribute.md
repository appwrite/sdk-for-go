```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/databases"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := databases.New(client)

	response, err := service.UpdateEnumAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<KEY>",
		[]string{"active", "inactive"},
		false,
		"active",
		service.WithUpdateEnumAttributeNewKey("<NEW_KEY>"),
	)
	fmt.Println(response, err)
}
```
