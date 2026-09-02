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

	response, err := service.Update(
		"<APP_ID>",
		"<NAME>",
		service.WithUpdateDescription("<DESCRIPTION>"),
		service.WithUpdateClientUri("https://example.com"),
		service.WithUpdateLogoUri("https://example.com"),
		service.WithUpdatePrivacyPolicyUrl("https://example.com"),
		service.WithUpdateTermsUrl("https://example.com"),
		service.WithUpdateContacts([]string{"example"}),
		service.WithUpdateTagline("<TAGLINE>"),
		service.WithUpdateTags([]string{"example"}),
		service.WithUpdateImages([]string{"example"}),
		service.WithUpdateSupportUrl("https://example.com"),
		service.WithUpdateDataDeletionUrl("https://example.com"),
		service.WithUpdateEnabled(false),
		service.WithUpdateRedirectUris([]string{"example"}),
		service.WithUpdatePostLogoutRedirectUris([]string{"example"}),
		service.WithUpdateType("public"),
		service.WithUpdateDeviceFlow(false),
		service.WithUpdateInstallationScopes([]string{"example"}),
		service.WithUpdateInstallationRedirectUrl("https://example.com"),
	)
	fmt.Println(response, err)
}
```
