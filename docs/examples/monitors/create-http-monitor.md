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

	response, err := service.CreateHttpMonitor(
		"<MONITOR_ID>",
		"<NAME>",
		"https://example.com/health",
		service.WithCreateHttpMonitorMethod("GET"),
		service.WithCreateHttpMonitorHeaders(map[string]interface{}{"Accept": "application\/json"}),
		service.WithCreateHttpMonitorBody("<BODY>"),
		service.WithCreateHttpMonitorRules([]string{"{\"method\":\"greaterThan\", \"attribute\":\"latency\", \"values\":[2000]}"}),
		service.WithCreateHttpMonitorInterval(1),
		service.WithCreateHttpMonitorTimeout(1),
		service.WithCreateHttpMonitorCooldown(300),
		service.WithCreateHttpMonitorEnabled(false),
		service.WithCreateHttpMonitorNotify([]string{"example"}),
		service.WithCreateHttpMonitorSeverity("info"),
		service.WithCreateHttpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
