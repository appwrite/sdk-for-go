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

	response, err := service.CreateTcpMonitor(
		"<MONITOR_ID>",
		"<NAME>",
		"db.example.com",
		5432,
		service.WithCreateTcpMonitorRules([]string{
			query.GreaterThan("latency", 1000),
		}),
		service.WithCreateTcpMonitorInterval(60),
		service.WithCreateTcpMonitorTimeout(10),
		service.WithCreateTcpMonitorCooldown(300),
		service.WithCreateTcpMonitorEnabled(true),
		service.WithCreateTcpMonitorNotify([]string{"owner"}),
		service.WithCreateTcpMonitorSeverity("info"),
		service.WithCreateTcpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
