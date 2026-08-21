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

	response, err := service.UpdateSendgridProvider(
		"<PROVIDER_ID>",
		service.WithUpdateSendgridProviderName("<NAME>"),
		service.WithUpdateSendgridProviderEnabled(false),
		service.WithUpdateSendgridProviderApiKey("<API_KEY>"),
		service.WithUpdateSendgridProviderFromName("<FROM_NAME>"),
		service.WithUpdateSendgridProviderFromEmail("email@example.com"),
		service.WithUpdateSendgridProviderReplyToName("<REPLY_TO_NAME>"),
		service.WithUpdateSendgridProviderReplyToEmail("<REPLY_TO_EMAIL>"),
	)
	fmt.Println(response, err)
}
```
