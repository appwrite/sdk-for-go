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
    response, error := account.CreateOAuth2Token(
        "{$example}",
        account.WithCreateOAuth2TokenSuccess("{$example}"),
        account.WithCreateOAuth2TokenFailure("{$example}"),
        account.WithCreateOAuth2TokenScopes([]interface{}{}),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
