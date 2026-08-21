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

	response, err := service.AuthorizePost(
		service.WithAuthorizePostClientId("<CLIENT_ID>"),
		service.WithAuthorizePostRedirectUri("https://example.com"),
		service.WithAuthorizePostResponseType(""),
		service.WithAuthorizePostScope("<SCOPE>"),
		service.WithAuthorizePostState("<STATE>"),
		service.WithAuthorizePostNonce("<NONCE>"),
		service.WithAuthorizePostCodeChallenge("<CODE_CHALLENGE>"),
		service.WithAuthorizePostCodeChallengeMethod("s256"),
		service.WithAuthorizePostPrompt("<PROMPT>"),
		service.WithAuthorizePostMaxAge(0),
		service.WithAuthorizePostAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
		service.WithAuthorizePostResource(""),
		service.WithAuthorizePostAudience("<AUDIENCE>"),
		service.WithAuthorizePostRequestUri("<REQUEST_URI>"),
	)
	fmt.Println(response, err)
}
```
