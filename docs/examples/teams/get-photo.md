```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/teams"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := teams.New(client)

	response, err := service.GetPhoto(
		"<TEAM_ID>",
		service.WithGetPhotoWidth(0),
		service.WithGetPhotoHeight(0),
		service.WithGetPhotoQuality(0),
		service.WithGetPhotoOutput("png"),
	)
	fmt.Println(response, err)
}
```
