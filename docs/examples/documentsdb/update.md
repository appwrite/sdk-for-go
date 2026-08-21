```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/documentsdb"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := documentsdb.New(client)

	response, err := service.Update(
		"<DATABASE_ID>",
		"<NAME>",
		documentsdb.WithUpdateEnabled(false),
		documentsdb.WithUpdateSpecification("serverless"),
		documentsdb.WithUpdateReplicas(0),
		documentsdb.WithUpdateSyncMode("async"),
	)
	fmt.Println(response, err)
}
```
