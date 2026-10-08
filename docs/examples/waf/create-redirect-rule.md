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

	response, err := service.CreateRedirectRule(
		"<RULE_ID>",
		"api",
		"<NAME>",
		"<LOCATION>",
		300,
		service.WithCreateRedirectRuleResourceId("<RESOURCE_ID>"),
		service.WithCreateRedirectRuleDescription("<DESCRIPTION>"),
		service.WithCreateRedirectRulePriority(-100000),
		service.WithCreateRedirectRuleEnabled(false),
		service.WithCreateRedirectRuleConditions([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
