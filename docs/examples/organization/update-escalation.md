```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/organization"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := organization.New(client)

	response, err := service.UpdateEscalation(
		"<ESCALATION_ID>",
		service.WithUpdateEscalationName("<NAME>"),
		service.WithUpdateEscalationSteps([]interface{}{map[string]interface{}{"delay": 0, "roles": []interface{}{"owner"}}, map[string]interface{}{"delay": 900, "roles": []interface{}{"developer"}, "userIds": []interface{}{"<USER_ID>"}}}),
		service.WithUpdateEscalationRepeat(1),
	)
	fmt.Println(response, err)
}
```
