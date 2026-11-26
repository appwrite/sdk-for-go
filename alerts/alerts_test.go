package alerts

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestAlerts(t *testing.T) {
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
    "alerts": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "monitorId": "5e5ea5c16897e",
            "projectId": "5e5ea5c16897e",
            "name": "Bandwidth spike",
            "status": "firing",
            "acknowledgedBy": "5e5ea5c16897e",
            "acknowledgedByName": "Jane Doe",
            "escalationId": "5e5ea5c16897e",
            "escalationStep": 1,
            "escalationRound": 0,
            "rules": [],
            "severity": "warning",
            "threshold": 5000000,
            "value": 7250000,
            "peak": 9100000,
            "type": "usage",
            "resourceType": "database",
            "resourceId": "5e5ea5c16897e",
            "notifications": 1
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
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "monitorId": "5e5ea5c16897e",
    "projectId": "5e5ea5c16897e",
    "name": "Bandwidth spike",
    "status": "firing",
    "acknowledgedBy": "5e5ea5c16897e",
    "acknowledgedByName": "Jane Doe",
    "escalationId": "5e5ea5c16897e",
    "escalationStep": 1,
    "escalationRound": 0,
    "rules": [],
    "severity": "warning",
    "threshold": 5000000,
    "value": 7250000,
    "peak": 9100000,
    "type": "usage",
    "resourceType": "database",
    "resourceId": "5e5ea5c16897e",
    "notifications": 1
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

		_, err := srv.Get("<ALERT_ID>")
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
    "monitorId": "5e5ea5c16897e",
    "projectId": "5e5ea5c16897e",
    "name": "Bandwidth spike",
    "status": "firing",
    "acknowledgedBy": "5e5ea5c16897e",
    "acknowledgedByName": "Jane Doe",
    "escalationId": "5e5ea5c16897e",
    "escalationStep": 1,
    "escalationRound": 0,
    "rules": [],
    "severity": "warning",
    "threshold": 5000000,
    "value": 7250000,
    "peak": 9100000,
    "type": "usage",
    "resourceType": "database",
    "resourceId": "5e5ea5c16897e",
    "notifications": 1
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

		_, err := srv.Update("<ALERT_ID>", true)
		if err != nil {
			t.Errorf("Method Update failed: %v", err)
		}
	})
}
