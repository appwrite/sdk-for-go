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

	response, err := service.CreateRelationshipAttribute(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<RELATED_COLLECTION_ID>",
		"oneToOne",
		service.WithCreateRelationshipAttributeTwoWay(false),
		service.WithCreateRelationshipAttributeKey("<KEY>"),
		service.WithCreateRelationshipAttributeTwoWayKey("<TWO_WAY_KEY>"),
		service.WithCreateRelationshipAttributeOnDelete("cascade"),
	)
	fmt.Println(response, err)
}
```
