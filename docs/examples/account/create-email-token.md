package main

import (
    "fmt"
	"github.com/appwrite/sdk-for-go/appwrite"
)

func main() {
	client := appwrite.NewClient(
        appwrite.WithEndpoint("https://cloud.appwrite.io/v1"), // Your API Endpoint
        appwrite.WithProject(""), // Your project ID
    )

    account := appwrite.NewAccount(client)
    response, error := account.CreateEmailToken(
        "{$example}",
        "{$example}",
        account.WithCreateEmailTokenPhrase(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
