```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/account"
	"github.com/appwrite/sdk-for-go/v7/appwrite"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := account.New(client)

	response, err := service.CreateMagicURLToken(
		"<USER_ID>",
		"email@example.com",
		service.WithCreateMagicURLTokenUrl("https://example.com"),
		service.WithCreateMagicURLTokenPhrase(false),
	)
	fmt.Println(response, err)
}
```
