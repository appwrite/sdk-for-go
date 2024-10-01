package main

import (
    "fmt"
	"github.com/appwrite/sdk-for-go/appwrite"
)

func main() {
	client := appwrite.NewClient(
        appwrite.WithEndpoint("https://cloud.appwrite.io/v1"), // Your API Endpoint
        appwrite.WithProject(""), // Your project ID
        appwrite.WithSession(""), // The user session to authenticate with
    )

    functions := appwrite.NewFunctions(client)
    response, error := functions.GetExecution(
        "{$example}",
        "{$example}",
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
