package storage

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"

	"github.com/appwrite/sdk-for-go/v2/client"
	"github.com/appwrite/sdk-for-go/v2/file")

func TestStorage(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	_, _ = tmpFile.WriteString("test content")
	_ = tmpFile.Close()
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test ListBuckets", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "buckets": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": [],
            "fileSecurity": true,
            "name": "Documents",
            "enabled": true,
            "maximumFileSize": 100,
            "allowedFileExtensions": [],
            "compression": "gzip",
            "encryption": true,
            "antivirus": true,
            "transformations": true,
            "totalSize": 128
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

		_, err := srv.ListBuckets()
		if err != nil {
			t.Errorf("Method ListBuckets failed: %v", err)
		}
	})

	t.Run("Test CreateBucket", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "fileSecurity": true,
    "name": "Documents",
    "enabled": true,
    "maximumFileSize": 100,
    "allowedFileExtensions": [],
    "compression": "gzip",
    "encryption": true,
    "antivirus": true,
    "transformations": true,
    "totalSize": 128
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

		_, err := srv.CreateBucket("<BUCKET_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateBucket failed: %v", err)
		}
	})

	t.Run("Test GetBucket", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "fileSecurity": true,
    "name": "Documents",
    "enabled": true,
    "maximumFileSize": 100,
    "allowedFileExtensions": [],
    "compression": "gzip",
    "encryption": true,
    "antivirus": true,
    "transformations": true,
    "totalSize": 128
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

		_, err := srv.GetBucket("<BUCKET_ID>")
		if err != nil {
			t.Errorf("Method GetBucket failed: %v", err)
		}
	})

	t.Run("Test UpdateBucket", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "fileSecurity": true,
    "name": "Documents",
    "enabled": true,
    "maximumFileSize": 100,
    "allowedFileExtensions": [],
    "compression": "gzip",
    "encryption": true,
    "antivirus": true,
    "transformations": true,
    "totalSize": 128
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateBucket("<BUCKET_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method UpdateBucket failed: %v", err)
		}
	})

	t.Run("Test DeleteBucket", func(t *testing.T) {
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

		_, err := srv.DeleteBucket("<BUCKET_ID>")
		if err != nil {
			t.Errorf("Method DeleteBucket failed: %v", err)
		}
	})

	t.Run("Test ListFiles", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "files": [
        {
            "$id": "5e5ea5c16897e",
            "bucketId": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": [],
            "name": "Pink.png",
            "signature": "5d529fd02b544198ae075bd57c1762bb",
            "mimeType": "image/png",
            "sizeOriginal": 17890,
            "chunksTotal": 17890,
            "chunksUploaded": 17890,
            "encryption": true,
            "compression": "gzip"
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

		_, err := srv.ListFiles("<BUCKET_ID>")
		if err != nil {
			t.Errorf("Method ListFiles failed: %v", err)
		}
	})

	t.Run("Test CreateFile", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "bucketId": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "name": "Pink.png",
    "signature": "5d529fd02b544198ae075bd57c1762bb",
    "mimeType": "image/png",
    "sizeOriginal": 17890,
    "chunksTotal": 17890,
    "chunksUploaded": 17890,
    "encryption": true,
    "compression": "gzip"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" && "POST" != "GET" {
				// Handle file upload resume check
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"chunksUploaded": 0}`))
				return
			}
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateFile("<BUCKET_ID>", "<FILE_ID>", file.NewInputFile(tmpFile.Name(), "test.txt"))
		if err != nil {
			t.Errorf("Method CreateFile failed: %v", err)
		}
	})

	t.Run("Test GetFile", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "bucketId": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "name": "Pink.png",
    "signature": "5d529fd02b544198ae075bd57c1762bb",
    "mimeType": "image/png",
    "sizeOriginal": 17890,
    "chunksTotal": 17890,
    "chunksUploaded": 17890,
    "encryption": true,
    "compression": "gzip"
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

		_, err := srv.GetFile("<BUCKET_ID>", "<FILE_ID>")
		if err != nil {
			t.Errorf("Method GetFile failed: %v", err)
		}
	})

	t.Run("Test UpdateFile", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "bucketId": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "name": "Pink.png",
    "signature": "5d529fd02b544198ae075bd57c1762bb",
    "mimeType": "image/png",
    "sizeOriginal": 17890,
    "chunksTotal": 17890,
    "chunksUploaded": 17890,
    "encryption": true,
    "compression": "gzip"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateFile("<BUCKET_ID>", "<FILE_ID>")
		if err != nil {
			t.Errorf("Method UpdateFile failed: %v", err)
		}
	})

	t.Run("Test DeleteFile", func(t *testing.T) {
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

		_, err := srv.DeleteFile("<BUCKET_ID>", "<FILE_ID>")
		if err != nil {
			t.Errorf("Method DeleteFile failed: %v", err)
		}
	})

	t.Run("Test GetFileDownload", func(t *testing.T) {
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

		_, err := srv.GetFileDownload("<BUCKET_ID>", "<FILE_ID>")
		if err != nil {
			t.Errorf("Method GetFileDownload failed: %v", err)
		}
	})

	t.Run("Test GetFilePreview", func(t *testing.T) {
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

		_, err := srv.GetFilePreview("<BUCKET_ID>", "<FILE_ID>")
		if err != nil {
			t.Errorf("Method GetFilePreview failed: %v", err)
		}
	})

	t.Run("Test GetFileView", func(t *testing.T) {
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

		_, err := srv.GetFileView("<BUCKET_ID>", "<FILE_ID>")
		if err != nil {
			t.Errorf("Method GetFileView failed: %v", err)
		}
	})
}
