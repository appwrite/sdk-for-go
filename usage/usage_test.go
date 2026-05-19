package usage

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v4/client"
)

func TestUsage(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test ListEvents", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "events": [
        {
            "metric": "bandwidth",
            "value": 5000,
            "time": "2026-04-09T12:00:00.000+00:00",
            "path": "/v1/storage/files",
            "method": "POST",
            "status": "201",
            "resourceType": "bucket",
            "resourceId": "abc123",
            "countryCode": "US",
            "userAgent": "AppwriteSDK/1.0"
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

		_, err := srv.ListEvents()
		if err != nil {
			t.Errorf("Method ListEvents failed: %v", err)
		}
	})

	t.Run("Test ListGauges", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "gauges": [
        {
            "metric": "users",
            "value": 1500,
            "time": "2026-04-09T12:00:00.000+00:00"
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

		_, err := srv.ListGauges()
		if err != nil {
			t.Errorf("Method ListGauges failed: %v", err)
		}
	})
}
