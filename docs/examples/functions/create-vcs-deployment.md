```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/functions"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := functions.New(client)

	response, err := service.CreateVcsDeployment(
		"<FUNCTION_ID>",
		"branch",
		"<REFERENCE>",
		functions.WithCreateVcsDeploymentActivate(false),
	)
	fmt.Println(response, err)
}
```
