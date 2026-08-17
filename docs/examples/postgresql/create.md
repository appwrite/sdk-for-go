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

response, error := service.Create(
    "<DATABASE_ID>",
    "<NAME>",
    postgresql.WithCreateVersion("17"),
    postgresql.WithCreateSpecification("<SPECIFICATION>"),
    postgresql.WithCreateReplicas(0),
    postgresql.WithCreateSyncMode("async"),
    postgresql.WithCreateNetworkIdleTimeoutSeconds(60),
    postgresql.WithCreateNetworkIPAllowlist([]string{}),
    postgresql.WithCreateIdleTimeoutMinutes(5),
    postgresql.WithCreatePitr(false),
    postgresql.WithCreatePitrRetentionDays(1),
    postgresql.WithCreateStorageAutoscaling(false),
    postgresql.WithCreateStorageAutoscalingThresholdPercent(50),
    postgresql.WithCreateStorageAutoscalingMaxGb(0),
)
```
