```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/domains"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := domains.New(client)

	response, err := service.CreateRecordCAA(
		"<DOMAIN_ID>",
		"",
		"",
		1,
		service.WithCreateRecordCAAComment("<COMMENT>"),
	)
	fmt.Println(response, err)
}
```
