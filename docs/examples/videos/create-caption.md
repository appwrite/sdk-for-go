```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/videos"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := videos.New(client)

	response, err := service.CreateCaption(
		"<VIDEO_ID>",
		"<BUCKET_ID>",
		"<FILE_ID>",
		"<NAME>",
		"aar",
		service.WithCreateCaptionDefault(false),
	)
	fmt.Println(response, err)
}
```
