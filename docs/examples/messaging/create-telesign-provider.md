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

	response, err := service.CreateTelesignProvider(
		"<PROVIDER_ID>",
		"<NAME>",
		service.WithCreateTelesignProviderFrom("+12065550100"),
		service.WithCreateTelesignProviderCustomerId("<CUSTOMER_ID>"),
		service.WithCreateTelesignProviderApiKey("<API_KEY>"),
		service.WithCreateTelesignProviderEnabled(false),
	)
	fmt.Println(response, err)
}
```
