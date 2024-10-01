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
    response, error := functions.CreateDeployment(
        "{$example}",
        payload.NewPayloadFromFile("/path/to/file.png"),
        false,
        functions.WithCreateDeploymentEntrypoint("{$example}"),
        functions.WithCreateDeploymentCommands("{$example}"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
