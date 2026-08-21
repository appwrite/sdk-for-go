```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/users"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := users.New(client)

	response, err := service.CreateToken(
		"<USER_ID>",
		service.WithCreateTokenLength(4),
		service.WithCreateTokenExpire(60),
	)
	fmt.Println(response, err)
}
```
