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
    response, error := databases.UpdateCollection(
        "{$example}",
        "{$example}",
        "{$example}",
        databases.WithUpdateCollectionPermissions(interface{}{"read("any")"}),
        databases.WithUpdateCollectionDocumentSecurity(false),
        databases.WithUpdateCollectionEnabled(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
