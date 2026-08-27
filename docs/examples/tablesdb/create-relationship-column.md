```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/tablesdb"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := tablesdb.New(client)

	response, err := service.CreateRelationshipColumn(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<RELATED_TABLE_ID>",
		"oneToOne",
		service.WithCreateRelationshipColumnTwoWay(false),
		service.WithCreateRelationshipColumnKey("<KEY>"),
		service.WithCreateRelationshipColumnTwoWayKey("<TWO_WAY_KEY>"),
		service.WithCreateRelationshipColumnOnDelete("cascade"),
	)
	fmt.Println(response, err)
}
```
