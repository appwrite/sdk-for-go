package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/functions"
)

func main() {
    client := client.NewClient()

    client.SetEndpoint("https://cloud.appwrite.io/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetKey("") // Your secret API key

    service := functions.NewFunctions(client)
    response, error := service.Create(
        "<FUNCTION_ID>",
        "<NAME>",
        "node-14.5",
        functions.WithCreateExecute(interface{}{"any"}),
        functions.WithCreateEvents([]interface{}{}),
        functions.WithCreateSchedule(""),
        functions.WithCreateTimeout(1),
        functions.WithCreateEnabled(false),
        functions.WithCreateLogging(false),
        functions.WithCreateEntrypoint("<ENTRYPOINT>"),
        functions.WithCreateCommands("<COMMANDS>"),
        functions.WithCreateScopes([]interface{}{}),
        functions.WithCreateInstallationId("<INSTALLATION_ID>"),
        functions.WithCreateProviderRepositoryId("<PROVIDER_REPOSITORY_ID>"),
        functions.WithCreateProviderBranch("<PROVIDER_BRANCH>"),
        functions.WithCreateProviderSilentMode(false),
        functions.WithCreateProviderRootDirectory("<PROVIDER_ROOT_DIRECTORY>"),
        functions.WithCreateTemplateRepository("<TEMPLATE_REPOSITORY>"),
        functions.WithCreateTemplateOwner("<TEMPLATE_OWNER>"),
        functions.WithCreateTemplateRootDirectory("<TEMPLATE_ROOT_DIRECTORY>"),
        functions.WithCreateTemplateBranch("<TEMPLATE_BRANCH>"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
