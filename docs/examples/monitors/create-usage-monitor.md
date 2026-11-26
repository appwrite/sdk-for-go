```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/monitors"
	"github.com/appwrite/sdk-for-go/v7/query"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := monitors.New(client)

	response, err := service.CreateUsageMonitor(
		"<MONITOR_ID>",
		"<NAME>",
		[]string{
			query.GreaterThan("network.outbound", 5000000000),
		},
		service.WithCreateUsageMonitorWindow(3600),
		service.WithCreateUsageMonitorCooldown(1800),
		service.WithCreateUsageMonitorEnabled(true),
		service.WithCreateUsageMonitorNotify([]string{"owner"}),
		service.WithCreateUsageMonitorSeverity("info"),
		service.WithCreateUsageMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
