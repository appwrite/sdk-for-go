package main

import (
    "fmt"
	"github.com/appwrite/sdk-for-go/appwrite"
)

func main() {
	client := appwrite.NewClient(
        appwrite.WithEndpoint("https://cloud.appwrite.io/v1"), // Your API Endpoint
        appwrite.WithProject(""), // Your project ID
        appwrite.WithKey(""), // Your secret API key
    )

    messaging := appwrite.NewMessaging(client)
    response, error := messaging.CreateMailgunProvider(
        "{$example}",
        "{$example}",
        messaging.WithCreateMailgunProviderApiKey("{$example}"),
        messaging.WithCreateMailgunProviderDomain("{$example}"),
        messaging.WithCreateMailgunProviderIsEuRegion(false),
        messaging.WithCreateMailgunProviderFromName("{$example}"),
        messaging.WithCreateMailgunProviderFromEmail("{$example}"),
        messaging.WithCreateMailgunProviderReplyToName("{$example}"),
        messaging.WithCreateMailgunProviderReplyToEmail("{$example}"),
        messaging.WithCreateMailgunProviderEnabled(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
