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
    response, error := messaging.UpdateMsg91Provider(
        "{$example}",
        messaging.WithUpdateMsg91ProviderName("{$example}"),
        messaging.WithUpdateMsg91ProviderEnabled(false),
        messaging.WithUpdateMsg91ProviderTemplateId("{$example}"),
        messaging.WithUpdateMsg91ProviderSenderId("{$example}"),
        messaging.WithUpdateMsg91ProviderAuthKey("{$example}"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
