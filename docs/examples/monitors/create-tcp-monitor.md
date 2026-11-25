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

	response, err := service.CreateTcpMonitor(
		"<MONITOR_ID>",
		"<NAME>",
		"example.com",
		443,
		service.WithCreateTcpMonitorRules([]string{"{\"method\":\"greaterThan\", \"attribute\":\"latency\", \"values\":[2000]}"}),
		service.WithCreateTcpMonitorInterval(1),
		service.WithCreateTcpMonitorTimeout(1),
		service.WithCreateTcpMonitorCooldown(300),
		service.WithCreateTcpMonitorEnabled(false),
		service.WithCreateTcpMonitorNotify([]string{"example"}),
		service.WithCreateTcpMonitorSeverity("info"),
		service.WithCreateTcpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
