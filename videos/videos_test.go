package videos

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestVideos(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test List", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "videos": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": [],
            "bucketId": "5e5ea5c16897e",
            "fileId": "d5fg5ehg1c168g7c",
            "name": "Product demo",
            "size": 23647142,
            "format": "MPEG-4",
            "duration": 92810,
            "width": 960,
            "height": 544,
            "aspectRatio": "16:9",
            "videoCodec": "h264",
            "videoFormat": "AVC",
            "videoFormatProfile": "Baseline@L3.1",
            "videoBitRate": 564790,
            "videoFrameRate": 29.97,
            "videoFrameRateMode": "CFR",
            "audioCodec": "aac",
            "audioFormat": "AAC",
            "audioBitRate": 127999,
            "audioSampleRate": 48000,
            "timelineStatus": "ready",
            "timelineError": "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.List()
		if err != nil {
			t.Errorf("Method List failed: %v", err)
		}
	})

	t.Run("Test Create", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "bucketId": "5e5ea5c16897e",
    "fileId": "d5fg5ehg1c168g7c",
    "name": "Product demo",
    "size": 23647142,
    "format": "MPEG-4",
    "duration": 92810,
    "width": 960,
    "height": 544,
    "aspectRatio": "16:9",
    "videoCodec": "h264",
    "videoFormat": "AVC",
    "videoFormatProfile": "Baseline@L3.1",
    "videoBitRate": 564790,
    "videoFrameRate": 29.97,
    "videoFrameRateMode": "CFR",
    "audioCodec": "aac",
    "audioFormat": "AAC",
    "audioBitRate": 127999,
    "audioSampleRate": 48000,
    "timelineStatus": "ready",
    "timelineError": "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Create("<VIDEO_ID>", "<BUCKET_ID>", "<FILE_ID>")
		if err != nil {
			t.Errorf("Method Create failed: %v", err)
		}
	})

	t.Run("Test ListCodecs", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "codecs": [
        {
            "$id": "h264",
            "name": "H.264",
            "outputs": []
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListCodecs()
		if err != nil {
			t.Errorf("Method ListCodecs failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "bucketId": "5e5ea5c16897e",
    "fileId": "d5fg5ehg1c168g7c",
    "name": "Product demo",
    "size": 23647142,
    "format": "MPEG-4",
    "duration": 92810,
    "width": 960,
    "height": 544,
    "aspectRatio": "16:9",
    "videoCodec": "h264",
    "videoFormat": "AVC",
    "videoFormatProfile": "Baseline@L3.1",
    "videoBitRate": 564790,
    "videoFrameRate": 29.97,
    "videoFrameRateMode": "CFR",
    "audioCodec": "aac",
    "audioFormat": "AAC",
    "audioBitRate": 127999,
    "audioSampleRate": 48000,
    "timelineStatus": "ready",
    "timelineError": "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Get("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
		}
	})

	t.Run("Test Update", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "bucketId": "5e5ea5c16897e",
    "fileId": "d5fg5ehg1c168g7c",
    "name": "Product demo",
    "size": 23647142,
    "format": "MPEG-4",
    "duration": 92810,
    "width": 960,
    "height": 544,
    "aspectRatio": "16:9",
    "videoCodec": "h264",
    "videoFormat": "AVC",
    "videoFormatProfile": "Baseline@L3.1",
    "videoBitRate": 564790,
    "videoFrameRate": 29.97,
    "videoFrameRateMode": "CFR",
    "audioCodec": "aac",
    "audioFormat": "AAC",
    "audioBitRate": 127999,
    "audioSampleRate": 48000,
    "timelineStatus": "ready",
    "timelineError": "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Update("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method Update failed: %v", err)
		}
	})

	t.Run("Test Delete", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Delete("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})

	t.Run("Test ListCaptions", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "captions": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "videoId": "5e5ea5c16897e",
            "bucketId": "d5fg5ehg1c168g7c",
            "fileId": "c5fg5emg1c168grr",
            "name": "English",
            "language": "eng",
            "default": true,
            "embedded": true,
            "targetDuration": 93,
            "status": "ready",
            "error": "Caption is not a proper .srt file (Invalid timestamp delimiter at line 2)",
            "size": 4096
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListCaptions("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method ListCaptions failed: %v", err)
		}
	})

	t.Run("Test CreateCaption", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "videoId": "5e5ea5c16897e",
    "bucketId": "d5fg5ehg1c168g7c",
    "fileId": "c5fg5emg1c168grr",
    "name": "English",
    "language": "eng",
    "default": true,
    "embedded": true,
    "targetDuration": 93,
    "status": "ready",
    "error": "Caption is not a proper .srt file (Invalid timestamp delimiter at line 2)",
    "size": 4096
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateCaption("<VIDEO_ID>", "<BUCKET_ID>", "<FILE_ID>", "<NAME>", "aar")
		if err != nil {
			t.Errorf("Method CreateCaption failed: %v", err)
		}
	})

	t.Run("Test UpdateCaption", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "videoId": "5e5ea5c16897e",
    "bucketId": "d5fg5ehg1c168g7c",
    "fileId": "c5fg5emg1c168grr",
    "name": "English",
    "language": "eng",
    "default": true,
    "embedded": true,
    "targetDuration": 93,
    "status": "ready",
    "error": "Caption is not a proper .srt file (Invalid timestamp delimiter at line 2)",
    "size": 4096
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateCaption("<VIDEO_ID>", "<CAPTION_ID>")
		if err != nil {
			t.Errorf("Method UpdateCaption failed: %v", err)
		}
	})

	t.Run("Test DeleteCaption", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.DeleteCaption("<VIDEO_ID>", "<CAPTION_ID>")
		if err != nil {
			t.Errorf("Method DeleteCaption failed: %v", err)
		}
	})

	t.Run("Test GetCmafHlsManifest", func(t *testing.T) {
		mockData := []byte("image_data")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData)
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetCmafHlsManifest("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method GetCmafHlsManifest failed: %v", err)
		}
	})

	t.Run("Test GetCmafDashManifest", func(t *testing.T) {
		mockData := []byte("image_data")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData)
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetCmafDashManifest("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method GetCmafDashManifest failed: %v", err)
		}
	})

	t.Run("Test GetDashManifest", func(t *testing.T) {
		mockData := []byte("image_data")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData)
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetDashManifest("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method GetDashManifest failed: %v", err)
		}
	})

	t.Run("Test GetHlsManifest", func(t *testing.T) {
		mockData := []byte("image_data")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData)
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetHlsManifest("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method GetHlsManifest failed: %v", err)
		}
	})

	t.Run("Test GetCaptionManifest", func(t *testing.T) {
		mockData := []byte("image_data")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData)
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetCaptionManifest("<VIDEO_ID>", "hls", "<CAPTION_ID>")
		if err != nil {
			t.Errorf("Method GetCaptionManifest failed: %v", err)
		}
	})

	t.Run("Test ListRenditions", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "renditions": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "videoId": "5e5ea5c16897e",
            "profileId": "d5fg5ehg1c168g7c",
            "name": "1280X720@3679",
            "startedAt": "2020-10-15T06:38:00.000+00:00",
            "endedAt": "2020-10-15T06:38:00.000+00:00",
            "width": 1280,
            "height": 720,
            "videoBitRate": 3551000,
            "audioBitRate": 128000,
            "targetDuration": 6.006,
            "status": "ready",
            "error": "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)",
            "progress": 88,
            "output": "hls",
            "codec": "h264",
            "size": 18350080
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListRenditions("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method ListRenditions failed: %v", err)
		}
	})

	t.Run("Test CreateRendition", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "videoId": "5e5ea5c16897e",
    "profileId": "d5fg5ehg1c168g7c",
    "name": "1280X720@3679",
    "startedAt": "2020-10-15T06:38:00.000+00:00",
    "endedAt": "2020-10-15T06:38:00.000+00:00",
    "width": 1280,
    "height": 720,
    "videoBitRate": 3551000,
    "audioBitRate": 128000,
    "targetDuration": 6.006,
    "status": "ready",
    "error": "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)",
    "progress": 88,
    "output": "hls",
    "codec": "h264",
    "size": 18350080
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateRendition("<VIDEO_ID>", "<PROFILE_ID>", "hls")
		if err != nil {
			t.Errorf("Method CreateRendition failed: %v", err)
		}
	})

	t.Run("Test GetRendition", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "videoId": "5e5ea5c16897e",
    "profileId": "d5fg5ehg1c168g7c",
    "name": "1280X720@3679",
    "startedAt": "2020-10-15T06:38:00.000+00:00",
    "endedAt": "2020-10-15T06:38:00.000+00:00",
    "width": 1280,
    "height": 720,
    "videoBitRate": 3551000,
    "audioBitRate": 128000,
    "targetDuration": 6.006,
    "status": "ready",
    "error": "The requested file is not a valid video or audio file (ffprobe: Invalid data found when processing input)",
    "progress": 88,
    "output": "hls",
    "codec": "h264",
    "size": 18350080
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetRendition("<VIDEO_ID>", "<RENDITION_ID>")
		if err != nil {
			t.Errorf("Method GetRendition failed: %v", err)
		}
	})

	t.Run("Test DeleteRendition", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.DeleteRendition("<VIDEO_ID>", "<RENDITION_ID>")
		if err != nil {
			t.Errorf("Method DeleteRendition failed: %v", err)
		}
	})

	t.Run("Test GetThumbnail", func(t *testing.T) {
		mockData := []byte("image_data")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData)
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetThumbnail("<VIDEO_ID>", "<THUMBNAIL_ID>")
		if err != nil {
			t.Errorf("Method GetThumbnail failed: %v", err)
		}
	})

	t.Run("Test GetTimeline", func(t *testing.T) {
		mockData := []byte("image_data")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData)
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetTimeline("<VIDEO_ID>")
		if err != nil {
			t.Errorf("Method GetTimeline failed: %v", err)
		}
	})
}
