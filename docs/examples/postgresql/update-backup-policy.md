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

response, error := service.UpdateBackupPolicy(
    "<DATABASE_ID>",
    "<POLICY_ID>",
    postgresql.WithUpdateBackupPolicyName("<NAME>"),
    postgresql.WithUpdateBackupPolicySchedule(""),
    postgresql.WithUpdateBackupPolicyRetention(1),
    postgresql.WithUpdateBackupPolicyEnabled(false),
)
```
