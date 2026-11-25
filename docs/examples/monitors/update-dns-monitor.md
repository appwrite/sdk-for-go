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

	response, err := service.UpdateDnsMonitor(
		"<MONITOR_ID>",
		service.WithUpdateDnsMonitorName("<NAME>"),
		service.WithUpdateDnsMonitorDomain("example.com"),
		service.WithUpdateDnsMonitorRecord("A"),
		service.WithUpdateDnsMonitorResolver("<RESOLVER>"),
		service.WithUpdateDnsMonitorRules([]string{"{\"method\":\"greaterThan\", \"attribute\":\"latency\", \"values\":[2000]}"}),
		service.WithUpdateDnsMonitorInterval(1),
		service.WithUpdateDnsMonitorTimeout(1),
		service.WithUpdateDnsMonitorCooldown(300),
		service.WithUpdateDnsMonitorEnabled(false),
		service.WithUpdateDnsMonitorNotify([]string{"example"}),
		service.WithUpdateDnsMonitorSeverity("info"),
		service.WithUpdateDnsMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
