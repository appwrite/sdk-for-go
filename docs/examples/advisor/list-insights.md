```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/advisor"
	"github.com/appwrite/sdk-for-go/v7/client"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := advisor.New(client)

	response, err := service.ListInsights(
		"<REPORT_ID>",
		advisor.WithListInsightsQueries([]string{}),
		advisor.WithListInsightsTotal(false),
	)
	fmt.Println(response, err)
}
```
