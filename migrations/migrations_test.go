package migrations

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestMigrations(t *testing.T) {
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
    "migrations": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "status": "pending",
            "stage": "init",
            "source": "Appwrite",
            "destination": "Appwrite",
            "resources": [],
            "resourceId": "collectionId",
            "resourceInternalId": "1",
            "resourceType": "collection",
            "parentResourceId": "databaseId",
            "parentResourceInternalId": "1",
            "parentResourceType": "database",
            "destinationResourceId": "databaseId",
            "destinationResourceInternalId": "1",
            "destinationResourceType": "database",
            "statusCounters": {},
            "resourceData": [],
            "errors": [],
            "options": {}
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

	t.Run("Test CreateAppwriteMigration", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.CreateAppwriteMigration([]string{}, "https://example.com", "<PROJECT_ID>", "<API_KEY>")
		if err != nil {
			t.Errorf("Method CreateAppwriteMigration failed: %v", err)
		}
	})

	t.Run("Test CreateCSVExport", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.CreateCSVExport("<DATABASE_ID>", "<COLLECTION_ID>", "<FILENAME>")
		if err != nil {
			t.Errorf("Method CreateCSVExport failed: %v", err)
		}
	})

	t.Run("Test CreateCSVImport", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.CreateCSVImport("<BUCKET_ID>", "<FILE_ID>", "<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method CreateCSVImport failed: %v", err)
		}
	})

	t.Run("Test CreateFirebaseMigration", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.CreateFirebaseMigration([]string{}, "<SERVICE_ACCOUNT>")
		if err != nil {
			t.Errorf("Method CreateFirebaseMigration failed: %v", err)
		}
	})

	t.Run("Test GetFirebaseReport", func(t *testing.T) {
		mockResponse := `
{
    "user": 20,
    "team": 20,
    "database": 20,
    "row": 20,
    "file": 20,
    "bucket": 20,
    "function": 20,
    "platform": 5,
    "api-key": 5,
    "project-variable": 5,
    "webhook": 5,
    "auth-methods": 1,
    "project-protocols": 1,
    "project-labels": 1,
    "project-services": 1,
    "policies": 1,
    "smtp": 1,
    "rule": 5,
    "project-email-template": 7,
    "project-environment": 2,
    "site": 5,
    "provider": 5,
    "topic": 10,
    "subscriber": 100,
    "message": 50,
    "size": 30000,
    "version": "1.4.0",
    "oauth2-provider": 5,
    "backup-policy": 5
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

		_, err := srv.GetFirebaseReport([]string{}, "<SERVICE_ACCOUNT>")
		if err != nil {
			t.Errorf("Method GetFirebaseReport failed: %v", err)
		}
	})

	t.Run("Test CreateJSONExport", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.CreateJSONExport("<DATABASE_ID>", "<COLLECTION_ID>", "<FILENAME>")
		if err != nil {
			t.Errorf("Method CreateJSONExport failed: %v", err)
		}
	})

	t.Run("Test CreateJSONImport", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.CreateJSONImport("<BUCKET_ID>", "<FILE_ID>", "<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method CreateJSONImport failed: %v", err)
		}
	})

	t.Run("Test CreateNHostMigration", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.CreateNHostMigration([]string{}, "<SUBDOMAIN>", "<REGION>", "<ADMIN_SECRET>", "<DATABASE>", "<USERNAME>", "password")
		if err != nil {
			t.Errorf("Method CreateNHostMigration failed: %v", err)
		}
	})

	t.Run("Test GetNHostReport", func(t *testing.T) {
		mockResponse := `
{
    "user": 20,
    "team": 20,
    "database": 20,
    "row": 20,
    "file": 20,
    "bucket": 20,
    "function": 20,
    "platform": 5,
    "api-key": 5,
    "project-variable": 5,
    "webhook": 5,
    "auth-methods": 1,
    "project-protocols": 1,
    "project-labels": 1,
    "project-services": 1,
    "policies": 1,
    "smtp": 1,
    "rule": 5,
    "project-email-template": 7,
    "project-environment": 2,
    "site": 5,
    "provider": 5,
    "topic": 10,
    "subscriber": 100,
    "message": 50,
    "size": 30000,
    "version": "1.4.0",
    "oauth2-provider": 5,
    "backup-policy": 5
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

		_, err := srv.GetNHostReport([]string{}, "<SUBDOMAIN>", "<REGION>", "<ADMIN_SECRET>", "<DATABASE>", "<USERNAME>", "password")
		if err != nil {
			t.Errorf("Method GetNHostReport failed: %v", err)
		}
	})

	t.Run("Test CreateSupabaseMigration", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.CreateSupabaseMigration([]string{}, "https://example.com", "<API_KEY>", "<DATABASE_HOST>", "<USERNAME>", "password")
		if err != nil {
			t.Errorf("Method CreateSupabaseMigration failed: %v", err)
		}
	})

	t.Run("Test GetSupabaseReport", func(t *testing.T) {
		mockResponse := `
{
    "user": 20,
    "team": 20,
    "database": 20,
    "row": 20,
    "file": 20,
    "bucket": 20,
    "function": 20,
    "platform": 5,
    "api-key": 5,
    "project-variable": 5,
    "webhook": 5,
    "auth-methods": 1,
    "project-protocols": 1,
    "project-labels": 1,
    "project-services": 1,
    "policies": 1,
    "smtp": 1,
    "rule": 5,
    "project-email-template": 7,
    "project-environment": 2,
    "site": 5,
    "provider": 5,
    "topic": 10,
    "subscriber": 100,
    "message": 50,
    "size": 30000,
    "version": "1.4.0",
    "oauth2-provider": 5,
    "backup-policy": 5
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

		_, err := srv.GetSupabaseReport([]string{}, "https://example.com", "<API_KEY>", "<DATABASE_HOST>", "<USERNAME>", "password")
		if err != nil {
			t.Errorf("Method GetSupabaseReport failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.Get("<MIGRATION_ID>")
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
		}
	})

	t.Run("Test Retry", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "stage": "init",
    "source": "Appwrite",
    "destination": "Appwrite",
    "resources": [],
    "resourceId": "collectionId",
    "resourceInternalId": "1",
    "resourceType": "collection",
    "parentResourceId": "databaseId",
    "parentResourceInternalId": "1",
    "parentResourceType": "database",
    "destinationResourceId": "databaseId",
    "destinationResourceInternalId": "1",
    "destinationResourceType": "database",
    "statusCounters": {},
    "resourceData": [],
    "errors": [],
    "options": {}
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

		_, err := srv.Retry("<MIGRATION_ID>")
		if err != nil {
			t.Errorf("Method Retry failed: %v", err)
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

		_, err := srv.Delete("<MIGRATION_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})
}
