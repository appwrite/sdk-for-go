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
    response, error := databases.CreateIntegerAttribute(
        "{$example}",
        "{$example}",
        "",
        false,
        databases.WithCreateIntegerAttributeMin(0),
        databases.WithCreateIntegerAttributeMax(0),
        databases.WithCreateIntegerAttributeDefault(0),
        databases.WithCreateIntegerAttributeArray(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
