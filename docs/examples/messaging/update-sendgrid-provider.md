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
    response, error := messaging.UpdateSendgridProvider(
        "{$example}",
        messaging.WithUpdateSendgridProviderName("{$example}"),
        messaging.WithUpdateSendgridProviderEnabled(false),
        messaging.WithUpdateSendgridProviderApiKey("{$example}"),
        messaging.WithUpdateSendgridProviderFromName("{$example}"),
        messaging.WithUpdateSendgridProviderFromEmail("{$example}"),
        messaging.WithUpdateSendgridProviderReplyToName("{$example}"),
        messaging.WithUpdateSendgridProviderReplyToEmail("{$example}"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
