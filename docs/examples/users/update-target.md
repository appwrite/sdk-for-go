```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/users"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := users.New(client)

	response, err := service.UpdateTarget(
		"<USER_ID>",
		"<TARGET_ID>",
		service.WithUpdateTargetIdentifier("<IDENTIFIER>"),
		service.WithUpdateTargetProviderId("<PROVIDER_ID>"),
		service.WithUpdateTargetName("<NAME>"),
	)
	fmt.Println(response, err)
}
```
