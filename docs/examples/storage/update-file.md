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
		appwrite.WithSession(""),
	)

	service := storage.New(client)

	response, err := service.UpdateFile(
		"<BUCKET_ID>",
		"<FILE_ID>",
		service.WithUpdateFileName("<NAME>"),
		service.WithUpdateFilePermissions([]string{"read(\"any\")"}),
	)
	fmt.Println(response, err)
}
```
