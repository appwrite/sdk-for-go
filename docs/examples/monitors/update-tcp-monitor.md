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

	response, err := service.UpdateTcpMonitor(
		"<MONITOR_ID>",
		service.WithUpdateTcpMonitorName("<NAME>"),
		service.WithUpdateTcpMonitorHost("db.example.com"),
		service.WithUpdateTcpMonitorPort(5432),
		service.WithUpdateTcpMonitorRules([]string{
			query.GreaterThan("latency", 1000),
		}),
		service.WithUpdateTcpMonitorInterval(60),
		service.WithUpdateTcpMonitorTimeout(10),
		service.WithUpdateTcpMonitorCooldown(300),
		service.WithUpdateTcpMonitorEnabled(true),
		service.WithUpdateTcpMonitorNotify([]string{"owner"}),
		service.WithUpdateTcpMonitorSeverity("info"),
		service.WithUpdateTcpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
