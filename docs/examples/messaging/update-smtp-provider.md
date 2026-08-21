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

	response, err := service.UpdateSMTPProvider(
		"<PROVIDER_ID>",
		service.WithUpdateSMTPProviderName("<NAME>"),
		service.WithUpdateSMTPProviderHost("<HOST>"),
		service.WithUpdateSMTPProviderPort(1),
		service.WithUpdateSMTPProviderUsername("<USERNAME>"),
		service.WithUpdateSMTPProviderPassword("password"),
		service.WithUpdateSMTPProviderEncryption("none"),
		service.WithUpdateSMTPProviderAutoTLS(false),
		service.WithUpdateSMTPProviderMailer("<MAILER>"),
		service.WithUpdateSMTPProviderFromName("<FROM_NAME>"),
		service.WithUpdateSMTPProviderFromEmail("email@example.com"),
		service.WithUpdateSMTPProviderReplyToName("<REPLY_TO_NAME>"),
		service.WithUpdateSMTPProviderReplyToEmail("<REPLY_TO_EMAIL>"),
		service.WithUpdateSMTPProviderEnabled(false),
	)
	fmt.Println(response, err)
}
```
