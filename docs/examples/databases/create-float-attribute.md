```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/databases"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := databases.New(client)

	response, err := service.CreateFloatAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"",
		false,
		databases.WithCreateFloatAttributeMin(0),
		databases.WithCreateFloatAttributeMax(0),
		databases.WithCreateFloatAttributeDefault(0),
		databases.WithCreateFloatAttributeArray(false),
	)
	fmt.Println(response, err)
}
```
