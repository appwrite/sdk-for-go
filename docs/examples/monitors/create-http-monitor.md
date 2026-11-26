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

	response, err := service.CreateHttpMonitor(
		"<MONITOR_ID>",
		"<NAME>",
		"https://api.example.com/health",
		service.WithCreateHttpMonitorMethod("GET"),
		service.WithCreateHttpMonitorHeaders(map[string]interface{}{"Authorization": "Bearer <TOKEN>"}),
		service.WithCreateHttpMonitorBody("<BODY>"),
		service.WithCreateHttpMonitorRules([]string{
			query.Or([]string{
				query.NotBetween("statusCode", 200, 299),
				query.GreaterThan("latency", 2000),
			}),
		}),
		service.WithCreateHttpMonitorInterval(60),
		service.WithCreateHttpMonitorTimeout(10),
		service.WithCreateHttpMonitorCooldown(300),
		service.WithCreateHttpMonitorEnabled(true),
		service.WithCreateHttpMonitorNotify([]string{"owner"}),
		service.WithCreateHttpMonitorSeverity("info"),
		service.WithCreateHttpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
