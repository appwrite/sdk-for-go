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

	response, err := service.UpdateHttpMonitor(
		"<MONITOR_ID>",
		service.WithUpdateHttpMonitorName("<NAME>"),
		service.WithUpdateHttpMonitorUrl("https://example.com/health"),
		service.WithUpdateHttpMonitorMethod("GET"),
		service.WithUpdateHttpMonitorHeaders(map[string]interface{}{"Accept": "application\/json"}),
		service.WithUpdateHttpMonitorBody("<BODY>"),
		service.WithUpdateHttpMonitorRules([]string{"{\"method\":\"greaterThan\", \"attribute\":\"latency\", \"values\":[2000]}"}),
		service.WithUpdateHttpMonitorInterval(1),
		service.WithUpdateHttpMonitorTimeout(1),
		service.WithUpdateHttpMonitorCooldown(300),
		service.WithUpdateHttpMonitorEnabled(false),
		service.WithUpdateHttpMonitorNotify([]string{"example"}),
		service.WithUpdateHttpMonitorSeverity("info"),
		service.WithUpdateHttpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
