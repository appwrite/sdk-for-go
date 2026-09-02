```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/mongo"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := mongo.New(client)

	response, err := service.Create(
		"<DATABASE_ID>",
		"<NAME>",
		service.WithCreateVersion("17"),
		service.WithCreateSpecification("<SPECIFICATION>"),
		service.WithCreateReplicas(0),
		service.WithCreateSyncMode("async"),
		service.WithCreateNetworkIdleTimeoutSeconds(60),
		service.WithCreateNetworkIPAllowlist([]string{"example"}),
		service.WithCreateIdleTimeoutMinutes(5),
		service.WithCreatePitr(false),
		service.WithCreatePitrRetentionDays(1),
		service.WithCreateStorageAutoscaling(false),
		service.WithCreateStorageAutoscalingThresholdPercent(50),
		service.WithCreateStorageAutoscalingMaxGb(0),
	)
	fmt.Println(response, err)
}
```
