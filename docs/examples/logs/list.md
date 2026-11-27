```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/logs"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := logs.New(client)

	response, err := service.List(
		service.WithListResourceType("functions"),
		service.WithListResourceId("<RESOURCE_ID>"),
		service.WithListQueries([]string{}),
		service.WithListTotal(false),
	)
	fmt.Println(response, err)
}
```
