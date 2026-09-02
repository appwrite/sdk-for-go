```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/teams"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := teams.New(client)

	response, err := service.CreateMembership(
		"<TEAM_ID>",
		[]string{"example"},
		service.WithCreateMembershipEmail("email@example.com"),
		service.WithCreateMembershipUserId("<USER_ID>"),
		service.WithCreateMembershipPhone("+12065550100"),
		service.WithCreateMembershipUrl("https://example.com"),
		service.WithCreateMembershipName("<NAME>"),
	)
	fmt.Println(response, err)
}
```
