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

	response, err := service.CreateSMTPProvider(
		"<PROVIDER_ID>",
		"<NAME>",
		"<HOST>",
		service.WithCreateSMTPProviderPort(587),
		service.WithCreateSMTPProviderUsername("<USERNAME>"),
		service.WithCreateSMTPProviderPassword("password"),
		service.WithCreateSMTPProviderEncryption("none"),
		service.WithCreateSMTPProviderAutoTLS(false),
		service.WithCreateSMTPProviderMailer("<MAILER>"),
		service.WithCreateSMTPProviderFromName("<FROM_NAME>"),
		service.WithCreateSMTPProviderFromEmail("email@example.com"),
		service.WithCreateSMTPProviderReplyToName("<REPLY_TO_NAME>"),
		service.WithCreateSMTPProviderReplyToEmail("email@example.com"),
		service.WithCreateSMTPProviderEnabled(false),
	)
	fmt.Println(response, err)
}
```
