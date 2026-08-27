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

	response, err := service.UpdateSMTP(
		service.WithUpdateSMTPHost("example.com"),
		service.WithUpdateSMTPPort(587),
		service.WithUpdateSMTPUsername("<USERNAME>"),
		service.WithUpdateSMTPPassword("password"),
		service.WithUpdateSMTPSenderEmail("email@example.com"),
		service.WithUpdateSMTPSenderName("<SENDER_NAME>"),
		service.WithUpdateSMTPReplyToEmail("email@example.com"),
		service.WithUpdateSMTPReplyToName("<REPLY_TO_NAME>"),
		service.WithUpdateSMTPSecure("tls"),
		service.WithUpdateSMTPEnabled(false),
	)
	fmt.Println(response, err)
}
```
