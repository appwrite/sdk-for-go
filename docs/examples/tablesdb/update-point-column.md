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

	response, err := service.UpdatePointColumn(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<KEY>",
		false,
		service.WithUpdatePointColumnDefault([]float64{1, 2}),
		service.WithUpdatePointColumnNewKey("<NEW_KEY>"),
	)
	fmt.Println(response, err)
}
```
