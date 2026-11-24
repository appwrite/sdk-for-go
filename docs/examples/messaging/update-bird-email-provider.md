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

	response, err := service.UpdateBirdEmailProvider(
		"<PROVIDER_ID>",
		service.WithUpdateBirdEmailProviderName("<NAME>"),
		service.WithUpdateBirdEmailProviderEnabled(false),
		service.WithUpdateBirdEmailProviderApiKey("<API_KEY>"),
		service.WithUpdateBirdEmailProviderFromName("<FROM_NAME>"),
		service.WithUpdateBirdEmailProviderFromEmail("email@example.com"),
		service.WithUpdateBirdEmailProviderReplyToName("<REPLY_TO_NAME>"),
		service.WithUpdateBirdEmailProviderReplyToEmail("email@example.com"),
	)
	fmt.Println(response, err)
}
```
