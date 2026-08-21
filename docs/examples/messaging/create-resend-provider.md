```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/messaging"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := messaging.New(client)

	response, err := service.CreateResendProvider(
		"<PROVIDER_ID>",
		"<NAME>",
		service.WithCreateResendProviderApiKey("<API_KEY>"),
		service.WithCreateResendProviderFromName("<FROM_NAME>"),
		service.WithCreateResendProviderFromEmail("email@example.com"),
		service.WithCreateResendProviderReplyToName("<REPLY_TO_NAME>"),
		service.WithCreateResendProviderReplyToEmail("email@example.com"),
		service.WithCreateResendProviderEnabled(false),
	)
	fmt.Println(response, err)
}
```
