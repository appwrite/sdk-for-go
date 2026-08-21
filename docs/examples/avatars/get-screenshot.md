```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/avatars"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithSession(""),
	)

	service := avatars.New(client)

	response, err := service.GetScreenshot(
		"https://example.com",
		service.WithGetScreenshotHeaders(map[string]interface{}{"Authorization": "Bearer token123", "X-Custom-Header": "value"}),
		service.WithGetScreenshotViewportWidth(1920),
		service.WithGetScreenshotViewportHeight(1080),
		service.WithGetScreenshotScale(2),
		service.WithGetScreenshotTheme("dark"),
		service.WithGetScreenshotUserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15"),
		service.WithGetScreenshotFullpage(true),
		service.WithGetScreenshotLocale("en-US"),
		service.WithGetScreenshotTimezone("America/New_York"),
		service.WithGetScreenshotLatitude(37.7749),
		service.WithGetScreenshotLongitude(-122.4194),
		service.WithGetScreenshotAccuracy(100),
		service.WithGetScreenshotTouch(true),
		service.WithGetScreenshotPermissions([]string{"geolocation", "notifications"}),
		service.WithGetScreenshotSleep(3),
		service.WithGetScreenshotWidth(800),
		service.WithGetScreenshotHeight(600),
		service.WithGetScreenshotQuality(85),
		service.WithGetScreenshotOutput("jpeg"),
	)
	fmt.Println(response, err)
}
```
