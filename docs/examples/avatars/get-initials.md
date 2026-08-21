```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/avatars"
	"github.com/appwrite/sdk-for-go/v7/client"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := avatars.New(client)

	response, err := service.GetInitials(
		avatars.WithGetInitialsName("<NAME>"),
		avatars.WithGetInitialsWidth(0),
		avatars.WithGetInitialsHeight(0),
		avatars.WithGetInitialsBackground(""),
	)
	fmt.Println(response, err)
}
```
