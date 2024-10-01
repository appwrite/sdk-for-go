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
    response, error := messaging.UpdateEmail(
        "{$example}",
        messaging.WithUpdateEmailTopics([]interface{}{}),
        messaging.WithUpdateEmailUsers([]interface{}{}),
        messaging.WithUpdateEmailTargets([]interface{}{}),
        messaging.WithUpdateEmailSubject("{$example}"),
        messaging.WithUpdateEmailContent("{$example}"),
        messaging.WithUpdateEmailDraft(false),
        messaging.WithUpdateEmailHtml(false),
        messaging.WithUpdateEmailCc([]interface{}{}),
        messaging.WithUpdateEmailBcc([]interface{}{}),
        messaging.WithUpdateEmailScheduledAt(""),
        messaging.WithUpdateEmailAttachments([]interface{}{}),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
