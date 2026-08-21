```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/tokens"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithKey("<YOUR_API_KEY>"),
	)

	service := tokens.New(client)

	response, err := service.CreateFileToken(
		"<BUCKET_ID>",
		"<FILE_ID>",
		tokens.WithCreateFileTokenExpire("2020-10-15T06:38:00.000+00:00"),
	)
	fmt.Println(response, err)
}
```
