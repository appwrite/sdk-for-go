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

	response, err := service.CreateDnsMonitor(
		"<MONITOR_ID>",
		"<NAME>",
		"example.com",
		service.WithCreateDnsMonitorRecord("A"),
		service.WithCreateDnsMonitorResolver("<RESOLVER>"),
		service.WithCreateDnsMonitorRules([]string{"{\"method\":\"greaterThan\", \"attribute\":\"latency\", \"values\":[2000]}"}),
		service.WithCreateDnsMonitorInterval(1),
		service.WithCreateDnsMonitorTimeout(1),
		service.WithCreateDnsMonitorCooldown(300),
		service.WithCreateDnsMonitorEnabled(false),
		service.WithCreateDnsMonitorNotify([]string{"example"}),
		service.WithCreateDnsMonitorSeverity("info"),
		service.WithCreateDnsMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
