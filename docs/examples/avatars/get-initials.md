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
    response, error := avatars.GetInitials(
        avatars.WithGetInitialsName("{$example}"),
        avatars.WithGetInitialsWidth(0),
        avatars.WithGetInitialsHeight(0),
        avatars.WithGetInitialsBackground(""),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
