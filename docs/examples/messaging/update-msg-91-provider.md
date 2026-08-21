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

	response, err := service.UpdateMsg91Provider(
		"<PROVIDER_ID>",
		service.WithUpdateMsg91ProviderName("<NAME>"),
		service.WithUpdateMsg91ProviderEnabled(false),
		service.WithUpdateMsg91ProviderTemplateId("<TEMPLATE_ID>"),
		service.WithUpdateMsg91ProviderSenderId("<SENDER_ID>"),
		service.WithUpdateMsg91ProviderAuthKey("<AUTH_KEY>"),
	)
	fmt.Println(response, err)
}
```
