```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/apps"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithSession("")
)

service := apps.New(client)

response, error := service.Create(
    "<APP_ID>",
    "<NAME>",
    []string{},
    apps.WithCreateDescription("<DESCRIPTION>"),
    apps.WithCreateClientUri("https://example.com"),
    apps.WithCreateLogoUri("https://example.com"),
    apps.WithCreatePrivacyPolicyUrl("https://example.com"),
    apps.WithCreateTermsUrl("https://example.com"),
    apps.WithCreateContacts([]string{}),
    apps.WithCreateTagline("<TAGLINE>"),
    apps.WithCreateTags([]string{}),
    apps.WithCreateImages([]string{}),
    apps.WithCreateSupportUrl("https://example.com"),
    apps.WithCreateDataDeletionUrl("https://example.com"),
    apps.WithCreatePostLogoutRedirectUris([]string{}),
    apps.WithCreateEnabled(false),
    apps.WithCreateType("public"),
    apps.WithCreateDeviceFlow(false),
    apps.WithCreateTeamId("<TEAM_ID>"),
)
```
