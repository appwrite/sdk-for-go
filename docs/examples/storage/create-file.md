```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/file"
	"github.com/appwrite/sdk-for-go/v7/storage"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := storage.New(client)

	response, err := service.CreateFile(
		"<BUCKET_ID>",
		"<FILE_ID>",
		file.NewInputFile("/path/to/file.png", "file.png"),
		service.WithCreateFilePermissions([]string{"read(\"any\")"}),
		service.WithCreateFileFolder("photos/2026"),
	)
	fmt.Println(response, err)
}
```
