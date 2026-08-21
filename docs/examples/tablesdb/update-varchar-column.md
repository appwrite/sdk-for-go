```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/tablesdb"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := tablesdb.New(client)

	response, err := service.UpdateVarcharColumn(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"",
		false,
		"<DEFAULT>",
		tablesdb.WithUpdateVarcharColumnSize(1),
		tablesdb.WithUpdateVarcharColumnNewKey(""),
	)
	fmt.Println(response, err)
}
```
