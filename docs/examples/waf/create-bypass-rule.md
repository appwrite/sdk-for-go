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

	response, err := service.CreateBypassRule(
		"<RULE_ID>",
		"api",
		"<NAME>",
		service.WithCreateBypassRuleResourceId("<RESOURCE_ID>"),
		service.WithCreateBypassRuleDescription("<DESCRIPTION>"),
		service.WithCreateBypassRulePriority(-100000),
		service.WithCreateBypassRuleEnabled(false),
		service.WithCreateBypassRuleConditions([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
