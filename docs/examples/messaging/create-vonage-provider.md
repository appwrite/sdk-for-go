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

	response, err := service.CreateVonageProvider(
		"<PROVIDER_ID>",
		"<NAME>",
		messaging.WithCreateVonageProviderFrom("+12065550100"),
		messaging.WithCreateVonageProviderApiKey("<API_KEY>"),
		messaging.WithCreateVonageProviderApiSecret("<API_SECRET>"),
		messaging.WithCreateVonageProviderEnabled(false),
	)
	fmt.Println(response, err)
}
```
