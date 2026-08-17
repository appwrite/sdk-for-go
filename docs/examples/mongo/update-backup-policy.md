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

response, error := service.UpdateBackupPolicy(
    "<DATABASE_ID>",
    "<POLICY_ID>",
    mongo.WithUpdateBackupPolicyName("<NAME>"),
    mongo.WithUpdateBackupPolicySchedule(""),
    mongo.WithUpdateBackupPolicyRetention(1),
    mongo.WithUpdateBackupPolicyEnabled(false),
)
```
