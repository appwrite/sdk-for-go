```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/functions"
)

func main() {
	client := client.New(
		client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		client.WithProject("<YOUR_PROJECT_ID>"),
		client.WithSession(""),
	)

	service := functions.New(client)

	response, err := service.GetExecution(
		"<FUNCTION_ID>",
		"<EXECUTION_ID>",
	)
	fmt.Println(response, err)
}
```
