```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v2/client"
    "github.com/appwrite/sdk-for-go/v2/messaging"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := messaging.New(client)

response, error := service.CreateEmail(
    "<MESSAGE_ID>",
    "<SUBJECT>",
    "<CONTENT>",
    messaging.WithCreateEmailTopics([]interface{}{}),
    messaging.WithCreateEmailUsers([]interface{}{}),
    messaging.WithCreateEmailTargets([]interface{}{}),
    messaging.WithCreateEmailCc([]interface{}{}),
    messaging.WithCreateEmailBcc([]interface{}{}),
    messaging.WithCreateEmailAttachments([]interface{}{}),
    messaging.WithCreateEmailDraft(false),
    messaging.WithCreateEmailHtml(false),
    messaging.WithCreateEmailScheduledAt("2020-10-15T06:38:00.000+00:00"),
)
```
