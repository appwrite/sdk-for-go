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

	response, err := service.GetThumbnail(
		"<VIDEO_ID>",
		"<THUMBNAIL_ID>",
		service.WithGetThumbnailWidth(0),
		service.WithGetThumbnailHeight(0),
		service.WithGetThumbnailOutput("jpg"),
	)
	fmt.Println(response, err)
}
```
