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
    response, error := messaging.UpdateSms(
        "{$example}",
        messaging.WithUpdateSmsTopics([]interface{}{}),
        messaging.WithUpdateSmsUsers([]interface{}{}),
        messaging.WithUpdateSmsTargets([]interface{}{}),
        messaging.WithUpdateSmsContent("{$example}"),
        messaging.WithUpdateSmsDraft(false),
        messaging.WithUpdateSmsScheduledAt(""),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
