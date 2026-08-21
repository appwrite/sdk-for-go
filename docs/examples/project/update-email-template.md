```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/project"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := project.New(client)

	response, err := service.UpdateEmailTemplate(
		"verification",
		service.WithUpdateEmailTemplateLocale("af"),
		service.WithUpdateEmailTemplateSubject("<SUBJECT>"),
		service.WithUpdateEmailTemplateMessage("<MESSAGE>"),
		service.WithUpdateEmailTemplateSenderName("<SENDER_NAME>"),
		service.WithUpdateEmailTemplateSenderEmail("email@example.com"),
		service.WithUpdateEmailTemplateReplyToEmail("email@example.com"),
		service.WithUpdateEmailTemplateReplyToName("<REPLY_TO_NAME>"),
	)
	fmt.Println(response, err)
}
```
