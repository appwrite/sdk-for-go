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

	response, err := service.CreateAppwriteProvider(
		"<PROVIDER_ID>",
		"<NAME>",
		service.WithCreateAppwriteProviderEnabled(false),
		service.WithCreateAppwriteProviderQos(0),
		service.WithCreateAppwriteProviderExpiry(0),
	)
	fmt.Println(response, err)
}
```
