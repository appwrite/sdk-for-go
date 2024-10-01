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

    avatars := appwrite.NewAvatars(client)
    response, error := avatars.GetQR(
        "{$example}",
        avatars.WithGetQRSize(1),
        avatars.WithGetQRMargin(0),
        avatars.WithGetQRDownload(false),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
