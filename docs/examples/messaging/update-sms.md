```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v2/client"
    "github.com/appwrite/sdk-for-go/v2/messaging"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := messaging.New(client)

response, error := service.UpdateSMS(
    "<MESSAGE_ID>",
    messaging.WithUpdateSMSTopics([]interface{}{}),
    messaging.WithUpdateSMSUsers([]interface{}{}),
    messaging.WithUpdateSMSTargets([]interface{}{}),
    messaging.WithUpdateSMSContent("<CONTENT>"),
    messaging.WithUpdateSMSDraft(false),
    messaging.WithUpdateSMSScheduledAt("2020-10-15T06:38:00.000+00:00"),
)
```
