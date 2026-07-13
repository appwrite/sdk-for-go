```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v6/client"
    "github.com/appwrite/sdk-for-go/v6/storage"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := storage.New(client)

response, error := service.CreateBucket(
    "<BUCKET_ID>",
    "<NAME>",
    storage.WithCreateBucketPermissions([]string{"read("any")"}),
    storage.WithCreateBucketFileSecurity(false),
    storage.WithCreateBucketEnabled(false),
    storage.WithCreateBucketMaximumFileSize(1),
    storage.WithCreateBucketAllowedFileExtensions([]string{}),
    storage.WithCreateBucketCompression("none"),
    storage.WithCreateBucketEncryption(false),
    storage.WithCreateBucketAntivirus(false),
    storage.WithCreateBucketTransformations(false),
)
```
