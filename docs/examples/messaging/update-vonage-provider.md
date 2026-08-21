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

	response, err := service.UpdateVonageProvider(
		"<PROVIDER_ID>",
		service.WithUpdateVonageProviderName("<NAME>"),
		service.WithUpdateVonageProviderEnabled(false),
		service.WithUpdateVonageProviderApiKey("<API_KEY>"),
		service.WithUpdateVonageProviderApiSecret("<API_SECRET>"),
		service.WithUpdateVonageProviderFrom("<FROM>"),
	)
	fmt.Println(response, err)
}
```
