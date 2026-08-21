```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/project"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := project.New(client)

	response, err := service.UpdateOAuth2Box(
		service.WithUpdateOAuth2BoxClientId("<CLIENT_ID>"),
		service.WithUpdateOAuth2BoxClientSecret("<CLIENT_SECRET>"),
		service.WithUpdateOAuth2BoxEnabled(false),
	)
	fmt.Println(response, err)
}
```
