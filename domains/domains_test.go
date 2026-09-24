package domains

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestDomains(t *testing.T) {
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
    "domains": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "domain": "example.com",
            "registrar": "appwrite",
            "nameservers": "Appwrite",
            "expire": "2020-10-15T06:38:00.000+00:00",
            "renewal": "2020-10-15T06:38:00.000+00:00",
            "autoRenewal": true,
            "renewalPrice": 2599,
            "teamId": "5e5ea5c16897e",
            "dnsRecords": [
                {
                    "$id": "5f40a6e10c65e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "type": "A",
                    "name": "mail",
                    "value": "192.0.2.1",
                    "ttl": 86400,
                    "priority": 10,
                    "lock": true,
                    "weight": 10,
                    "port": 443,
                    "comment": "Mail server record"
                }
            ]
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

	t.Run("Test Create", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "example.com",
    "registrar": "appwrite",
    "nameservers": "Appwrite",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "renewal": "2020-10-15T06:38:00.000+00:00",
    "autoRenewal": true,
    "renewalPrice": 2599,
    "teamId": "5e5ea5c16897e",
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.Create("<TEAM_ID>", "example.com")
		if err != nil {
			t.Errorf("Method Create failed: %v", err)
		}
	})

	t.Run("Test GetPrice", func(t *testing.T) {
		mockResponse := `
{
    "domain": "example.com",
    "tld": "com",
    "available": true,
    "periodYears": 1,
    "premium": true,
    "renewalPeriodYears": 1
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

		_, err := srv.GetPrice("example.com")
		if err != nil {
			t.Errorf("Method GetPrice failed: %v", err)
		}
	})

	t.Run("Test ListPrices", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "prices": [
        {
            "domain": "example.com",
            "tld": "com",
            "available": true,
            "periodYears": 1,
            "premium": true,
            "renewalPeriodYears": 1
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

		_, err := srv.ListPrices([]string{})
		if err != nil {
			t.Errorf("Method ListPrices failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "example.com",
    "registrar": "appwrite",
    "nameservers": "Appwrite",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "renewal": "2020-10-15T06:38:00.000+00:00",
    "autoRenewal": true,
    "renewalPrice": 2599,
    "teamId": "5e5ea5c16897e",
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
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

		_, err := srv.Get("<DOMAIN_ID>")
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

		_, err := srv.Delete("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})

	t.Run("Test UpdateNameservers", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "example.com",
    "registrar": "appwrite",
    "nameservers": "Appwrite",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "renewal": "2020-10-15T06:38:00.000+00:00",
    "autoRenewal": true,
    "renewalPrice": 2599,
    "teamId": "5e5ea5c16897e",
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.UpdateNameservers("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method UpdateNameservers failed: %v", err)
		}
	})

	t.Run("Test VerifyNameservers", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "example.com",
    "registrar": "appwrite",
    "nameservers": "Appwrite",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "renewal": "2020-10-15T06:38:00.000+00:00",
    "autoRenewal": true,
    "renewalPrice": 2599,
    "teamId": "5e5ea5c16897e",
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.VerifyNameservers("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method VerifyNameservers failed: %v", err)
		}
	})

	t.Run("Test GetPresetGoogleWorkspace", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
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

		_, err := srv.GetPresetGoogleWorkspace("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method GetPresetGoogleWorkspace failed: %v", err)
		}
	})

	t.Run("Test CreatePresetGoogleWorkspace", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.CreatePresetGoogleWorkspace("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method CreatePresetGoogleWorkspace failed: %v", err)
		}
	})

	t.Run("Test GetPresetICloud", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
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

		_, err := srv.GetPresetICloud("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method GetPresetICloud failed: %v", err)
		}
	})

	t.Run("Test CreatePresetICloud", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.CreatePresetICloud("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method CreatePresetICloud failed: %v", err)
		}
	})

	t.Run("Test GetPresetMailgun", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
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

		_, err := srv.GetPresetMailgun("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method GetPresetMailgun failed: %v", err)
		}
	})

	t.Run("Test CreatePresetMailgun", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.CreatePresetMailgun("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method CreatePresetMailgun failed: %v", err)
		}
	})

	t.Run("Test GetPresetOutlook", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
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

		_, err := srv.GetPresetOutlook("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method GetPresetOutlook failed: %v", err)
		}
	})

	t.Run("Test CreatePresetOutlook", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.CreatePresetOutlook("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method CreatePresetOutlook failed: %v", err)
		}
	})

	t.Run("Test GetPresetProtonMail", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
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

		_, err := srv.GetPresetProtonMail("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method GetPresetProtonMail failed: %v", err)
		}
	})

	t.Run("Test CreatePresetProtonMail", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.CreatePresetProtonMail("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method CreatePresetProtonMail failed: %v", err)
		}
	})

	t.Run("Test GetPresetZoho", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
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

		_, err := srv.GetPresetZoho("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method GetPresetZoho failed: %v", err)
		}
	})

	t.Run("Test CreatePresetZoho", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.CreatePresetZoho("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method CreatePresetZoho failed: %v", err)
		}
	})

	t.Run("Test ListRecords", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
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

		_, err := srv.ListRecords("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method ListRecords failed: %v", err)
		}
	})

	t.Run("Test CreateRecordA", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordA("<DOMAIN_ID>", "", "", 1)
		if err != nil {
			t.Errorf("Method CreateRecordA failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordA", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordA("<DOMAIN_ID>", "<RECORD_ID>", "", "", 1)
		if err != nil {
			t.Errorf("Method UpdateRecordA failed: %v", err)
		}
	})

	t.Run("Test CreateRecordAAAA", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordAAAA("<DOMAIN_ID>", "", "", 1)
		if err != nil {
			t.Errorf("Method CreateRecordAAAA failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordAAAA", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordAAAA("<DOMAIN_ID>", "<RECORD_ID>", "", "", 1)
		if err != nil {
			t.Errorf("Method UpdateRecordAAAA failed: %v", err)
		}
	})

	t.Run("Test CreateRecordAlias", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordAlias("<DOMAIN_ID>", "", "<VALUE>", 1)
		if err != nil {
			t.Errorf("Method CreateRecordAlias failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordAlias", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordAlias("<DOMAIN_ID>", "<RECORD_ID>", "", "<VALUE>", 1)
		if err != nil {
			t.Errorf("Method UpdateRecordAlias failed: %v", err)
		}
	})

	t.Run("Test CreateRecordCAA", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordCAA("<DOMAIN_ID>", "", "", 1)
		if err != nil {
			t.Errorf("Method CreateRecordCAA failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordCAA", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordCAA("<DOMAIN_ID>", "<RECORD_ID>", "", "", 1)
		if err != nil {
			t.Errorf("Method UpdateRecordCAA failed: %v", err)
		}
	})

	t.Run("Test CreateRecordCNAME", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordCNAME("<DOMAIN_ID>", "", "<VALUE>", 1)
		if err != nil {
			t.Errorf("Method CreateRecordCNAME failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordCNAME", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordCNAME("<DOMAIN_ID>", "<RECORD_ID>", "", "<VALUE>", 1)
		if err != nil {
			t.Errorf("Method UpdateRecordCNAME failed: %v", err)
		}
	})

	t.Run("Test CreateRecordHTTPS", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordHTTPS("<DOMAIN_ID>", "", "<VALUE>", 1)
		if err != nil {
			t.Errorf("Method CreateRecordHTTPS failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordHTTPS", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordHTTPS("<DOMAIN_ID>", "<RECORD_ID>", "", "<VALUE>", 1)
		if err != nil {
			t.Errorf("Method UpdateRecordHTTPS failed: %v", err)
		}
	})

	t.Run("Test CreateRecordMX", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordMX("<DOMAIN_ID>", "", "<VALUE>", 1, 1)
		if err != nil {
			t.Errorf("Method CreateRecordMX failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordMX", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordMX("<DOMAIN_ID>", "<RECORD_ID>", "", "<VALUE>", 1, 1)
		if err != nil {
			t.Errorf("Method UpdateRecordMX failed: %v", err)
		}
	})

	t.Run("Test CreateRecordNS", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordNS("<DOMAIN_ID>", "", "<VALUE>", 1)
		if err != nil {
			t.Errorf("Method CreateRecordNS failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordNS", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordNS("<DOMAIN_ID>", "<RECORD_ID>", "", "<VALUE>", 1)
		if err != nil {
			t.Errorf("Method UpdateRecordNS failed: %v", err)
		}
	})

	t.Run("Test CreateRecordSRV", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordSRV("<DOMAIN_ID>", "", "<VALUE>", 1, 1, 1, 1)
		if err != nil {
			t.Errorf("Method CreateRecordSRV failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordSRV", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordSRV("<DOMAIN_ID>", "<RECORD_ID>", "", "<VALUE>", 1, 1, 1, 1)
		if err != nil {
			t.Errorf("Method UpdateRecordSRV failed: %v", err)
		}
	})

	t.Run("Test CreateRecordTXT", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.CreateRecordTXT("<DOMAIN_ID>", "", 1)
		if err != nil {
			t.Errorf("Method CreateRecordTXT failed: %v", err)
		}
	})

	t.Run("Test UpdateRecordTXT", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecordTXT("<DOMAIN_ID>", "<RECORD_ID>", "", "<VALUE>", 1)
		if err != nil {
			t.Errorf("Method UpdateRecordTXT failed: %v", err)
		}
	})

	t.Run("Test GetRecord", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5f40a6e10c65e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "type": "A",
    "name": "mail",
    "value": "192.0.2.1",
    "ttl": 86400,
    "priority": 10,
    "lock": true,
    "weight": 10,
    "port": 443,
    "comment": "Mail server record"
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

		_, err := srv.GetRecord("<DOMAIN_ID>", "<RECORD_ID>")
		if err != nil {
			t.Errorf("Method GetRecord failed: %v", err)
		}
	})

	t.Run("Test DeleteRecord", func(t *testing.T) {
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

		_, err := srv.DeleteRecord("<DOMAIN_ID>", "<RECORD_ID>")
		if err != nil {
			t.Errorf("Method DeleteRecord failed: %v", err)
		}
	})

	t.Run("Test UpdateTeam", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "example.com",
    "registrar": "appwrite",
    "nameservers": "Appwrite",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "renewal": "2020-10-15T06:38:00.000+00:00",
    "autoRenewal": true,
    "renewalPrice": 2599,
    "teamId": "5e5ea5c16897e",
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
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

		_, err := srv.UpdateTeam("<DOMAIN_ID>", "<TEAM_ID>")
		if err != nil {
			t.Errorf("Method UpdateTeam failed: %v", err)
		}
	})

	t.Run("Test GetTransferStatus", func(t *testing.T) {
		mockResponse := `
{
    "status": "pending_registry",
    "reason": "Transfer in progress",
    "timestamp": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.GetTransferStatus("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method GetTransferStatus failed: %v", err)
		}
	})

	t.Run("Test GetZone", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
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

		_, err := srv.GetZone("<DOMAIN_ID>")
		if err != nil {
			t.Errorf("Method GetZone failed: %v", err)
		}
	})

	t.Run("Test UpdateZone", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "example.com",
    "registrar": "appwrite",
    "nameservers": "Appwrite",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "renewal": "2020-10-15T06:38:00.000+00:00",
    "autoRenewal": true,
    "renewalPrice": 2599,
    "teamId": "5e5ea5c16897e",
    "dnsRecords": [
        {
            "$id": "5f40a6e10c65e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "type": "A",
            "name": "mail",
            "value": "192.0.2.1",
            "ttl": 86400,
            "priority": 10,
            "lock": true,
            "weight": 10,
            "port": 443,
            "comment": "Mail server record"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateZone("<DOMAIN_ID>", "<CONTENT>")
		if err != nil {
			t.Errorf("Method UpdateZone failed: %v", err)
		}
	})
}
