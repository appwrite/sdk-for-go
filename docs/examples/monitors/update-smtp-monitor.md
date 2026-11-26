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

	response, err := service.UpdateSmtpMonitor(
		"<MONITOR_ID>",
		service.WithUpdateSmtpMonitorName("<NAME>"),
		service.WithUpdateSmtpMonitorHost("smtp.example.com"),
		service.WithUpdateSmtpMonitorPort(587),
		service.WithUpdateSmtpMonitorTlsMode("none"),
		service.WithUpdateSmtpMonitorRules([]string{
			query.Or([]string{
				query.LessThan("certificateDays", 14),
				query.GreaterThan("latency", 2000),
			}),
		}),
		service.WithUpdateSmtpMonitorInterval(60),
		service.WithUpdateSmtpMonitorTimeout(10),
		service.WithUpdateSmtpMonitorCooldown(300),
		service.WithUpdateSmtpMonitorEnabled(true),
		service.WithUpdateSmtpMonitorNotify([]string{"owner"}),
		service.WithUpdateSmtpMonitorSeverity("info"),
		service.WithUpdateSmtpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
