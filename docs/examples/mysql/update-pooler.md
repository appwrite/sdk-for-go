```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/mysql"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := mysql.New(client)

	response, err := service.UpdatePooler(
		"<DATABASE_ID>",
		service.WithUpdatePoolerMode("transaction"),
		service.WithUpdatePoolerMaxConnections(10),
		service.WithUpdatePoolerDefaultPoolSize(1),
		service.WithUpdatePoolerReadWriteSplitting(false),
		service.WithUpdatePoolerPoolerCpuRequest("<POOLER_CPU_REQUEST>"),
		service.WithUpdatePoolerPoolerCpuLimit("<POOLER_CPU_LIMIT>"),
		service.WithUpdatePoolerPoolerMemoryRequest("<POOLER_MEMORY_REQUEST>"),
		service.WithUpdatePoolerPoolerMemoryLimit("<POOLER_MEMORY_LIMIT>"),
	)
	fmt.Println(response, err)
}
```
