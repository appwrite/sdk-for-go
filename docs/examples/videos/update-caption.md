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

	response, err := service.UpdateCaption(
		"<VIDEO_ID>",
		"<CAPTION_ID>",
		service.WithUpdateCaptionBucketId("<BUCKET_ID>"),
		service.WithUpdateCaptionFileId("<FILE_ID>"),
		service.WithUpdateCaptionName("<NAME>"),
		service.WithUpdateCaptionLanguage("aar"),
		service.WithUpdateCaptionDefault(false),
	)
	fmt.Println(response, err)
}
```
