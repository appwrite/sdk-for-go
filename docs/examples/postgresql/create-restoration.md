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

response, error := service.CreateRestoration(
    "<DATABASE_ID>",
    postgresql.WithCreateRestorationType("backup"),
    postgresql.WithCreateRestorationBackupId("<BACKUP_ID>"),
    postgresql.WithCreateRestorationTargetDatabaseId("<TARGET_DATABASE_ID>"),
    postgresql.WithCreateRestorationTargetTime("2020-10-15T06:38:00.000+00:00"),
)
```
