package avatars

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/file"
)

func TestAvatars(t *testing.T) {
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

	t.Run("Test GetBrowser", func(t *testing.T) {
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

		_, err := srv.GetBrowser("aa")
		if err != nil {
			t.Errorf("Method GetBrowser failed: %v", err)
		}
	})

	t.Run("Test GetCreditCard", func(t *testing.T) {
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

		_, err := srv.GetCreditCard("amex")
		if err != nil {
			t.Errorf("Method GetCreditCard failed: %v", err)
		}
	})

	t.Run("Test GetFavicon", func(t *testing.T) {
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

		_, err := srv.GetFavicon("https://example.com")
		if err != nil {
			t.Errorf("Method GetFavicon failed: %v", err)
		}
	})

	t.Run("Test GetFlag", func(t *testing.T) {
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

		_, err := srv.GetFlag("af")
		if err != nil {
			t.Errorf("Method GetFlag failed: %v", err)
		}
	})

	t.Run("Test GetImage", func(t *testing.T) {
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

		_, err := srv.GetImage("https://example.com")
		if err != nil {
			t.Errorf("Method GetImage failed: %v", err)
		}
	})

	t.Run("Test GetInitials", func(t *testing.T) {
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

		_, err := srv.GetInitials()
		if err != nil {
			t.Errorf("Method GetInitials failed: %v", err)
		}
	})

	t.Run("Test GetPhoto", func(t *testing.T) {
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

		_, err := srv.GetPhoto()
		if err != nil {
			t.Errorf("Method GetPhoto failed: %v", err)
		}
	})

	t.Run("Test UpdatePhoto", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" && "PUT" != "GET" {
				// Handle file upload resume check
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"chunksUploaded": 0}`))
				return
			}
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdatePhoto(file.NewInputFile(tmpFile.Name(), "test.txt"))
		if err != nil {
			t.Errorf("Method UpdatePhoto failed: %v", err)
		}
	})

	t.Run("Test DeletePhoto", func(t *testing.T) {
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

		_, err := srv.DeletePhoto()
		if err != nil {
			t.Errorf("Method DeletePhoto failed: %v", err)
		}
	})

	t.Run("Test GetQR", func(t *testing.T) {
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

		_, err := srv.GetQR("<TEXT>")
		if err != nil {
			t.Errorf("Method GetQR failed: %v", err)
		}
	})

	t.Run("Test GetScreenshot", func(t *testing.T) {
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

		_, err := srv.GetScreenshot("https://example.com")
		if err != nil {
			t.Errorf("Method GetScreenshot failed: %v", err)
		}
	})
}
