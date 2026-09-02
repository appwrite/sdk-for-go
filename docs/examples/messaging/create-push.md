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

	response, err := service.CreatePush(
		"<MESSAGE_ID>",
		service.WithCreatePushTitle("<TITLE>"),
		service.WithCreatePushBody("<BODY>"),
		service.WithCreatePushTopics([]string{"example"}),
		service.WithCreatePushUsers([]string{"example"}),
		service.WithCreatePushTargets([]string{"example"}),
		service.WithCreatePushData([]interface{}{}),
		service.WithCreatePushAction("<ACTION>"),
		service.WithCreatePushImage("<ID1:ID2>"),
		service.WithCreatePushIcon("<ICON>"),
		service.WithCreatePushSound("<SOUND>"),
		service.WithCreatePushColor("<COLOR>"),
		service.WithCreatePushTag("<TAG>"),
		service.WithCreatePushBadge(1),
		service.WithCreatePushDraft(false),
		service.WithCreatePushScheduledAt("2020-10-15T06:38:00.000+00:00"),
		service.WithCreatePushContentAvailable(false),
		service.WithCreatePushCritical(false),
		service.WithCreatePushPriority("normal"),
	)
	fmt.Println(response, err)
}
```
