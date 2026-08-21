```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/users"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := users.New(client)

	response, err := service.Create(
		"<USER_ID>",
		service.WithCreateEmail("email@example.com"),
		service.WithCreatePhone("+12065550100"),
		service.WithCreatePassword("password"),
		service.WithCreateName("<NAME>"),
	)
	fmt.Println(response, err)
}
```
