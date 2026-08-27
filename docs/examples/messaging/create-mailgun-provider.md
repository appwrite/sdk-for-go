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

	response, err := service.CreateMailgunProvider(
		"<PROVIDER_ID>",
		"<NAME>",
		service.WithCreateMailgunProviderApiKey("<API_KEY>"),
		service.WithCreateMailgunProviderDomain("example.com"),
		service.WithCreateMailgunProviderIsEuRegion(false),
		service.WithCreateMailgunProviderFromName("<FROM_NAME>"),
		service.WithCreateMailgunProviderFromEmail("email@example.com"),
		service.WithCreateMailgunProviderReplyToName("<REPLY_TO_NAME>"),
		service.WithCreateMailgunProviderReplyToEmail("email@example.com"),
		service.WithCreateMailgunProviderEnabled(false),
	)
	fmt.Println(response, err)
}
```
