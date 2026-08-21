```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/postgresql"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := postgresql.New(client)

	response, err := service.CreateBackup(
		"<DATABASE_ID>",
		postgresql.WithCreateBackupType("full"),
	)
	fmt.Println(response, err)
}
```
