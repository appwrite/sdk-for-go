```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v5/client"
    "github.com/appwrite/sdk-for-go/v5/advisor"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := advisor.New(client)

response, error := service.ListInsights(
    "<REPORT_ID>",
    advisor.WithListInsightsQueries([]string{}),
    advisor.WithListInsightsTotal(false),
)
```
