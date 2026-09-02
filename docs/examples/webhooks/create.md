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

	response, err := service.Create(
		"<WEBHOOK_ID>",
		"https://example.com/webhook",
		"<NAME>",
		[]string{"example"},
		service.WithCreateEnabled(false),
		service.WithCreateTls(false),
		service.WithCreateAuthUsername("<AUTH_USERNAME>"),
		service.WithCreateAuthPassword("password"),
		service.WithCreateSecret("<SECRET>"),
	)
	fmt.Println(response, err)
}
```
