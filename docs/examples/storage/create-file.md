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

    storage := appwrite.NewStorage(client)
    response, error := storage.CreateFile(
        "{$example}",
        "{$example}",
        payload.NewPayloadFromFile("/path/to/file.png"),
        storage.WithCreateFilePermissions(interface{}{"read("any")"}),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
