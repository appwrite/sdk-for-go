package backups

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v2/client"
)

func TestBackups(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test ListArchives", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "archives": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "policyId": "did8jx6ws45jana098ab7",
            "size": 100000,
            "status": "completed",
            "startedAt": "2020-10-15T06:38:00.000+00:00",
            "migrationId": "did8jx6ws45jana098ab7",
            "services": [],
            "resources": []
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

		_, err := srv.ListArchives()
		if err != nil {
			t.Errorf("Method ListArchives failed: %v", err)
		}
	})

	t.Run("Test CreateArchive", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "policyId": "did8jx6ws45jana098ab7",
    "size": 100000,
    "status": "completed",
    "startedAt": "2020-10-15T06:38:00.000+00:00",
    "migrationId": "did8jx6ws45jana098ab7",
    "services": [],
    "resources": []
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

		_, err := srv.CreateArchive([]string{})
		if err != nil {
			t.Errorf("Method CreateArchive failed: %v", err)
		}
	})

	t.Run("Test GetArchive", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "policyId": "did8jx6ws45jana098ab7",
    "size": 100000,
    "status": "completed",
    "startedAt": "2020-10-15T06:38:00.000+00:00",
    "migrationId": "did8jx6ws45jana098ab7",
    "services": [],
    "resources": []
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

		_, err := srv.GetArchive("<ARCHIVE_ID>")
		if err != nil {
			t.Errorf("Method GetArchive failed: %v", err)
		}
	})

	t.Run("Test DeleteArchive", func(t *testing.T) {
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

		_, err := srv.DeleteArchive("<ARCHIVE_ID>")
		if err != nil {
			t.Errorf("Method DeleteArchive failed: %v", err)
		}
	})

	t.Run("Test ListPolicies", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "policies": [
        {
            "$id": "5e5ea5c16897e",
            "name": "Hourly backups",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "services": [],
            "resources": [],
            "retention": 7,
            "schedule": "0 * * * *",
            "enabled": true
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

		_, err := srv.ListPolicies()
		if err != nil {
			t.Errorf("Method ListPolicies failed: %v", err)
		}
	})

	t.Run("Test CreatePolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "name": "Hourly backups",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "services": [],
    "resources": [],
    "retention": 7,
    "schedule": "0 * * * *",
    "enabled": true
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

		_, err := srv.CreatePolicy("<POLICY_ID>", []string{}, 1, "")
		if err != nil {
			t.Errorf("Method CreatePolicy failed: %v", err)
		}
	})

	t.Run("Test GetPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "name": "Hourly backups",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "services": [],
    "resources": [],
    "retention": 7,
    "schedule": "0 * * * *",
    "enabled": true
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

		_, err := srv.GetPolicy("<POLICY_ID>")
		if err != nil {
			t.Errorf("Method GetPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdatePolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "name": "Hourly backups",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "services": [],
    "resources": [],
    "retention": 7,
    "schedule": "0 * * * *",
    "enabled": true
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

		_, err := srv.UpdatePolicy("<POLICY_ID>")
		if err != nil {
			t.Errorf("Method UpdatePolicy failed: %v", err)
		}
	})

	t.Run("Test DeletePolicy", func(t *testing.T) {
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

		_, err := srv.DeletePolicy("<POLICY_ID>")
		if err != nil {
			t.Errorf("Method DeletePolicy failed: %v", err)
		}
	})

	t.Run("Test CreateRestoration", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "archiveId": "did8jx6ws45jana098ab7",
    "policyId": "did8jx6ws45jana098ab7",
    "status": "completed",
    "startedAt": "2020-10-15T06:38:00.000+00:00",
    "migrationId": "did8jx6ws45jana098ab7",
    "services": [],
    "resources": [],
    "options": "{databases.database[{oldId, newId, newName}]}"
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

		_, err := srv.CreateRestoration("<ARCHIVE_ID>", []string{})
		if err != nil {
			t.Errorf("Method CreateRestoration failed: %v", err)
		}
	})

	t.Run("Test ListRestorations", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "restorations": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "archiveId": "did8jx6ws45jana098ab7",
            "policyId": "did8jx6ws45jana098ab7",
            "status": "completed",
            "startedAt": "2020-10-15T06:38:00.000+00:00",
            "migrationId": "did8jx6ws45jana098ab7",
            "services": [],
            "resources": [],
            "options": "{databases.database[{oldId, newId, newName}]}"
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

		_, err := srv.ListRestorations()
		if err != nil {
			t.Errorf("Method ListRestorations failed: %v", err)
		}
	})

	t.Run("Test GetRestoration", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "archiveId": "did8jx6ws45jana098ab7",
    "policyId": "did8jx6ws45jana098ab7",
    "status": "completed",
    "startedAt": "2020-10-15T06:38:00.000+00:00",
    "migrationId": "did8jx6ws45jana098ab7",
    "services": [],
    "resources": [],
    "options": "{databases.database[{oldId, newId, newName}]}"
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

		_, err := srv.GetRestoration("<RESTORATION_ID>")
		if err != nil {
			t.Errorf("Method GetRestoration failed: %v", err)
		}
	})
}
