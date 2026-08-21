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

	response, err := service.UpdateStringColumn(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"",
		false,
		"<DEFAULT>",
		service.WithUpdateStringColumnSize(1),
		service.WithUpdateStringColumnNewKey(""),
	)
	fmt.Println(response, err)
}
```
