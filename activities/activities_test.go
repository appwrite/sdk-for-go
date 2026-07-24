package activities

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v6/client"
)

func TestActivities(t *testing.T) {
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
            "$id": "5e5ea5c16897e",
            "actorType": "user",
            "actorId": "610fc2f985ee0",
            "actorEmail": "john@appwrite.io",
            "actorName": "John Doe",
            "resourceParent": "database/ID",
            "resourceType": "collection",
            "resourceId": "610fc2f985ee0",
            "resource": "collections/610fc2f985ee0",
            "event": "account.sessions.create",
            "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",
            "ip": "127.0.0.1",
            "mode": "admin",
            "country": "US",
            "continentCode": "NA",
            "city": "Mountain View",
            "subdivisions": "California",
            "isp": "Google",
            "autonomousSystemNumber": "15169",
            "autonomousSystemOrganization": "GOOGLE",
            "connectionType": "cable",
            "connectionUsageType": "residential",
            "connectionOrganization": "Google LLC",
            "time": "2020-10-15T06:38:00.000+00:00",
            "projectId": "610fc2f985ee0",
            "teamId": "610fc2f985ee0",
            "hostname": "appwrite.io",
            "sdk": "web",
            "sdkVersion": "14.0.0"
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

	t.Run("Test GetEvent", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "actorType": "user",
    "actorId": "610fc2f985ee0",
    "actorEmail": "john@appwrite.io",
    "actorName": "John Doe",
    "resourceParent": "database/ID",
    "resourceType": "collection",
    "resourceId": "610fc2f985ee0",
    "resource": "collections/610fc2f985ee0",
    "event": "account.sessions.create",
    "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",
    "ip": "127.0.0.1",
    "mode": "admin",
    "country": "US",
    "continentCode": "NA",
    "city": "Mountain View",
    "subdivisions": "California",
    "isp": "Google",
    "autonomousSystemNumber": "15169",
    "autonomousSystemOrganization": "GOOGLE",
    "connectionType": "cable",
    "connectionUsageType": "residential",
    "connectionOrganization": "Google LLC",
    "time": "2020-10-15T06:38:00.000+00:00",
    "projectId": "610fc2f985ee0",
    "teamId": "610fc2f985ee0",
    "hostname": "appwrite.io",
    "sdk": "web",
    "sdkVersion": "14.0.0"
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

		_, err := srv.GetEvent("<EVENT_ID>")
		if err != nil {
			t.Errorf("Method GetEvent failed: %v", err)
		}
	})
}
