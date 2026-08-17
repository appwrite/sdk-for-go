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

response, error := service.UpdateBackupPolicy(
    "<DATABASE_ID>",
    "<POLICY_ID>",
    mysql.WithUpdateBackupPolicyName("<NAME>"),
    mysql.WithUpdateBackupPolicySchedule(""),
    mysql.WithUpdateBackupPolicyRetention(1),
    mysql.WithUpdateBackupPolicyEnabled(false),
)
```
