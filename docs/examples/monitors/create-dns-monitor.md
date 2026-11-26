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

	response, err := service.CreateDnsMonitor(
		"<MONITOR_ID>",
		"<NAME>",
		"example.com",
		service.WithCreateDnsMonitorRecord("A"),
		service.WithCreateDnsMonitorResolver("1.1.1.1"),
		service.WithCreateDnsMonitorRules([]string{
			query.Or([]string{
				query.Equal("answers", 0),
				query.GreaterThan("latency", 1000),
			}),
		}),
		service.WithCreateDnsMonitorInterval(60),
		service.WithCreateDnsMonitorTimeout(10),
		service.WithCreateDnsMonitorCooldown(300),
		service.WithCreateDnsMonitorEnabled(true),
		service.WithCreateDnsMonitorNotify([]string{"owner"}),
		service.WithCreateDnsMonitorSeverity("info"),
		service.WithCreateDnsMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
