```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/presences"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := presences.New(client)

	response, err := service.Update(
		"<PRESENCE_ID>",
		"<USER_ID>",
		service.WithUpdateStatus("<STATUS>"),
		service.WithUpdateExpiresAt("2020-10-15T06:38:00.000+00:00"),
		service.WithUpdateMetadata([]interface{}{}),
		service.WithUpdatePermissions([]string{"read(\"any\")"}),
		service.WithUpdatePurge(false),
	)
	fmt.Println(response, err)
}
```
