```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/account"
	"github.com/appwrite/sdk-for-go/v7/client"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := account.New(client)

	response, err := service.CreateOAuth2Token(
		"amazon",
		account.WithCreateOAuth2TokenSuccess("https://example.com"),
		account.WithCreateOAuth2TokenFailure("https://example.com"),
		account.WithCreateOAuth2TokenScopes([]string{}),
	)
	fmt.Println(response, err)
}
```
