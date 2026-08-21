```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/mongo"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := mongo.New(client)

	response, err := service.GetBackupPolicy(
		"<DATABASE_ID>",
		"<POLICY_ID>",
	)
	fmt.Println(response, err)
}
```
