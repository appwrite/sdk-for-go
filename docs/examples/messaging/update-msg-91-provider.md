```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/messaging"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := messaging.New(client)

	response, err := service.UpdateMsg91Provider(
		"<PROVIDER_ID>",
		messaging.WithUpdateMsg91ProviderName("<NAME>"),
		messaging.WithUpdateMsg91ProviderEnabled(false),
		messaging.WithUpdateMsg91ProviderTemplateId("<TEMPLATE_ID>"),
		messaging.WithUpdateMsg91ProviderSenderId("<SENDER_ID>"),
		messaging.WithUpdateMsg91ProviderAuthKey("<AUTH_KEY>"),
	)
	fmt.Println(response, err)
}
```
