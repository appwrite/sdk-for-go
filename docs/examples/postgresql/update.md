```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/postgresql"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := postgresql.New(client)

	response, err := service.Update(
		"<DATABASE_ID>",
		service.WithUpdateName("<NAME>"),
		service.WithUpdateStatus("ready"),
		service.WithUpdateSpecification("<SPECIFICATION>"),
		service.WithUpdateReplicas(0),
		service.WithUpdateSyncMode("async"),
		service.WithUpdateNetworkIdleTimeoutSeconds(60),
		service.WithUpdateNetworkIPAllowlist([]string{"example"}),
		service.WithUpdateIdleTimeoutMinutes(5),
		service.WithUpdatePitr(false),
		service.WithUpdatePitrRetentionDays(1),
		service.WithUpdateStorageAutoscaling(false),
		service.WithUpdateStorageAutoscalingThresholdPercent(50),
		service.WithUpdateStorageAutoscalingMaxGb(0),
		service.WithUpdateMetricsTraceSampleRate(0),
		service.WithUpdateMetricsSlowQueryLogThresholdMs(0),
		service.WithUpdateSqlApiEnabled(false),
		service.WithUpdateSqlApiAllowedStatements([]string{"example"}),
		service.WithUpdateSqlApiMaxRows(1),
		service.WithUpdateSqlApiMaxBytes(1024),
		service.WithUpdateSqlApiTimeoutSeconds(1),
	)
	fmt.Println(response, err)
}
```
