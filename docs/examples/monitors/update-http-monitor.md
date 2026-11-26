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

	response, err := service.UpdateHttpMonitor(
		"<MONITOR_ID>",
		service.WithUpdateHttpMonitorName("<NAME>"),
		service.WithUpdateHttpMonitorUrl("https://api.example.com/health"),
		service.WithUpdateHttpMonitorMethod("GET"),
		service.WithUpdateHttpMonitorHeaders(map[string]interface{}{"Authorization": "Bearer <TOKEN>"}),
		service.WithUpdateHttpMonitorBody("<BODY>"),
		service.WithUpdateHttpMonitorRules([]string{
			query.Or([]string{
				query.NotBetween("statusCode", 200, 299),
				query.GreaterThan("latency", 2000),
			}),
		}),
		service.WithUpdateHttpMonitorInterval(60),
		service.WithUpdateHttpMonitorTimeout(10),
		service.WithUpdateHttpMonitorCooldown(300),
		service.WithUpdateHttpMonitorEnabled(true),
		service.WithUpdateHttpMonitorNotify([]string{"owner"}),
		service.WithUpdateHttpMonitorSeverity("info"),
		service.WithUpdateHttpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
