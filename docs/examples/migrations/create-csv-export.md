```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/migrations"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := migrations.New(client)

	response, err := service.CreateCSVExport(
		"<DATABASE_ID>",
		"<COLLECTION_ID>",
		"<FILENAME>",
		service.WithCreateCSVExportColumns([]string{"example"}),
		service.WithCreateCSVExportQueries([]string{"example"}),
		service.WithCreateCSVExportDelimiter("<DELIMITER>"),
		service.WithCreateCSVExportEnclosure("<ENCLOSURE>"),
		service.WithCreateCSVExportEscape("<ESCAPE>"),
		service.WithCreateCSVExportHeader(false),
		service.WithCreateCSVExportNotify(false),
	)
	fmt.Println(response, err)
}
```
