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
    response, error := functions.Update(
        "{$example}",
        "{$example}",
        functions.WithUpdateRuntime("{$example}"),
        functions.WithUpdateExecute(interface{}{"any"}),
        functions.WithUpdateEvents([]interface{}{}),
        functions.WithUpdateSchedule(""),
        functions.WithUpdateTimeout(1),
        functions.WithUpdateEnabled(false),
        functions.WithUpdateLogging(false),
        functions.WithUpdateEntrypoint("{$example}"),
        functions.WithUpdateCommands("{$example}"),
        functions.WithUpdateScopes([]interface{}{}),
        functions.WithUpdateInstallationId("{$example}"),
        functions.WithUpdateProviderRepositoryId("{$example}"),
        functions.WithUpdateProviderBranch("{$example}"),
        functions.WithUpdateProviderSilentMode(false),
        functions.WithUpdateProviderRootDirectory("{$example}"),
        functions.WithUpdateSpecification(""),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
