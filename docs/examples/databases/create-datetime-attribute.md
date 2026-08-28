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

	response, err := service.CreateDatetimeAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<KEY>",
		false,
		service.WithCreateDatetimeAttributeDefault("2020-10-15T06:38:00.000+00:00"),
		service.WithCreateDatetimeAttributeArray(false),
	)
	fmt.Println(response, err)
}
```
