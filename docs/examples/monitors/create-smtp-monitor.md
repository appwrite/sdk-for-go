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

	response, err := service.CreateSmtpMonitor(
		"<MONITOR_ID>",
		"<NAME>",
		"smtp.example.com",
		587,
		service.WithCreateSmtpMonitorTlsMode("none"),
		service.WithCreateSmtpMonitorRules([]string{
			query.Or([]string{
				query.LessThan("certificateDays", 14),
				query.GreaterThan("latency", 2000),
			}),
		}),
		service.WithCreateSmtpMonitorInterval(60),
		service.WithCreateSmtpMonitorTimeout(10),
		service.WithCreateSmtpMonitorCooldown(300),
		service.WithCreateSmtpMonitorEnabled(true),
		service.WithCreateSmtpMonitorNotify([]string{"owner"}),
		service.WithCreateSmtpMonitorSeverity("info"),
		service.WithCreateSmtpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
