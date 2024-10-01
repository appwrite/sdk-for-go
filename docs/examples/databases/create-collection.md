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

    databases := appwrite.NewDatabases(client)
    response, error := databases.CreateCollection(
        "{$example}",
        "{$example}",
        "{$example}",
        databases.WithCreateCollectionPermissions(interface{}{"read("any")"}),
        databases.WithCreateCollectionDocumentSecurity(false),
        databases.WithCreateCollectionEnabled(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
