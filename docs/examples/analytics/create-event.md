```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/analytics"
	"github.com/appwrite/sdk-for-go/v7/appwrite"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := analytics.New(client)

	response, err := service.CreateEvent(
		"<PROPERTY_ID>",
		"<NAME>",
		"https://example.com",
		service.WithCreateEventDomain("<DOMAIN>"),
		service.WithCreateEventReferrer("<REFERRER>"),
		service.WithCreateEventScreenWidth(0),
		service.WithCreateEventSessionHash("<SESSION_HASH>"),
		service.WithCreateEventScrollDepth(0),
		service.WithCreateEventEngagementTime(0),
		service.WithCreateEventProps([]string{"example"}),
		service.WithCreateEventUserId("<USER_ID>"),
		service.WithCreateEventIp("<IP>"),
		service.WithCreateEventUserAgent("<USER_AGENT>"),
	)
	fmt.Println(response, err)
}
```
