```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/postgresql"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := postgresql.New(client)

response, error := service.Update(
    "<DATABASE_ID>",
    postgresql.WithUpdateName("<NAME>"),
    postgresql.WithUpdateStatus("ready"),
    postgresql.WithUpdateSpecification("<SPECIFICATION>"),
    postgresql.WithUpdateReplicas(0),
    postgresql.WithUpdateSyncMode("async"),
    postgresql.WithUpdateNetworkIdleTimeoutSeconds(60),
    postgresql.WithUpdateNetworkIPAllowlist([]string{}),
    postgresql.WithUpdateIdleTimeoutMinutes(5),
    postgresql.WithUpdatePitr(false),
    postgresql.WithUpdatePitrRetentionDays(1),
    postgresql.WithUpdateStorageAutoscaling(false),
    postgresql.WithUpdateStorageAutoscalingThresholdPercent(50),
    postgresql.WithUpdateStorageAutoscalingMaxGb(0),
    postgresql.WithUpdateMetricsTraceSampleRate(0),
    postgresql.WithUpdateMetricsSlowQueryLogThresholdMs(0),
    postgresql.WithUpdateSqlApiEnabled(false),
    postgresql.WithUpdateSqlApiAllowedStatements([]string{}),
    postgresql.WithUpdateSqlApiMaxRows(1),
    postgresql.WithUpdateSqlApiMaxBytes(1024),
    postgresql.WithUpdateSqlApiTimeoutSeconds(1),
)
```
