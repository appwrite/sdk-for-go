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

	response, err := service.UpdateOAuth2Server(
		false,
		"https://example.com",
		service.WithUpdateOAuth2ServerScopes([]string{"example"}),
		service.WithUpdateOAuth2ServerAuthorizationDetailsTypes([]string{"example"}),
		service.WithUpdateOAuth2ServerAccessTokenDuration(60),
		service.WithUpdateOAuth2ServerRefreshTokenDuration(60),
		service.WithUpdateOAuth2ServerPublicAccessTokenDuration(60),
		service.WithUpdateOAuth2ServerPublicRefreshTokenDuration(60),
		service.WithUpdateOAuth2ServerInstallationAccessTokenDuration(60),
		service.WithUpdateOAuth2ServerConfidentialPkce(false),
		service.WithUpdateOAuth2ServerVerificationUrl("https://example.com"),
		service.WithUpdateOAuth2ServerUserCodeLength(6),
		service.WithUpdateOAuth2ServerUserCodeFormat("numeric"),
		service.WithUpdateOAuth2ServerDeviceCodeDuration(60),
		service.WithUpdateOAuth2ServerDefaultScopes([]string{"example"}),
		service.WithUpdateOAuth2ServerInstallationScopes([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
