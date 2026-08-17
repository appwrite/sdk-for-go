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

response, error := service.Update(
    "<DATABASE_ID>",
    mysql.WithUpdateName("<NAME>"),
    mysql.WithUpdateStatus("ready"),
    mysql.WithUpdateSpecification("<SPECIFICATION>"),
    mysql.WithUpdateReplicas(0),
    mysql.WithUpdateSyncMode("async"),
    mysql.WithUpdateNetworkIdleTimeoutSeconds(60),
    mysql.WithUpdateNetworkIPAllowlist([]string{}),
    mysql.WithUpdateIdleTimeoutMinutes(5),
    mysql.WithUpdatePitr(false),
    mysql.WithUpdatePitrRetentionDays(1),
    mysql.WithUpdateStorageAutoscaling(false),
    mysql.WithUpdateStorageAutoscalingThresholdPercent(50),
    mysql.WithUpdateStorageAutoscalingMaxGb(0),
    mysql.WithUpdateMetricsTraceSampleRate(0),
    mysql.WithUpdateMetricsSlowQueryLogThresholdMs(0),
    mysql.WithUpdateSqlApiEnabled(false),
    mysql.WithUpdateSqlApiAllowedStatements([]string{}),
    mysql.WithUpdateSqlApiMaxRows(1),
    mysql.WithUpdateSqlApiMaxBytes(1024),
    mysql.WithUpdateSqlApiTimeoutSeconds(1),
)
```
