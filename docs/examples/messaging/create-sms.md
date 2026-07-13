```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v6/client"
    "github.com/appwrite/sdk-for-go/v6/messaging"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := messaging.New(client)

response, error := service.CreateSMS(
    "<MESSAGE_ID>",
    "<CONTENT>",
    messaging.WithCreateSMSTopics([]string{}),
    messaging.WithCreateSMSUsers([]string{}),
    messaging.WithCreateSMSTargets([]string{}),
    messaging.WithCreateSMSDraft(false),
    messaging.WithCreateSMSScheduledAt("2020-10-15T06:38:00.000+00:00"),
)
```
