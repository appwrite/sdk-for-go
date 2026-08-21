```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/postgresql"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := postgresql.New(client)

	response, err := service.CreateBackupPolicy(
		"<DATABASE_ID>",
		"<POLICY_ID>",
		"<NAME>",
		"",
		1,
		service.WithCreateBackupPolicyType("full"),
		service.WithCreateBackupPolicyEnabled(false),
	)
	fmt.Println(response, err)
}
```
