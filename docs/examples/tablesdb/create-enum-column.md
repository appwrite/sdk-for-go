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

	response, err := service.CreateEnumColumn(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<KEY>",
		[]string{"active", "inactive"},
		false,
		service.WithCreateEnumColumnDefault("active"),
		service.WithCreateEnumColumnArray(false),
	)
	fmt.Println(response, err)
}
```
