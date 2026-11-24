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

	response, err := service.UpdateBirdSMSProvider(
		"<PROVIDER_ID>",
		service.WithUpdateBirdSMSProviderName("<NAME>"),
		service.WithUpdateBirdSMSProviderEnabled(false),
		service.WithUpdateBirdSMSProviderApiKey("<API_KEY>"),
		service.WithUpdateBirdSMSProviderFrom("<FROM>"),
	)
	fmt.Println(response, err)
}
```
