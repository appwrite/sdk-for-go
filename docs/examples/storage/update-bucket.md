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
    response, error := storage.UpdateBucket(
        "{$example}",
        "{$example}",
        storage.WithUpdateBucketPermissions(interface{}{"read("any")"}),
        storage.WithUpdateBucketFileSecurity(false),
        storage.WithUpdateBucketEnabled(false),
        storage.WithUpdateBucketMaximumFileSize(1),
        storage.WithUpdateBucketAllowedFileExtensions([]interface{}{}),
        storage.WithUpdateBucketCompression("{$example}"),
        storage.WithUpdateBucketEncryption(false),
        storage.WithUpdateBucketAntivirus(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
