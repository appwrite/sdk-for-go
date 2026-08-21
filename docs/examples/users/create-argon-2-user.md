```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/users"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := users.New(client)

	response, err := service.CreateArgon2User(
		"<USER_ID>",
		"email@example.com",
		"password",
		users.WithCreateArgon2UserName("<NAME>"),
	)
	fmt.Println(response, err)
}
```
