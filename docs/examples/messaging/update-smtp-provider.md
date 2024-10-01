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
    response, error := messaging.UpdateSmtpProvider(
        "{$example}",
        messaging.WithUpdateSmtpProviderName("{$example}"),
        messaging.WithUpdateSmtpProviderHost("{$example}"),
        messaging.WithUpdateSmtpProviderPort(1),
        messaging.WithUpdateSmtpProviderUsername("{$example}"),
        messaging.WithUpdateSmtpProviderPassword("{$example}"),
        messaging.WithUpdateSmtpProviderEncryption("{$example}"),
        messaging.WithUpdateSmtpProviderAutoTLS(false),
        messaging.WithUpdateSmtpProviderMailer("{$example}"),
        messaging.WithUpdateSmtpProviderFromName("{$example}"),
        messaging.WithUpdateSmtpProviderFromEmail("{$example}"),
        messaging.WithUpdateSmtpProviderReplyToName("{$example}"),
        messaging.WithUpdateSmtpProviderReplyToEmail("{$example}"),
        messaging.WithUpdateSmtpProviderEnabled(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
