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

response, error := service.UpdateSesProvider(
    "<PROVIDER_ID>",
    messaging.WithUpdateSesProviderName("<NAME>"),
    messaging.WithUpdateSesProviderEnabled(false),
    messaging.WithUpdateSesProviderAccessKey("<ACCESS_KEY>"),
    messaging.WithUpdateSesProviderSecretKey("<SECRET_KEY>"),
    messaging.WithUpdateSesProviderRegion("<REGION>"),
    messaging.WithUpdateSesProviderFromName("<FROM_NAME>"),
    messaging.WithUpdateSesProviderFromEmail("email@example.com"),
    messaging.WithUpdateSesProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithUpdateSesProviderReplyToEmail("<REPLY_TO_EMAIL>"),
)
```
