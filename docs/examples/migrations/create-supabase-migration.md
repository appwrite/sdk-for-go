```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/migrations"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := migrations.New(client)

	response, err := service.CreateSupabaseMigration(
		[]string{"user"},
		"https://example.com",
		"<API_KEY>",
		"<DATABASE_HOST>",
		"<USERNAME>",
		"password",
		service.WithCreateSupabaseMigrationPort(5432),
	)
	fmt.Println(response, err)
}
```
