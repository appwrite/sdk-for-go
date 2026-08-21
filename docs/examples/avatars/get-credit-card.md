```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/avatars"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := avatars.New(client)

	response, err := service.GetCreditCard(
		"amex",
		service.WithGetCreditCardWidth(0),
		service.WithGetCreditCardHeight(0),
		service.WithGetCreditCardQuality(-1),
	)
	fmt.Println(response, err)
}
```
