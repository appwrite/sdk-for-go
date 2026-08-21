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

	response, err := service.UpdateTextmagicProvider(
		"<PROVIDER_ID>",
		messaging.WithUpdateTextmagicProviderName("<NAME>"),
		messaging.WithUpdateTextmagicProviderEnabled(false),
		messaging.WithUpdateTextmagicProviderUsername("<USERNAME>"),
		messaging.WithUpdateTextmagicProviderApiKey("<API_KEY>"),
		messaging.WithUpdateTextmagicProviderFrom("<FROM>"),
	)
	fmt.Println(response, err)
}
```
