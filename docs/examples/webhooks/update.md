```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/webhooks"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := webhooks.New(client)

	response, err := service.Update(
		"<WEBHOOK_ID>",
		"<NAME>",
		"https://example.com/webhook",
		[]string{"example"},
		service.WithUpdateEnabled(false),
		service.WithUpdateTls(false),
		service.WithUpdateAuthUsername("<AUTH_USERNAME>"),
		service.WithUpdateAuthPassword("password"),
	)
	fmt.Println(response, err)
}
```
