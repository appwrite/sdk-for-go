```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/avatars"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := avatars.New(client)

	response, err := service.GetPhoto(
		service.WithGetPhotoWidth(0),
		service.WithGetPhotoHeight(0),
		service.WithGetPhotoQuality(0),
		service.WithGetPhotoOutput("png"),
		service.WithGetPhotoRating("g"),
		service.WithGetPhotoUserId("current()"),
		service.WithGetPhotoEmailHash("<EMAIL_HASH>"),
		service.WithGetPhotoName("<NAME>"),
	)
	fmt.Println(response, err)
}
```
