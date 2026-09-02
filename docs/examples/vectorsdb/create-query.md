```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/vectorsdb"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := vectorsdb.New(client)

	response, err := service.CreateQuery(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		service.WithCreateQueryQueries([]string{"example"}),
		service.WithCreateQueryTransactionId("<TRANSACTION_ID>"),
		service.WithCreateQueryTotal(false),
		service.WithCreateQueryTtl(0),
	)
	fmt.Println(response, err)
}
```
