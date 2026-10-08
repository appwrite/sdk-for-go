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

	response, err := service.CreateChallengeRule(
		"<RULE_ID>",
		"api",
		"<NAME>",
		service.WithCreateChallengeRuleResourceId("<RESOURCE_ID>"),
		service.WithCreateChallengeRuleDescription("<DESCRIPTION>"),
		service.WithCreateChallengeRuleChallengeType("compute"),
		service.WithCreateChallengeRulePriority(-100000),
		service.WithCreateChallengeRuleEnabled(false),
		service.WithCreateChallengeRuleConditions([]string{"example"}),
		service.WithCreateChallengeRuleDifficulty(1),
		service.WithCreateChallengeRuleTtl(900),
	)
	fmt.Println(response, err)
}
```
