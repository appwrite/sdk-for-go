package main

import (
    "fmt"
	"github.com/appwrite/sdk-for-go/appwrite"
)

func main() {
	client := appwrite.NewClient(
        appwrite.WithEndpoint("https://cloud.appwrite.io/v1"), // Your API Endpoint
        appwrite.WithProject(""), // Your project ID
        appwrite.WithKey(""), // Your secret API key
    )

    users := appwrite.NewUsers(client)
    response, error := users.CreateToken(
        "{$example}",
        users.WithCreateTokenLength(4),
        users.WithCreateTokenExpire(60),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
