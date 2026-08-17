```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/mysql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := mysql.New(client)

response, error := service.Create(
    "<DATABASE_ID>",
    "<NAME>",
    mysql.WithCreateVersion("17"),
    mysql.WithCreateSpecification("<SPECIFICATION>"),
    mysql.WithCreateReplicas(0),
    mysql.WithCreateSyncMode("async"),
    mysql.WithCreateNetworkIdleTimeoutSeconds(60),
    mysql.WithCreateNetworkIPAllowlist([]string{}),
    mysql.WithCreateIdleTimeoutMinutes(5),
    mysql.WithCreatePitr(false),
    mysql.WithCreatePitrRetentionDays(1),
    mysql.WithCreateStorageAutoscaling(false),
    mysql.WithCreateStorageAutoscalingThresholdPercent(50),
    mysql.WithCreateStorageAutoscalingMaxGb(0),
)
```
