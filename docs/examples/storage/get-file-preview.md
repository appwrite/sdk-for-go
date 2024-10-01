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
    response, error := storage.GetFilePreview(
        "{$example}",
        "{$example}",
        storage.WithGetFilePreviewWidth(0),
        storage.WithGetFilePreviewHeight(0),
        storage.WithGetFilePreviewGravity("{$example}"),
        storage.WithGetFilePreviewQuality(0),
        storage.WithGetFilePreviewBorderWidth(0),
        storage.WithGetFilePreviewBorderColor(""),
        storage.WithGetFilePreviewBorderRadius(0),
        storage.WithGetFilePreviewOpacity(0),
        storage.WithGetFilePreviewRotation(-360),
        storage.WithGetFilePreviewBackground(""),
        storage.WithGetFilePreviewOutput("{$example}"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
