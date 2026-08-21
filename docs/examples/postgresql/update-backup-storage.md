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

	response, err := service.UpdateBackupStorage(
		"<DATABASE_ID>",
		"s3",
		"<BUCKET>",
		"<ACCESS_KEY>",
		"<SECRET_KEY>",
		postgresql.WithUpdateBackupStorageRegion("<REGION>"),
		postgresql.WithUpdateBackupStoragePrefix("<PREFIX>"),
		postgresql.WithUpdateBackupStorageEndpoint("<ENDPOINT>"),
	)
	fmt.Println(response, err)
}
```
