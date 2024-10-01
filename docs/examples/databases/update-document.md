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
    response, error := databases.UpdateDocument(
        "{$example}",
        "{$example}",
        "{$example}",
        databases.WithUpdateDocumentData(map[string]interface{}{}),
        databases.WithUpdateDocumentPermissions(interface{}{"read("any")"}),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
