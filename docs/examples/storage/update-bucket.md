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

	response, err := service.UpdateBucket(
		"<BUCKET_ID>",
		"<NAME>",
		service.WithUpdateBucketPermissions([]string{"read(\"any\")"}),
		service.WithUpdateBucketFileSecurity(false),
		service.WithUpdateBucketEnabled(false),
		service.WithUpdateBucketMaximumFileSize(1),
		service.WithUpdateBucketAllowedFileExtensions([]string{"example"}),
		service.WithUpdateBucketCompression("none"),
		service.WithUpdateBucketEncryption(false),
		service.WithUpdateBucketAntivirus(false),
		service.WithUpdateBucketTransformations(false),
	)
	fmt.Println(response, err)
}
```
