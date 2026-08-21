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

	response, err := service.UpdateStringAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"",
		false,
		"<DEFAULT>",
		service.WithUpdateStringAttributeSize(1),
		service.WithUpdateStringAttributeNewKey(""),
	)
	fmt.Println(response, err)
}
```
