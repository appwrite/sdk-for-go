```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v4/client"
    "github.com/appwrite/sdk-for-go/v4/project"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := project.New(client)

response, error := service.UpdateEmailTemplate(
    "verification",
    project.WithUpdateEmailTemplateLocale("af"),
    project.WithUpdateEmailTemplateSubject("<SUBJECT>"),
    project.WithUpdateEmailTemplateMessage("<MESSAGE>"),
    project.WithUpdateEmailTemplateSenderName("<SENDER_NAME>"),
    project.WithUpdateEmailTemplateSenderEmail("email@example.com"),
    project.WithUpdateEmailTemplateReplyToEmail("email@example.com"),
    project.WithUpdateEmailTemplateReplyToName("<REPLY_TO_NAME>"),
)
```
