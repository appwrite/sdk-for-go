```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/account"
	"github.com/appwrite/sdk-for-go/v7/appwrite"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := account.New(client)

	response, err := service.CreatePasskey(
		service.WithCreatePasskeyPasskeyId("<PASSKEY_ID>"),
		service.WithCreatePasskeyName("<NAME>"),
	)
	fmt.Println(response, err)
}
```
