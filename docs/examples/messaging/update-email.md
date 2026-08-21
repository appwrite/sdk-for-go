```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/messaging"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := messaging.New(client)

	response, err := service.UpdateEmail(
		"<MESSAGE_ID>",
		service.WithUpdateEmailTopics([]string{}),
		service.WithUpdateEmailUsers([]string{}),
		service.WithUpdateEmailTargets([]string{}),
		service.WithUpdateEmailSubject("<SUBJECT>"),
		service.WithUpdateEmailContent("<CONTENT>"),
		service.WithUpdateEmailDraft(false),
		service.WithUpdateEmailHtml(false),
		service.WithUpdateEmailCc([]string{}),
		service.WithUpdateEmailBcc([]string{}),
		service.WithUpdateEmailScheduledAt("2020-10-15T06:38:00.000+00:00"),
		service.WithUpdateEmailAttachments([]string{}),
	)
	fmt.Println(response, err)
}
```
