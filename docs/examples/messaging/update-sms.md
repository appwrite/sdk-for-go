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

	response, err := service.UpdateSMS(
		"<MESSAGE_ID>",
		service.WithUpdateSMSTopics([]string{"example"}),
		service.WithUpdateSMSUsers([]string{"example"}),
		service.WithUpdateSMSTargets([]string{"example"}),
		service.WithUpdateSMSContent("<CONTENT>"),
		service.WithUpdateSMSDraft(false),
		service.WithUpdateSMSScheduledAt("2020-10-15T06:38:00.000+00:00"),
	)
	fmt.Println(response, err)
}
```
