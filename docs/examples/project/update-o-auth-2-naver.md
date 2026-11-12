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

	response, err := service.UpdateOAuth2Naver(
		service.WithUpdateOAuth2NaverClientId("<CLIENT_ID>"),
		service.WithUpdateOAuth2NaverClientSecret("<CLIENT_SECRET>"),
		service.WithUpdateOAuth2NaverEnabled(false),
	)
	fmt.Println(response, err)
}
```
