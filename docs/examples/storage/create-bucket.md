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

    storage := appwrite.NewStorage(client)
    response, error := storage.CreateBucket(
        "{$example}",
        "{$example}",
        storage.WithCreateBucketPermissions(interface{}{"read("any")"}),
        storage.WithCreateBucketFileSecurity(false),
        storage.WithCreateBucketEnabled(false),
        storage.WithCreateBucketMaximumFileSize(1),
        storage.WithCreateBucketAllowedFileExtensions([]interface{}{}),
        storage.WithCreateBucketCompression("{$example}"),
        storage.WithCreateBucketEncryption(false),
        storage.WithCreateBucketAntivirus(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
