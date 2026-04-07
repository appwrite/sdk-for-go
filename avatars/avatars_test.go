package avatars

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v2/client"
)

func TestAvatars(t *testing.T) {
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
