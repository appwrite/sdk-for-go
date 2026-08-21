```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/mongo"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := mongo.New(client)

	response, err := service.UpdateBackupStorage(
		"<DATABASE_ID>",
		"s3",
		"<BUCKET>",
		"<ACCESS_KEY>",
		"<SECRET_KEY>",
		service.WithUpdateBackupStorageRegion("<REGION>"),
		service.WithUpdateBackupStoragePrefix("<PREFIX>"),
		service.WithUpdateBackupStorageEndpoint("<ENDPOINT>"),
	)
	fmt.Println(response, err)
}
```
