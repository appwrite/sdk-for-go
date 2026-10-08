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

	response, err := service.CreateDenyRule(
		"<RULE_ID>",
		"api",
		"<NAME>",
		service.WithCreateDenyRuleResourceId("<RESOURCE_ID>"),
		service.WithCreateDenyRuleDescription("<DESCRIPTION>"),
		service.WithCreateDenyRulePriority(-100000),
		service.WithCreateDenyRuleEnabled(false),
		service.WithCreateDenyRuleConditions([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
