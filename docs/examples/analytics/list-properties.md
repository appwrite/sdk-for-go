```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/analytics"
	"github.com/appwrite/sdk-for-go/v7/appwrite"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := analytics.New(client)

	response, err := service.ListProperties(
		service.WithListPropertiesQueries([]string{"example"}),
		service.WithListPropertiesSearch("<SEARCH>"),
		service.WithListPropertiesTotal(false),
	)
	fmt.Println(response, err)
}
```
