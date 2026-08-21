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

	response, err := service.GetQR(
		"<TEXT>",
		service.WithGetQRSize(1),
		service.WithGetQRMargin(0),
		service.WithGetQRDownload(false),
	)
	fmt.Println(response, err)
}
```
