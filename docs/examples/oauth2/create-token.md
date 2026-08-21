```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/oauth2"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithSession(""),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
	)

	service := oauth2.New(client)

	response, err := service.CreateToken(
		"<GRANT_TYPE>",
		service.WithCreateTokenCode("<CODE>"),
		service.WithCreateTokenRefreshToken("<REFRESH_TOKEN>"),
		service.WithCreateTokenDeviceCode("<DEVICE_CODE>"),
		service.WithCreateTokenClientId("<CLIENT_ID>"),
		service.WithCreateTokenClientSecret("<CLIENT_SECRET>"),
		service.WithCreateTokenCodeVerifier("<CODE_VERIFIER>"),
		service.WithCreateTokenRedirectUri("https://example.com"),
		service.WithCreateTokenResource(""),
		service.WithCreateTokenAudience("<AUDIENCE>"),
	)
	fmt.Println(response, err)
}
```
