package main

import (
    "fmt"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/payload"
)

func main() {
	client := appwrite.NewClient(
        appwrite.WithEndpoint("https://cloud.appwrite.io/v1"), // Your API Endpoint
        appwrite.WithProject(""), // Your project ID
        appwrite.WithSession(""), // The user session to authenticate with
    )

    functions := appwrite.NewFunctions(client)
    response, error := functions.CreateExecution(
        "{$example}",
        functions.WithCreateExecutionBody(payload.NewPayloadFromJson(map[string]interface{}{ "x": "y" })),
        functions.WithCreateExecutionAsync(false),
        functions.WithCreateExecutionPath("{$example}"),
        functions.WithCreateExecutionMethod("{$example}"),
        functions.WithCreateExecutionHeaders(map[string]interface{}{}),
        functions.WithCreateExecutionScheduledAt(""),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
