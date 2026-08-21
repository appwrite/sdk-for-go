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

	response, err := service.UpdateAPNSProvider(
		"<PROVIDER_ID>",
		service.WithUpdateAPNSProviderName("<NAME>"),
		service.WithUpdateAPNSProviderEnabled(false),
		service.WithUpdateAPNSProviderAuthKey("<AUTH_KEY>"),
		service.WithUpdateAPNSProviderAuthKeyId("<AUTH_KEY_ID>"),
		service.WithUpdateAPNSProviderTeamId("<TEAM_ID>"),
		service.WithUpdateAPNSProviderBundleId("<BUNDLE_ID>"),
		service.WithUpdateAPNSProviderSandbox(false),
	)
	fmt.Println(response, err)
}
```
