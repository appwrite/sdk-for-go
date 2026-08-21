```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/oauth2"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithSession(""),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
	)

	service := oauth2.New(client)

	response, err := service.Approve(
		"<GRANT_ID>",
		service.WithApproveAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
		service.WithApproveScope("<SCOPE>"),
	)
	fmt.Println(response, err)
}
```
