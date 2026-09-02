```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/backups"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := backups.New(client)

	response, err := service.CreatePolicy(
		"<POLICY_ID>",
		[]string{"example"},
		1,
		"",
		service.WithCreatePolicyName("<NAME>"),
		service.WithCreatePolicyResourceId("<RESOURCE_ID>"),
		service.WithCreatePolicyEnabled(false),
	)
	fmt.Println(response, err)
}
```
