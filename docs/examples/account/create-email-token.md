```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v2/client"
    "github.com/appwrite/sdk-for-go/v2/account"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithSession("")
)

service := account.New(client)

response, error := service.CreateEmailToken(
    "<USER_ID>",
    "email@example.com",
    account.WithCreateEmailTokenPhrase(false),
)
```
