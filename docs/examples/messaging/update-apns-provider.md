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
    response, error := messaging.UpdateApnsProvider(
        "{$example}",
        messaging.WithUpdateApnsProviderName("{$example}"),
        messaging.WithUpdateApnsProviderEnabled(false),
        messaging.WithUpdateApnsProviderAuthKey("{$example}"),
        messaging.WithUpdateApnsProviderAuthKeyId("{$example}"),
        messaging.WithUpdateApnsProviderTeamId("{$example}"),
        messaging.WithUpdateApnsProviderBundleId("{$example}"),
        messaging.WithUpdateApnsProviderSandbox(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
