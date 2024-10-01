package main

import (
    "fmt"
	"github.com/appwrite/sdk-for-go/appwrite"
)

func main() {
	client := appwrite.NewClient(
        appwrite.WithEndpoint("https://cloud.appwrite.io/v1"), // Your API Endpoint
        appwrite.WithProject(""), // Your project ID
        appwrite.WithJWT(""), // Your secret JSON Web Token
    )

    messaging := appwrite.NewMessaging(client)
    response, error := messaging.CreateSubscriber(
        "{$example}",
        "{$example}",
        "{$example}",
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
