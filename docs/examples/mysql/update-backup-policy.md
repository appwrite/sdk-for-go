```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/mysql"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := mysql.New(client)

	response, err := service.UpdateBackupPolicy(
		"<DATABASE_ID>",
		"<POLICY_ID>",
		service.WithUpdateBackupPolicyName("<NAME>"),
		service.WithUpdateBackupPolicySchedule(""),
		service.WithUpdateBackupPolicyRetention(1),
		service.WithUpdateBackupPolicyEnabled(false),
	)
	fmt.Println(response, err)
}
```
