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

    databases := appwrite.NewDatabases(client)
    response, error := databases.CreateDocument(
        "{$example}",
        "{$example}",
        "{$example}",
        map[string]interface{}{},
        databases.WithCreateDocumentPermissions(interface{}{"read("any")"}),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
