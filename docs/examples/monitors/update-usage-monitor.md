```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/monitors"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := monitors.New(client)

	response, err := service.UpdateUsageMonitor(
		"<MONITOR_ID>",
		service.WithUpdateUsageMonitorName("<NAME>"),
		service.WithUpdateUsageMonitorRules([]string{"{\"method\":\"greaterThan\", \"attribute\":\"network.outbound\", \"values\":[5000000000]}"}),
		service.WithUpdateUsageMonitorWindow(60),
		service.WithUpdateUsageMonitorCooldown(300),
		service.WithUpdateUsageMonitorEnabled(false),
		service.WithUpdateUsageMonitorNotify([]string{"owner"}),
		service.WithUpdateUsageMonitorSeverity("info"),
		service.WithUpdateUsageMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
