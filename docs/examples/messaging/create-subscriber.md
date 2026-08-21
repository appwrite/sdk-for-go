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
		appwrite.WithJWT("<YOUR_JWT>"),
	)

	service := messaging.New(client)

	response, err := service.CreateSubscriber(
		"<TOPIC_ID>",
		"<SUBSCRIBER_ID>",
		"<TARGET_ID>",
	)
	fmt.Println(response, err)
}
```
