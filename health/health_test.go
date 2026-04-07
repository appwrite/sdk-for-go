package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v2/client"
)

func TestHealth(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "name": "database",
    "ping": 128,
    "status": "pass"
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

		_, err := srv.Get()
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
		}
	})

	t.Run("Test GetAntivirus", func(t *testing.T) {
		mockResponse := `
{
    "version": "1.0.0",
    "status": "online"
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

		_, err := srv.GetAntivirus()
		if err != nil {
			t.Errorf("Method GetAntivirus failed: %v", err)
		}
	})

	t.Run("Test GetCache", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "statuses": [
        {
            "name": "database",
            "ping": 128,
            "status": "pass"
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

		_, err := srv.GetCache()
		if err != nil {
			t.Errorf("Method GetCache failed: %v", err)
		}
	})

	t.Run("Test GetCertificate", func(t *testing.T) {
		mockResponse := `
{
    "name": "/CN=www.google.com",
    "subjectSN": "string",
    "issuerOrganisation": "string",
    "validFrom": "1704200998",
    "validTo": "1711458597",
    "signatureTypeSN": "RSA-SHA256"
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

		_, err := srv.GetCertificate()
		if err != nil {
			t.Errorf("Method GetCertificate failed: %v", err)
		}
	})

	t.Run("Test GetConsolePausing", func(t *testing.T) {
		mockResponse := `
{
    "name": "database",
    "ping": 128,
    "status": "pass"
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

		_, err := srv.GetConsolePausing()
		if err != nil {
			t.Errorf("Method GetConsolePausing failed: %v", err)
		}
	})

	t.Run("Test GetDB", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "statuses": [
        {
            "name": "database",
            "ping": 128,
            "status": "pass"
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

		_, err := srv.GetDB()
		if err != nil {
			t.Errorf("Method GetDB failed: %v", err)
		}
	})

	t.Run("Test GetPubSub", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "statuses": [
        {
            "name": "database",
            "ping": 128,
            "status": "pass"
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

		_, err := srv.GetPubSub()
		if err != nil {
			t.Errorf("Method GetPubSub failed: %v", err)
		}
	})

	t.Run("Test GetQueueAudits", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueAudits()
		if err != nil {
			t.Errorf("Method GetQueueAudits failed: %v", err)
		}
	})

	t.Run("Test GetQueueBuilds", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueBuilds()
		if err != nil {
			t.Errorf("Method GetQueueBuilds failed: %v", err)
		}
	})

	t.Run("Test GetQueueCertificates", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueCertificates()
		if err != nil {
			t.Errorf("Method GetQueueCertificates failed: %v", err)
		}
	})

	t.Run("Test GetQueueDatabases", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueDatabases()
		if err != nil {
			t.Errorf("Method GetQueueDatabases failed: %v", err)
		}
	})

	t.Run("Test GetQueueDeletes", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueDeletes()
		if err != nil {
			t.Errorf("Method GetQueueDeletes failed: %v", err)
		}
	})

	t.Run("Test GetFailedJobs", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetFailedJobs("v1-database")
		if err != nil {
			t.Errorf("Method GetFailedJobs failed: %v", err)
		}
	})

	t.Run("Test GetQueueFunctions", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueFunctions()
		if err != nil {
			t.Errorf("Method GetQueueFunctions failed: %v", err)
		}
	})

	t.Run("Test GetQueueLogs", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueLogs()
		if err != nil {
			t.Errorf("Method GetQueueLogs failed: %v", err)
		}
	})

	t.Run("Test GetQueueMails", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueMails()
		if err != nil {
			t.Errorf("Method GetQueueMails failed: %v", err)
		}
	})

	t.Run("Test GetQueueMessaging", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueMessaging()
		if err != nil {
			t.Errorf("Method GetQueueMessaging failed: %v", err)
		}
	})

	t.Run("Test GetQueueMigrations", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueMigrations()
		if err != nil {
			t.Errorf("Method GetQueueMigrations failed: %v", err)
		}
	})

	t.Run("Test GetQueueStatsResources", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueStatsResources()
		if err != nil {
			t.Errorf("Method GetQueueStatsResources failed: %v", err)
		}
	})

	t.Run("Test GetQueueUsage", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueUsage()
		if err != nil {
			t.Errorf("Method GetQueueUsage failed: %v", err)
		}
	})

	t.Run("Test GetQueueWebhooks", func(t *testing.T) {
		mockResponse := `
{
    "size": 8
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

		_, err := srv.GetQueueWebhooks()
		if err != nil {
			t.Errorf("Method GetQueueWebhooks failed: %v", err)
		}
	})

	t.Run("Test GetStorage", func(t *testing.T) {
		mockResponse := `
{
    "name": "database",
    "ping": 128,
    "status": "pass"
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

		_, err := srv.GetStorage()
		if err != nil {
			t.Errorf("Method GetStorage failed: %v", err)
		}
	})

	t.Run("Test GetStorageLocal", func(t *testing.T) {
		mockResponse := `
{
    "name": "database",
    "ping": 128,
    "status": "pass"
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

		_, err := srv.GetStorageLocal()
		if err != nil {
			t.Errorf("Method GetStorageLocal failed: %v", err)
		}
	})

	t.Run("Test GetTime", func(t *testing.T) {
		mockResponse := `
{
    "remoteTime": 1639490751,
    "localTime": 1639490844,
    "diff": 93
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

		_, err := srv.GetTime()
		if err != nil {
			t.Errorf("Method GetTime failed: %v", err)
		}
	})
}
