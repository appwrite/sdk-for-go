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

	response, err := service.CreateSesProvider(
		"<PROVIDER_ID>",
		"<NAME>",
		service.WithCreateSesProviderAccessKey("<ACCESS_KEY>"),
		service.WithCreateSesProviderSecretKey("<SECRET_KEY>"),
		service.WithCreateSesProviderRegion("<REGION>"),
		service.WithCreateSesProviderFromName("<FROM_NAME>"),
		service.WithCreateSesProviderFromEmail("email@example.com"),
		service.WithCreateSesProviderReplyToName("<REPLY_TO_NAME>"),
		service.WithCreateSesProviderReplyToEmail("email@example.com"),
		service.WithCreateSesProviderEnabled(false),
	)
	fmt.Println(response, err)
}
```
