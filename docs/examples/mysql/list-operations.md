```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/mysql"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := mysql.New(client)

	response, err := service.ListOperations(
		"<DATABASE_ID>",
		mysql.WithListOperationsStatus("running"),
		mysql.WithListOperationsLimit(1),
		mysql.WithListOperationsOffset(0),
	)
	fmt.Println(response, err)
}
```
