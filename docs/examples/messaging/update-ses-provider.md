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

	response, err := service.UpdateSesProvider(
		"<PROVIDER_ID>",
		service.WithUpdateSesProviderName("<NAME>"),
		service.WithUpdateSesProviderEnabled(false),
		service.WithUpdateSesProviderAccessKey("<ACCESS_KEY>"),
		service.WithUpdateSesProviderSecretKey("<SECRET_KEY>"),
		service.WithUpdateSesProviderRegion("<REGION>"),
		service.WithUpdateSesProviderFromName("<FROM_NAME>"),
		service.WithUpdateSesProviderFromEmail("email@example.com"),
		service.WithUpdateSesProviderReplyToName("<REPLY_TO_NAME>"),
		service.WithUpdateSesProviderReplyToEmail("<REPLY_TO_EMAIL>"),
	)
	fmt.Println(response, err)
}
```
