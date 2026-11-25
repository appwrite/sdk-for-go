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

	response, err := service.CreateSmtpMonitor(
		"<MONITOR_ID>",
		"<NAME>",
		"example.com",
		587,
		service.WithCreateSmtpMonitorTlsMode("none"),
		service.WithCreateSmtpMonitorRules([]string{"{\"method\":\"greaterThan\", \"attribute\":\"latency\", \"values\":[2000]}"}),
		service.WithCreateSmtpMonitorInterval(1),
		service.WithCreateSmtpMonitorTimeout(1),
		service.WithCreateSmtpMonitorCooldown(300),
		service.WithCreateSmtpMonitorEnabled(false),
		service.WithCreateSmtpMonitorNotify([]string{"example"}),
		service.WithCreateSmtpMonitorSeverity("info"),
		service.WithCreateSmtpMonitorEscalationId("<ESCALATION_ID>"),
	)
	fmt.Println(response, err)
}
```
