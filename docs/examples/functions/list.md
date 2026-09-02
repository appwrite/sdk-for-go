```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/functions"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := functions.New(client)

	response, err := service.List(
		service.WithListQueries([]string{"example"}),
		service.WithListSearch("<SEARCH>"),
		service.WithListTotal(false),
	)
	fmt.Println(response, err)
}
```
