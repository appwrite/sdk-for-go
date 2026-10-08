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

	response, err := service.UpdateRateLimitRule(
		"<RULE_ID>",
		service.WithUpdateRateLimitRuleResourceType("api"),
		service.WithUpdateRateLimitRuleResourceId("<RESOURCE_ID>"),
		service.WithUpdateRateLimitRuleName("<NAME>"),
		service.WithUpdateRateLimitRuleDescription("<DESCRIPTION>"),
		service.WithUpdateRateLimitRuleLimit(1),
		service.WithUpdateRateLimitRuleInterval(1),
		service.WithUpdateRateLimitRuleKey("ip"),
		service.WithUpdateRateLimitRuleMaxBucketSize(1),
		service.WithUpdateRateLimitRulePriority(-100000),
		service.WithUpdateRateLimitRuleEnabled(false),
		service.WithUpdateRateLimitRuleConditions([]string{"example"}),
	)
	fmt.Println(response, err)
}
```
