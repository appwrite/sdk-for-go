package monitors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestMonitors(t *testing.T) {
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
    "monitors": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Bandwidth spike",
            "projectId": "5e5ea5c16897e",
            "rules": [],
            "type": "usage",
            "resourceType": "database",
            "resourceId": "main",
            "window": 3600,
            "cooldown": 3600,
            "enabled": true,
            "notify": [],
            "escalationId": "5e5ea5c16897e",
            "severity": "warning",
            "status": "ok",
            "alertId": "5e5ea5c16897e",
            "interval": 60,
            "timeout": 10,
            "url": "https://example.com/health",
            "body": "string",
            "domain": "example.com",
            "resolver": "1.1.1.1",
            "host": "mail.example.com",
            "port": 587
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

	t.Run("Test CreateDnsMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.CreateDnsMonitor("<MONITOR_ID>", "<NAME>", "example.com")
		if err != nil {
			t.Errorf("Method CreateDnsMonitor failed: %v", err)
		}
	})

	t.Run("Test UpdateDnsMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.UpdateDnsMonitor("<MONITOR_ID>")
		if err != nil {
			t.Errorf("Method UpdateDnsMonitor failed: %v", err)
		}
	})

	t.Run("Test CreateHttpMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.CreateHttpMonitor("<MONITOR_ID>", "<NAME>", "https://api.example.com/health")
		if err != nil {
			t.Errorf("Method CreateHttpMonitor failed: %v", err)
		}
	})

	t.Run("Test UpdateHttpMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.UpdateHttpMonitor("<MONITOR_ID>")
		if err != nil {
			t.Errorf("Method UpdateHttpMonitor failed: %v", err)
		}
	})

	t.Run("Test CreateSmtpMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.CreateSmtpMonitor("<MONITOR_ID>", "<NAME>", "smtp.example.com", 1)
		if err != nil {
			t.Errorf("Method CreateSmtpMonitor failed: %v", err)
		}
	})

	t.Run("Test UpdateSmtpMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.UpdateSmtpMonitor("<MONITOR_ID>")
		if err != nil {
			t.Errorf("Method UpdateSmtpMonitor failed: %v", err)
		}
	})

	t.Run("Test CreateTcpMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.CreateTcpMonitor("<MONITOR_ID>", "<NAME>", "db.example.com", 1)
		if err != nil {
			t.Errorf("Method CreateTcpMonitor failed: %v", err)
		}
	})

	t.Run("Test UpdateTcpMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.UpdateTcpMonitor("<MONITOR_ID>")
		if err != nil {
			t.Errorf("Method UpdateTcpMonitor failed: %v", err)
		}
	})

	t.Run("Test CreateUsageMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.CreateUsageMonitor("<MONITOR_ID>", "<NAME>", []string{})
		if err != nil {
			t.Errorf("Method CreateUsageMonitor failed: %v", err)
		}
	})

	t.Run("Test UpdateUsageMonitor", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.UpdateUsageMonitor("<MONITOR_ID>")
		if err != nil {
			t.Errorf("Method UpdateUsageMonitor failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Bandwidth spike",
    "projectId": "5e5ea5c16897e",
    "rules": [],
    "type": "usage",
    "resourceType": "database",
    "resourceId": "main",
    "window": 3600,
    "cooldown": 3600,
    "enabled": true,
    "notify": [],
    "escalationId": "5e5ea5c16897e",
    "severity": "warning",
    "status": "ok",
    "alertId": "5e5ea5c16897e",
    "interval": 60,
    "timeout": 10,
    "url": "https://example.com/health",
    "body": "string",
    "domain": "example.com",
    "resolver": "1.1.1.1",
    "host": "mail.example.com",
    "port": 587
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

		_, err := srv.Get("<MONITOR_ID>")
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
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

		_, err := srv.Delete("<MONITOR_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})
}
