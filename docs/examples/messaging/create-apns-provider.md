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
    response, error := messaging.CreateApnsProvider(
        "{$example}",
        "{$example}",
        messaging.WithCreateApnsProviderAuthKey("{$example}"),
        messaging.WithCreateApnsProviderAuthKeyId("{$example}"),
        messaging.WithCreateApnsProviderTeamId("{$example}"),
        messaging.WithCreateApnsProviderBundleId("{$example}"),
        messaging.WithCreateApnsProviderSandbox(false),
        messaging.WithCreateApnsProviderEnabled(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
