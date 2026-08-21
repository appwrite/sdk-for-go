```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/backups"
	"github.com/appwrite/sdk-for-go/v7/client"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := backups.New(client)

	response, err := service.CreatePolicy(
		"<POLICY_ID>",
		[]string{},
		1,
		"",
		backups.WithCreatePolicyName("<NAME>"),
		backups.WithCreatePolicyResourceId("<RESOURCE_ID>"),
		backups.WithCreatePolicyEnabled(false),
	)
	fmt.Println(response, err)
}
```
