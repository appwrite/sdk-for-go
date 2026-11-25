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

	response, err := service.UpdateSmtpMonitor(
		"<MONITOR_ID>",
		service.WithUpdateSmtpMonitorName("<NAME>"),
		service.WithUpdateSmtpMonitorHost("example.com"),
		service.WithUpdateSmtpMonitorPort(587),
		service.WithUpdateSmtpMonitorTlsMode("none"),
		service.WithUpdateSmtpMonitorRules([]string{"{\"method\":\"greaterThan\", \"attribute\":\"latency\", \"values\":[2000]}"}),
		service.WithUpdateSmtpMonitorInterval(1),
		service.WithUpdateSmtpMonitorTimeout(1),
		service.WithUpdateSmtpMonitorCooldown(300),
		service.WithUpdateSmtpMonitorEnabled(false),
		service.WithUpdateSmtpMonitorNotify([]string{"example"}),
		service.WithUpdateSmtpMonitorSeverity("info"),
		service.WithUpdateSmtpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
