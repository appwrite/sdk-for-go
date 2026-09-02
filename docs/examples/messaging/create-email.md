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

	response, err := service.CreateEmail(
		"<MESSAGE_ID>",
		"<SUBJECT>",
		"<CONTENT>",
		service.WithCreateEmailTopics([]string{"example"}),
		service.WithCreateEmailUsers([]string{"example"}),
		service.WithCreateEmailTargets([]string{"example"}),
		service.WithCreateEmailCc([]string{"example"}),
		service.WithCreateEmailBcc([]string{"example"}),
		service.WithCreateEmailAttachments([]string{"example"}),
		service.WithCreateEmailDraft(false),
		service.WithCreateEmailHtml(false),
		service.WithCreateEmailScheduledAt("2020-10-15T06:38:00.000+00:00"),
	)
	fmt.Println(response, err)
}
```
