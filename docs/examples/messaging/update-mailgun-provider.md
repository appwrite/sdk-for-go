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
    response, error := messaging.UpdateMailgunProvider(
        "{$example}",
        messaging.WithUpdateMailgunProviderName("{$example}"),
        messaging.WithUpdateMailgunProviderApiKey("{$example}"),
        messaging.WithUpdateMailgunProviderDomain("{$example}"),
        messaging.WithUpdateMailgunProviderIsEuRegion(false),
        messaging.WithUpdateMailgunProviderEnabled(false),
        messaging.WithUpdateMailgunProviderFromName("{$example}"),
        messaging.WithUpdateMailgunProviderFromEmail("{$example}"),
        messaging.WithUpdateMailgunProviderReplyToName("{$example}"),
        messaging.WithUpdateMailgunProviderReplyToEmail("{$example}"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
