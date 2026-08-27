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

	response, err := service.UpdateLineAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<KEY>",
		false,
		service.WithUpdateLineAttributeDefault([][]interface{}{[]interface{}{1, 2}, []interface{}{3, 4}, []interface{}{5, 6}}),
		service.WithUpdateLineAttributeNewKey("<NEW_KEY>"),
	)
	fmt.Println(response, err)
}
```
