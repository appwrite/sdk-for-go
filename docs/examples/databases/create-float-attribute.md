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
    response, error := databases.CreateFloatAttribute(
        "{$example}",
        "{$example}",
        "",
        false,
        databases.WithCreateFloatAttributeMin(0),
        databases.WithCreateFloatAttributeMax(0),
        databases.WithCreateFloatAttributeDefault(0),
        databases.WithCreateFloatAttributeArray(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
