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

	response, err := service.UpdateOAuth2Oidc(
		service.WithUpdateOAuth2OidcClientId("<CLIENT_ID>"),
		service.WithUpdateOAuth2OidcClientSecret("<CLIENT_SECRET>"),
		service.WithUpdateOAuth2OidcWellKnownURL("https://example.com"),
		service.WithUpdateOAuth2OidcAuthorizationURL("https://example.com"),
		service.WithUpdateOAuth2OidcTokenURL("https://example.com"),
		service.WithUpdateOAuth2OidcUserInfoURL("https://example.com"),
		service.WithUpdateOAuth2OidcPrompt([]string{"example"}),
		service.WithUpdateOAuth2OidcMaxAge(0),
		service.WithUpdateOAuth2OidcEnabled(false),
	)
	fmt.Println(response, err)
}
```
