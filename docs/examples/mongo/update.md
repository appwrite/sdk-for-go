```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v7/client"
    "github.com/appwrite/sdk-for-go/v7/mongo"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := mongo.New(client)

response, error := service.Update(
    "<DATABASE_ID>",
    mongo.WithUpdateName("<NAME>"),
    mongo.WithUpdateStatus("ready"),
    mongo.WithUpdateSpecification("<SPECIFICATION>"),
    mongo.WithUpdateReplicas(0),
    mongo.WithUpdateSyncMode("async"),
    mongo.WithUpdateNetworkIdleTimeoutSeconds(60),
    mongo.WithUpdateNetworkIPAllowlist([]string{}),
    mongo.WithUpdateIdleTimeoutMinutes(5),
    mongo.WithUpdatePitr(false),
    mongo.WithUpdatePitrRetentionDays(1),
    mongo.WithUpdateStorageAutoscaling(false),
    mongo.WithUpdateStorageAutoscalingThresholdPercent(50),
    mongo.WithUpdateStorageAutoscalingMaxGb(0),
    mongo.WithUpdateMetricsTraceSampleRate(0),
    mongo.WithUpdateMetricsSlowQueryLogThresholdMs(0),
    mongo.WithUpdateSqlApiEnabled(false),
    mongo.WithUpdateSqlApiAllowedStatements([]string{}),
    mongo.WithUpdateSqlApiMaxRows(1),
    mongo.WithUpdateSqlApiMaxBytes(1024),
    mongo.WithUpdateSqlApiTimeoutSeconds(1),
)
```
