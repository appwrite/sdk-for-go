```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/account"
	"github.com/appwrite/sdk-for-go/v7/client"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := account.New(client)

	response, err := service.UpdateRecovery(
		"<USER_ID>",
		"<SECRET>",
		"password",
	)
	fmt.Println(response, err)
}
```
