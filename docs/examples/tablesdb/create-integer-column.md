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

	response, err := service.CreateIntegerColumn(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<KEY>",
		false,
		service.WithCreateIntegerColumnMin(0),
		service.WithCreateIntegerColumnMax(100),
		service.WithCreateIntegerColumnDefault(10),
		service.WithCreateIntegerColumnArray(false),
	)
	fmt.Println(response, err)
}
```
