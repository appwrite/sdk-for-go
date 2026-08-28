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

	response, err := service.CreateBigIntAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<KEY>",
		false,
		service.WithCreateBigIntAttributeMin(0),
		service.WithCreateBigIntAttributeMax(1000000),
		service.WithCreateBigIntAttributeDefault(0),
		service.WithCreateBigIntAttributeArray(false),
	)
	fmt.Println(response, err)
}
```
