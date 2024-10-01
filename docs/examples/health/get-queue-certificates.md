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

    health := appwrite.NewHealth(client)
    response, error := health.GetQueueCertificates(
        health.WithGetQueueCertificatesThreshold(0),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
