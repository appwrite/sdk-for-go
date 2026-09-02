```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/proxy"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := proxy.New(client)

	response, err := service.ListRules(
		service.WithListRulesQueries([]string{"example"}),
		service.WithListRulesTotal(false),
	)
	fmt.Println(response, err)
}
```
