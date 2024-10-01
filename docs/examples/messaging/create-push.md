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
    response, error := messaging.CreatePush(
        "{$example}",
        "{$example}",
        "{$example}",
        messaging.WithCreatePushTopics([]interface{}{}),
        messaging.WithCreatePushUsers([]interface{}{}),
        messaging.WithCreatePushTargets([]interface{}{}),
        messaging.WithCreatePushData(map[string]interface{}{}),
        messaging.WithCreatePushAction("{$example}"),
        messaging.WithCreatePushImage("{$example}"),
        messaging.WithCreatePushIcon("{$example}"),
        messaging.WithCreatePushSound("{$example}"),
        messaging.WithCreatePushColor("{$example}"),
        messaging.WithCreatePushTag("{$example}"),
        messaging.WithCreatePushBadge("{$example}"),
        messaging.WithCreatePushDraft(false),
        messaging.WithCreatePushScheduledAt(""),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
