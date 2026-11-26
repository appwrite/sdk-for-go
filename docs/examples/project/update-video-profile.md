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

	response, err := service.UpdateVideoProfile(
		"<PROFILE_ID>",
		service.WithUpdateVideoProfileName("<NAME>"),
		service.WithUpdateVideoProfileVideoBitRate(32000),
		service.WithUpdateVideoProfileAudioBitRate(32000),
		service.WithUpdateVideoProfileWidth(16),
		service.WithUpdateVideoProfileHeight(16),
	)
	fmt.Println(response, err)
}
```
