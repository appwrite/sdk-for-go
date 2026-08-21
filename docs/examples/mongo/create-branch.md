```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/mongo"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := mongo.New(client)

	response, err := service.CreateBranch(
		"<DATABASE_ID>",
		service.WithCreateBranchBranchId("<BRANCH_ID>"),
		service.WithCreateBranchTtl(300),
	)
	fmt.Println(response, err)
}
```
