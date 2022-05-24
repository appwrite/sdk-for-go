package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go"
)

func main() {
    var client := appwrite.Client{}

    client.setEndpoint("url") // url to your appwrite instance
    client.SetProject("5df5acd0d48c2") // Your project ID
    client.SetKey("919c2d18fb5d4...a2ae413da83346ad2") // Your secret API key

    appwrite_db := appwrite.NewDatabase(client)
    
    response, error :=appwrite_db.CreateCollection("unique()", "Collection_name", "document", []string{"role:all"}, []string{"role:all"})
    fmt.Println(response)
}