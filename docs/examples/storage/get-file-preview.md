```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/storage"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := storage.New(client)

	response, err := service.GetFilePreview(
		"<BUCKET_ID>",
		"<FILE_ID>",
		service.WithGetFilePreviewWidth(0),
		service.WithGetFilePreviewHeight(0),
		service.WithGetFilePreviewGravity("center"),
		service.WithGetFilePreviewQuality(-1),
		service.WithGetFilePreviewBorderWidth(0),
		service.WithGetFilePreviewBorderColor("FFFFFF"),
		service.WithGetFilePreviewBorderRadius(0),
		service.WithGetFilePreviewOpacity(0),
		service.WithGetFilePreviewRotation(-360),
		service.WithGetFilePreviewBackground("FFFFFF"),
		service.WithGetFilePreviewOutput("jpg"),
		service.WithGetFilePreviewToken("<TOKEN>"),
	)
	fmt.Println(response, err)
}
```
