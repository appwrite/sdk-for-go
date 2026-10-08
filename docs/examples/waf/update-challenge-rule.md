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

	response, err := service.UpdateChallengeRule(
		"<RULE_ID>",
		service.WithUpdateChallengeRuleResourceType("api"),
		service.WithUpdateChallengeRuleResourceId("<RESOURCE_ID>"),
		service.WithUpdateChallengeRuleName("<NAME>"),
		service.WithUpdateChallengeRuleDescription("<DESCRIPTION>"),
		service.WithUpdateChallengeRuleChallengeType("compute"),
		service.WithUpdateChallengeRulePriority(-100000),
		service.WithUpdateChallengeRuleEnabled(false),
		service.WithUpdateChallengeRuleConditions([]string{"example"}),
		service.WithUpdateChallengeRuleDifficulty(1),
		service.WithUpdateChallengeRuleTtl(900),
	)
	fmt.Println(response, err)
}
```
