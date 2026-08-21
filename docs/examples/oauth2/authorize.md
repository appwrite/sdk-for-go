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

	response, err := service.Authorize(
		service.WithAuthorizeClientId("<CLIENT_ID>"),
		service.WithAuthorizeRedirectUri("https://example.com"),
		service.WithAuthorizeResponseType(""),
		service.WithAuthorizeScope("<SCOPE>"),
		service.WithAuthorizeState("<STATE>"),
		service.WithAuthorizeNonce("<NONCE>"),
		service.WithAuthorizeCodeChallenge("<CODE_CHALLENGE>"),
		service.WithAuthorizeCodeChallengeMethod("s256"),
		service.WithAuthorizePrompt("<PROMPT>"),
		service.WithAuthorizeMaxAge(0),
		service.WithAuthorizeAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
		service.WithAuthorizeResource(""),
		service.WithAuthorizeAudience("<AUDIENCE>"),
		service.WithAuthorizeRequestUri("<REQUEST_URI>"),
	)
	fmt.Println(response, err)
}
```
