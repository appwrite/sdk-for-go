```go
package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/v2/client"
    "github.com/appwrite/sdk-for-go/v2/sites"
)

client := client.New(
    client.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1")
    client.WithProject("<YOUR_PROJECT_ID>")
    client.WithKey("<YOUR_API_KEY>")
)

service := sites.New(client)

response, error := service.CreateDeployment(
    "<SITE_ID>",
    file.NewInputFile("/path/to/file.png", "file.png"),
    sites.WithCreateDeploymentInstallCommand("<INSTALL_COMMAND>"),
    sites.WithCreateDeploymentBuildCommand("<BUILD_COMMAND>"),
    sites.WithCreateDeploymentOutputDirectory("<OUTPUT_DIRECTORY>"),
    sites.WithCreateDeploymentActivate(false),
)
```
