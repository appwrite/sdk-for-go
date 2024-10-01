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
    response, error := messaging.CreateEmail(
        "{$example}",
        "{$example}",
        "{$example}",
        messaging.WithCreateEmailTopics([]interface{}{}),
        messaging.WithCreateEmailUsers([]interface{}{}),
        messaging.WithCreateEmailTargets([]interface{}{}),
        messaging.WithCreateEmailCc([]interface{}{}),
        messaging.WithCreateEmailBcc([]interface{}{}),
        messaging.WithCreateEmailAttachments([]interface{}{}),
        messaging.WithCreateEmailDraft(false),
        messaging.WithCreateEmailHtml(false),
        messaging.WithCreateEmailScheduledAt(""),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
