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

	response, err := service.UpdateTcpMonitor(
		"<MONITOR_ID>",
		service.WithUpdateTcpMonitorName("<NAME>"),
		service.WithUpdateTcpMonitorHost("example.com"),
		service.WithUpdateTcpMonitorPort(443),
		service.WithUpdateTcpMonitorRules([]string{"{\"method\":\"greaterThan\", \"attribute\":\"latency\", \"values\":[2000]}"}),
		service.WithUpdateTcpMonitorInterval(1),
		service.WithUpdateTcpMonitorTimeout(1),
		service.WithUpdateTcpMonitorCooldown(300),
		service.WithUpdateTcpMonitorEnabled(false),
		service.WithUpdateTcpMonitorNotify([]string{"example"}),
		service.WithUpdateTcpMonitorSeverity("info"),
		service.WithUpdateTcpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
