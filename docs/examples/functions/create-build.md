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

    functions := appwrite.NewFunctions(client)
    response, error := functions.CreateBuild(
        "{$example}",
        "{$example}",
        functions.WithCreateBuildBuildId("{$example}"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
