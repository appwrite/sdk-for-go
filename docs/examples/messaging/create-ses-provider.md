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

response, error := service.CreateSesProvider(
    "<PROVIDER_ID>",
    "<NAME>",
    messaging.WithCreateSesProviderAccessKey("<ACCESS_KEY>"),
    messaging.WithCreateSesProviderSecretKey("<SECRET_KEY>"),
    messaging.WithCreateSesProviderRegion("<REGION>"),
    messaging.WithCreateSesProviderFromName("<FROM_NAME>"),
    messaging.WithCreateSesProviderFromEmail("email@example.com"),
    messaging.WithCreateSesProviderReplyToName("<REPLY_TO_NAME>"),
    messaging.WithCreateSesProviderReplyToEmail("email@example.com"),
    messaging.WithCreateSesProviderEnabled(false),
)
```
