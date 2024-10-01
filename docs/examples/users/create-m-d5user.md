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
    response, error := users.CreateMD5User(
        "{$example}",
        "{$example}",
        "{$example}",
        users.WithCreateMD5UserName("{$example}"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
