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

	response, err := service.CreateSesProvider(
		"<PROVIDER_ID>",
		"<NAME>",
		messaging.WithCreateSesProviderAccessKey("<ACCESS_KEY>"),
		messaging.WithCreateSesProviderSecretKey("<SECRET_KEY>"),
		messaging.WithCreateSesProviderRegion("<REGION>"),
		messaging.WithCreateSesProviderFromName("<FROM_NAME>"),
		messaging.WithCreateSesProviderFromEmail("email@example.com"),
		messaging.WithCreateSesProviderReplyToName("<REPLY_TO_NAME>"),
		messaging.WithCreateSesProviderReplyToEmail("email@example.com"),
		messaging.WithCreateSesProviderEnabled(false),
	)
	fmt.Println(response, err)
}
```
