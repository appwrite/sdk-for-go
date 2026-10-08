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

	response, err := service.UpdateBypassRule(
		"<RULE_ID>",
		service.WithUpdateBypassRuleResourceType("api"),
		service.WithUpdateBypassRuleResourceId("<RESOURCE_ID>"),
		service.WithUpdateBypassRuleName("<NAME>"),
		service.WithUpdateBypassRuleDescription("<DESCRIPTION>"),
		service.WithUpdateBypassRulePriority(-100000),
		service.WithUpdateBypassRuleEnabled(false),
		service.WithUpdateBypassRuleConditions([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
