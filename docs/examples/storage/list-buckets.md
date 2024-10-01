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
    response, error := storage.ListBuckets(
        storage.WithListBucketsQueries([]interface{}{}),
        storage.WithListBucketsSearch("{$example}"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
