```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/functions"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := functions.New(client)

	response, err := service.CreateExecution(
		"<FUNCTION_ID>",
		service.WithCreateExecutionBody("<BODY>"),
		service.WithCreateExecutionAsync(false),
		service.WithCreateExecutionPath("<PATH>"),
		service.WithCreateExecutionMethod("GET"),
		service.WithCreateExecutionHeaders([]interface{}{}),
		service.WithCreateExecutionScheduledAt("<SCHEDULED_AT>"),
	)
	fmt.Println(response, err)
}
```
