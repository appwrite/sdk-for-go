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

	response, err := service.CreateRestoration(
		"<DATABASE_ID>",
		service.WithCreateRestorationType("backup"),
		service.WithCreateRestorationBackupId("<BACKUP_ID>"),
		service.WithCreateRestorationTargetDatabaseId("<TARGET_DATABASE_ID>"),
		service.WithCreateRestorationTargetTime("2020-10-15T06:38:00.000+00:00"),
	)
	fmt.Println(response, err)
}
```
