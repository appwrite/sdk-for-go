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
    response, error := functions.Create(
        "{$example}",
        "{$example}",
        "{$example}",
        functions.WithCreateExecute(interface{}{"any"}),
        functions.WithCreateEvents([]interface{}{}),
        functions.WithCreateSchedule(""),
        functions.WithCreateTimeout(1),
        functions.WithCreateEnabled(false),
        functions.WithCreateLogging(false),
        functions.WithCreateEntrypoint("{$example}"),
        functions.WithCreateCommands("{$example}"),
        functions.WithCreateScopes([]interface{}{}),
        functions.WithCreateInstallationId("{$example}"),
        functions.WithCreateProviderRepositoryId("{$example}"),
        functions.WithCreateProviderBranch("{$example}"),
        functions.WithCreateProviderSilentMode(false),
        functions.WithCreateProviderRootDirectory("{$example}"),
        functions.WithCreateTemplateRepository("{$example}"),
        functions.WithCreateTemplateOwner("{$example}"),
        functions.WithCreateTemplateRootDirectory("{$example}"),
        functions.WithCreateTemplateVersion("{$example}"),
        functions.WithCreateSpecification(""),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
