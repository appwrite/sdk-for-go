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

	response, err := service.CreateIntegerAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"",
		false,
		service.WithCreateIntegerAttributeMin(0),
		service.WithCreateIntegerAttributeMax(0),
		service.WithCreateIntegerAttributeDefault(0),
		service.WithCreateIntegerAttributeArray(false),
	)
	fmt.Println(response, err)
}
```
