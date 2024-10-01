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
    response, error := messaging.UpdatePush(
        "{$example}",
        messaging.WithUpdatePushTopics([]interface{}{}),
        messaging.WithUpdatePushUsers([]interface{}{}),
        messaging.WithUpdatePushTargets([]interface{}{}),
        messaging.WithUpdatePushTitle("{$example}"),
        messaging.WithUpdatePushBody("{$example}"),
        messaging.WithUpdatePushData(map[string]interface{}{}),
        messaging.WithUpdatePushAction("{$example}"),
        messaging.WithUpdatePushImage("{$example}"),
        messaging.WithUpdatePushIcon("{$example}"),
        messaging.WithUpdatePushSound("{$example}"),
        messaging.WithUpdatePushColor("{$example}"),
        messaging.WithUpdatePushTag("{$example}"),
        messaging.WithUpdatePushBadge(0),
        messaging.WithUpdatePushDraft(false),
        messaging.WithUpdatePushScheduledAt(""),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
