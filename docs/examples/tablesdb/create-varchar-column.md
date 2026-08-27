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

	response, err := service.CreateVarcharColumn(
		"<DATABASE_ID>",
		"<TABLE_ID>",
		"<KEY>",
		1,
		false,
		service.WithCreateVarcharColumnDefault("Hello World"),
		service.WithCreateVarcharColumnArray(false),
		service.WithCreateVarcharColumnEncrypt(false),
	)
	fmt.Println(response, err)
}
```
