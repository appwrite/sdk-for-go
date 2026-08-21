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

	response, err := service.CreateTable(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<NAME>",
		service.WithCreateTablePermissions([]string{"read(\"any\")"}),
		service.WithCreateTableRowSecurity(false),
		service.WithCreateTableEnabled(false),
		service.WithCreateTableColumns([]interface{}{}),
		service.WithCreateTableIndexes([]interface{}{}),
	)
	fmt.Println(response, err)
}
```
