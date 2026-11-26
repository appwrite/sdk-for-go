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

	response, err := service.CreateVideoProfile(
		"<PROFILE_ID>",
		"<NAME>",
		32000,
		32000,
		16,
		16,
		service.WithCreateVideoProfileCodec("h264"),
	)
	fmt.Println(response, err)
}
```
