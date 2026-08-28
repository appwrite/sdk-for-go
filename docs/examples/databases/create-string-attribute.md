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

	response, err := service.CreateStringAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<KEY>",
		1,
		false,
		service.WithCreateStringAttributeDefault("Hello World"),
		service.WithCreateStringAttributeArray(false),
		service.WithCreateStringAttributeEncrypt(false),
	)
	fmt.Println(response, err)
}
```
