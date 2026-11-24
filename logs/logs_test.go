package logs

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestLogs(t *testing.T) {
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
    "logs": [
        {
            "$id": "5e5ea5c16897e",
            "resourceType": "functions",
            "resourceId": "610fc2f985ee0",
            "resourceInternalId": "42",
            "deploymentId": "610fc2f985ee0",
            "region": "fra",
            "timestamp": "2020-10-15T06:38:00.000+00:00",
            "stream": "stdout",
            "message": "Listening on port 3000"
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

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "resourceType": "functions",
    "resourceId": "610fc2f985ee0",
    "resourceInternalId": "42",
    "deploymentId": "610fc2f985ee0",
    "region": "fra",
    "timestamp": "2020-10-15T06:38:00.000+00:00",
    "stream": "stdout",
    "message": "Listening on port 3000"
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

		_, err := srv.Get("<LOG_ID>", "functions", "<RESOURCE_ID>")
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
		}
	})
}
