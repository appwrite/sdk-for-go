```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/waf"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := waf.New(client)

	response, err := service.UpdateDenyRule(
		"<RULE_ID>",
		service.WithUpdateDenyRuleResourceType("api"),
		service.WithUpdateDenyRuleResourceId("<RESOURCE_ID>"),
		service.WithUpdateDenyRuleName("<NAME>"),
		service.WithUpdateDenyRuleDescription("<DESCRIPTION>"),
		service.WithUpdateDenyRulePriority(-100000),
		service.WithUpdateDenyRuleEnabled(false),
		service.WithUpdateDenyRuleConditions([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
