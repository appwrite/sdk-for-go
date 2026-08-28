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

	response, err := service.CreateScryptUser(
		"<USER_ID>",
		"email@example.com",
		"password",
		"<PASSWORD_SALT>",
		8,
		65536,
		1,
		64,
		service.WithCreateScryptUserName("<NAME>"),
	)
	fmt.Println(response, err)
}
```
