```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/analytics"
	"github.com/appwrite/sdk-for-go/v7/appwrite"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := analytics.New(client)

	response, err := service.ListMetrics(
		"<PROPERTY_ID>",
		service.WithListMetricsQueries([]string{"example"}),
		service.WithListMetricsInterval("1h"),
		service.WithListMetricsDimensions([]string{"example"}),
		service.WithListMetricsDateRange("<DATE_RANGE>"),
		service.WithListMetricsStartAt("2020-10-15T06:38:00.000+00:00"),
		service.WithListMetricsEndAt("2020-10-15T06:38:00.000+00:00"),
		service.WithListMetricsLimit(1),
	)
	fmt.Println(response, err)
}
```
