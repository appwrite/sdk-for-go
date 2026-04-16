```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v3/client"
    "github.com/appwrite/sdk-for-go/v3/messaging"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := messaging.New(client)

response, error := service.UpdateEmail(
    "<MESSAGE_ID>",
    messaging.WithUpdateEmailTopics([]interface{}{}),
    messaging.WithUpdateEmailUsers([]interface{}{}),
    messaging.WithUpdateEmailTargets([]interface{}{}),
    messaging.WithUpdateEmailSubject("<SUBJECT>"),
    messaging.WithUpdateEmailContent("<CONTENT>"),
    messaging.WithUpdateEmailDraft(false),
    messaging.WithUpdateEmailHtml(false),
    messaging.WithUpdateEmailCc([]interface{}{}),
    messaging.WithUpdateEmailBcc([]interface{}{}),
    messaging.WithUpdateEmailScheduledAt("2020-10-15T06:38:00.000+00:00"),
    messaging.WithUpdateEmailAttachments([]interface{}{}),
)
```
