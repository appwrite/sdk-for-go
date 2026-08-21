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

	response, err := service.UpdateResendProvider(
		"<PROVIDER_ID>",
		service.WithUpdateResendProviderName("<NAME>"),
		service.WithUpdateResendProviderEnabled(false),
		service.WithUpdateResendProviderApiKey("<API_KEY>"),
		service.WithUpdateResendProviderFromName("<FROM_NAME>"),
		service.WithUpdateResendProviderFromEmail("email@example.com"),
		service.WithUpdateResendProviderReplyToName("<REPLY_TO_NAME>"),
		service.WithUpdateResendProviderReplyToEmail("<REPLY_TO_EMAIL>"),
	)
	fmt.Println(response, err)
}
```
