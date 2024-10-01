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
    response, error := messaging.CreateSendgridProvider(
        "{$example}",
        "{$example}",
        messaging.WithCreateSendgridProviderApiKey("{$example}"),
        messaging.WithCreateSendgridProviderFromName("{$example}"),
        messaging.WithCreateSendgridProviderFromEmail("{$example}"),
        messaging.WithCreateSendgridProviderReplyToName("{$example}"),
        messaging.WithCreateSendgridProviderReplyToEmail("{$example}"),
        messaging.WithCreateSendgridProviderEnabled(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
