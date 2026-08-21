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

	response, err := service.UpdateVonageProvider(
		"<PROVIDER_ID>",
		messaging.WithUpdateVonageProviderName("<NAME>"),
		messaging.WithUpdateVonageProviderEnabled(false),
		messaging.WithUpdateVonageProviderApiKey("<API_KEY>"),
		messaging.WithUpdateVonageProviderApiSecret("<API_SECRET>"),
		messaging.WithUpdateVonageProviderFrom("<FROM>"),
	)
	fmt.Println(response, err)
}
```
