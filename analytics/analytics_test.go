package analytics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestAnalytics(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test ListProperties", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "properties": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Website",
            "domain": "example.com",
            "timezone": "UTC",
            "enabled": true,
            "public": true,
            "allowedOrigins": [],
            "snippetId": "snp_a1b2c3d4e5"
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

		_, err := srv.ListProperties()
		if err != nil {
			t.Errorf("Method ListProperties failed: %v", err)
		}
	})

	t.Run("Test CreateProperty", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Website",
    "domain": "example.com",
    "timezone": "UTC",
    "enabled": true,
    "public": true,
    "allowedOrigins": [],
    "snippetId": "snp_a1b2c3d4e5"
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

		_, err := srv.CreateProperty("<PROPERTY_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateProperty failed: %v", err)
		}
	})

	t.Run("Test GetProperty", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Website",
    "domain": "example.com",
    "timezone": "UTC",
    "enabled": true,
    "public": true,
    "allowedOrigins": [],
    "snippetId": "snp_a1b2c3d4e5"
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

		_, err := srv.GetProperty("<PROPERTY_ID>")
		if err != nil {
			t.Errorf("Method GetProperty failed: %v", err)
		}
	})

	t.Run("Test UpdateProperty", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Website",
    "domain": "example.com",
    "timezone": "UTC",
    "enabled": true,
    "public": true,
    "allowedOrigins": [],
    "snippetId": "snp_a1b2c3d4e5"
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

		_, err := srv.UpdateProperty("<PROPERTY_ID>")
		if err != nil {
			t.Errorf("Method UpdateProperty failed: %v", err)
		}
	})

	t.Run("Test DeleteProperty", func(t *testing.T) {
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

		_, err := srv.DeleteProperty("<PROPERTY_ID>")
		if err != nil {
			t.Errorf("Method DeleteProperty failed: %v", err)
		}
	})

	t.Run("Test CreateEvent", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
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

		_, err := srv.CreateEvent("<PROPERTY_ID>", "<NAME>", "https://example.com")
		if err != nil {
			t.Errorf("Method CreateEvent failed: %v", err)
		}
	})

	t.Run("Test ListMetrics", func(t *testing.T) {
		mockResponse := `
{
    "total": 30,
    "metrics": [
        {
            "visitors": 1234,
            "sessions": 4567,
            "events": 8910
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

		_, err := srv.ListMetrics("<PROPERTY_ID>")
		if err != nil {
			t.Errorf("Method ListMetrics failed: %v", err)
		}
	})
}
