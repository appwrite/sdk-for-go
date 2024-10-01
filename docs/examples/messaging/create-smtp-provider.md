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
    response, error := messaging.CreateSmtpProvider(
        "{$example}",
        "{$example}",
        "{$example}",
        messaging.WithCreateSmtpProviderPort(1),
        messaging.WithCreateSmtpProviderUsername("{$example}"),
        messaging.WithCreateSmtpProviderPassword("{$example}"),
        messaging.WithCreateSmtpProviderEncryption("{$example}"),
        messaging.WithCreateSmtpProviderAutoTLS(false),
        messaging.WithCreateSmtpProviderMailer("{$example}"),
        messaging.WithCreateSmtpProviderFromName("{$example}"),
        messaging.WithCreateSmtpProviderFromEmail("{$example}"),
        messaging.WithCreateSmtpProviderReplyToName("{$example}"),
        messaging.WithCreateSmtpProviderReplyToEmail("{$example}"),
        messaging.WithCreateSmtpProviderEnabled(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
