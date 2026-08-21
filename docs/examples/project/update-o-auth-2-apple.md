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

	response, err := service.UpdateOAuth2Apple(
		service.WithUpdateOAuth2AppleServiceId("<SERVICE_ID>"),
		service.WithUpdateOAuth2AppleKeyId("<KEY_ID>"),
		service.WithUpdateOAuth2AppleTeamId("<TEAM_ID>"),
		service.WithUpdateOAuth2AppleP8File("<P8_FILE>"),
		service.WithUpdateOAuth2AppleEnabled(false),
	)
	fmt.Println(response, err)
}
```
