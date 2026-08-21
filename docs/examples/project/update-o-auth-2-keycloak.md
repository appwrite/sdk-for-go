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

	response, err := service.UpdateOAuth2Keycloak(
		service.WithUpdateOAuth2KeycloakClientId("<CLIENT_ID>"),
		service.WithUpdateOAuth2KeycloakClientSecret("<CLIENT_SECRET>"),
		service.WithUpdateOAuth2KeycloakEndpoint("<ENDPOINT>"),
		service.WithUpdateOAuth2KeycloakRealmName("<REALM_NAME>"),
		service.WithUpdateOAuth2KeycloakEnabled(false),
	)
	fmt.Println(response, err)
}
```
