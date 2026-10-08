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

	response, err := service.UpdateRedirectRule(
		"<RULE_ID>",
		service.WithUpdateRedirectRuleResourceType("api"),
		service.WithUpdateRedirectRuleResourceId("<RESOURCE_ID>"),
		service.WithUpdateRedirectRuleName("<NAME>"),
		service.WithUpdateRedirectRuleDescription("<DESCRIPTION>"),
		service.WithUpdateRedirectRuleLocation("<LOCATION>"),
		service.WithUpdateRedirectRuleStatusCode(300),
		service.WithUpdateRedirectRulePriority(-100000),
		service.WithUpdateRedirectRuleEnabled(false),
		service.WithUpdateRedirectRuleConditions([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
