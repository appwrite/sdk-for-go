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

	response, err := service.UpdatePush(
		"<MESSAGE_ID>",
		service.WithUpdatePushTopics([]string{"example"}),
		service.WithUpdatePushUsers([]string{"example"}),
		service.WithUpdatePushTargets([]string{"example"}),
		service.WithUpdatePushTitle("<TITLE>"),
		service.WithUpdatePushBody("<BODY>"),
		service.WithUpdatePushData([]interface{}{}),
		service.WithUpdatePushAction("<ACTION>"),
		service.WithUpdatePushImage("<ID1:ID2>"),
		service.WithUpdatePushIcon("<ICON>"),
		service.WithUpdatePushSound("<SOUND>"),
		service.WithUpdatePushColor("<COLOR>"),
		service.WithUpdatePushTag("<TAG>"),
		service.WithUpdatePushBadge(1),
		service.WithUpdatePushDraft(false),
		service.WithUpdatePushScheduledAt("2020-10-15T06:38:00.000+00:00"),
		service.WithUpdatePushContentAvailable(false),
		service.WithUpdatePushCritical(false),
		service.WithUpdatePushPriority("normal"),
	)
	fmt.Println(response, err)
}
```
