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

	response, err := service.UpdateBigIntAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<KEY>",
		false,
		0,
		service.WithUpdateBigIntAttributeMin(0),
		service.WithUpdateBigIntAttributeMax(1000000),
		service.WithUpdateBigIntAttributeNewKey("<NEW_KEY>"),
	)
	fmt.Println(response, err)
}
```
