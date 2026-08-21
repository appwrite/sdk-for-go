```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/postgresql"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := postgresql.New(client)

	response, err := service.ListRestorations(
		"<DATABASE_ID>",
		service.WithListRestorationsStatus("pending"),
		service.WithListRestorationsType("backup"),
		service.WithListRestorationsLimit(1),
		service.WithListRestorationsOffset(0),
	)
	fmt.Println(response, err)
}
```
