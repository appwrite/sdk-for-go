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

	response, err := service.CreateRateLimitRule(
		"<RULE_ID>",
		"api",
		"<NAME>",
		1,
		1,
		service.WithCreateRateLimitRuleResourceId("<RESOURCE_ID>"),
		service.WithCreateRateLimitRuleDescription("<DESCRIPTION>"),
		service.WithCreateRateLimitRuleKey("ip"),
		service.WithCreateRateLimitRuleStrategy("fixedWindow"),
		service.WithCreateRateLimitRuleMaxBucketSize(1),
		service.WithCreateRateLimitRulePriority(-100000),
		service.WithCreateRateLimitRuleEnabled(false),
		service.WithCreateRateLimitRuleConditions([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
