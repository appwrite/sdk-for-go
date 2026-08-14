```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/oauth2"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithSession("")
    client.WithProject("<YOUR_PROJECT_ID>")
)

service := oauth2.New(client)

response, error := service.Authorize(
    oauth2.WithAuthorizeClientId("<CLIENT_ID>"),
    oauth2.WithAuthorizeRedirectUri("https://example.com"),
    oauth2.WithAuthorizeResponseType(""),
    oauth2.WithAuthorizeScope("<SCOPE>"),
    oauth2.WithAuthorizeState("<STATE>"),
    oauth2.WithAuthorizeNonce("<NONCE>"),
    oauth2.WithAuthorizeCodeChallenge("<CODE_CHALLENGE>"),
    oauth2.WithAuthorizeCodeChallengeMethod("s256"),
    oauth2.WithAuthorizePrompt("<PROMPT>"),
    oauth2.WithAuthorizeMaxAge(0),
    oauth2.WithAuthorizeAuthorizationDetails("<AUTHORIZATION_DETAILS>"),
    oauth2.WithAuthorizeResource(""),
    oauth2.WithAuthorizeAudience("<AUDIENCE>"),
    oauth2.WithAuthorizeRequestUri("<REQUEST_URI>"),
)
```
