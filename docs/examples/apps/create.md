```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/apps"
	"github.com/appwrite/sdk-for-go/v7/appwrite"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := apps.New(client)

	response, err := service.Create(
		"<APP_ID>",
		"<NAME>",
		[]string{},
		service.WithCreateDescription("<DESCRIPTION>"),
		service.WithCreateClientUri("https://example.com"),
		service.WithCreateLogoUri("https://example.com"),
		service.WithCreatePrivacyPolicyUrl("https://example.com"),
		service.WithCreateTermsUrl("https://example.com"),
		service.WithCreateContacts([]string{}),
		service.WithCreateTagline("<TAGLINE>"),
		service.WithCreateTags([]string{}),
		service.WithCreateImages([]string{}),
		service.WithCreateSupportUrl("https://example.com"),
		service.WithCreateDataDeletionUrl("https://example.com"),
		service.WithCreatePostLogoutRedirectUris([]string{}),
		service.WithCreateEnabled(false),
		service.WithCreateType("public"),
		service.WithCreateDeviceFlow(false),
		service.WithCreateTeamId("<TEAM_ID>"),
	)
	fmt.Println(response, err)
}
```
