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
    response, error := messaging.CreateMsg91Provider(
        "{$example}",
        "{$example}",
        messaging.WithCreateMsg91ProviderTemplateId("{$example}"),
        messaging.WithCreateMsg91ProviderSenderId("{$example}"),
        messaging.WithCreateMsg91ProviderAuthKey("{$example}"),
        messaging.WithCreateMsg91ProviderEnabled(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
