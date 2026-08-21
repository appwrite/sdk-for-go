```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/project"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := project.New(client)

	response, err := service.UpdateOAuth2HuggingFace(
		project.WithUpdateOAuth2HuggingFaceClientId("<CLIENT_ID>"),
		project.WithUpdateOAuth2HuggingFaceClientSecret("<CLIENT_SECRET>"),
		project.WithUpdateOAuth2HuggingFaceEnabled(false),
	)
	fmt.Println(response, err)
}
```
