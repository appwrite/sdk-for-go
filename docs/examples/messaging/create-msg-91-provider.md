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

	response, err := service.CreateMsg91Provider(
		"<PROVIDER_ID>",
		"<NAME>",
		service.WithCreateMsg91ProviderTemplateId("<TEMPLATE_ID>"),
		service.WithCreateMsg91ProviderSenderId("<SENDER_ID>"),
		service.WithCreateMsg91ProviderAuthKey("<AUTH_KEY>"),
		service.WithCreateMsg91ProviderEnabled(false),
	)
	fmt.Println(response, err)
}
```
