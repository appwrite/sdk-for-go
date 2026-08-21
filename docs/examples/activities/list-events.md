```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/activities"
	"github.com/appwrite/sdk-for-go/v7/client"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := activities.New(client)

	response, err := service.ListEvents(
		activities.WithListEventsQueries([]string{}),
	)
	fmt.Println(response, err)
}
```
