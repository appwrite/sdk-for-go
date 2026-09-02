```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/storage"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := storage.New(client)

	response, err := service.CreateBucket(
		"<BUCKET_ID>",
		"<NAME>",
		service.WithCreateBucketPermissions([]string{"read(\"any\")"}),
		service.WithCreateBucketFileSecurity(false),
		service.WithCreateBucketEnabled(false),
		service.WithCreateBucketMaximumFileSize(1),
		service.WithCreateBucketAllowedFileExtensions([]string{"example"}),
		service.WithCreateBucketCompression("none"),
		service.WithCreateBucketEncryption(false),
		service.WithCreateBucketAntivirus(false),
		service.WithCreateBucketTransformations(false),
	)
	fmt.Println(response, err)
}
```
