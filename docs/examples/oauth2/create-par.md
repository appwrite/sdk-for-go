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

	response, err := service.CreatePAR(
		"<CLIENT_ID>",
		"https://example.com",
		"code",
		service.WithCreatePARScope("<SCOPE>"),
		service.WithCreatePARState("<STATE>"),
		service.WithCreatePARNonce("<NONCE>"),
		service.WithCreatePARCodeChallenge("<CODE_CHALLENGE>"),
		service.WithCreatePARCodeChallengeMethod("s256"),
		service.WithCreatePARPrompt("<PROMPT>"),
		service.WithCreatePARMaxAge(0),
		service.WithCreatePARAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
		service.WithCreatePARResource(""),
		service.WithCreatePARAudience("<AUDIENCE>"),
	)
	fmt.Println(response, err)
}
```
