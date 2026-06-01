package organization

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v5/client"
)

func TestOrganization(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test ListKeys", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "keys": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My API Key",
            "expire": "2020-10-15T06:38:00.000+00:00",
            "scopes": [],
            "secret": "919c2d18fb5d4...a2ae413da83346ad2",
            "accessedAt": "2020-10-15T06:38:00.000+00:00",
            "sdks": []
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

		_, err := srv.ListKeys()
		if err != nil {
			t.Errorf("Method ListKeys failed: %v", err)
		}
	})

	t.Run("Test CreateKey", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My API Key",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "scopes": [],
    "secret": "919c2d18fb5d4...a2ae413da83346ad2",
    "accessedAt": "2020-10-15T06:38:00.000+00:00",
    "sdks": []
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

		_, err := srv.CreateKey("<KEY_ID>", "<NAME>", []string{})
		if err != nil {
			t.Errorf("Method CreateKey failed: %v", err)
		}
	})

	t.Run("Test GetKey", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My API Key",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "scopes": [],
    "secret": "919c2d18fb5d4...a2ae413da83346ad2",
    "accessedAt": "2020-10-15T06:38:00.000+00:00",
    "sdks": []
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

		_, err := srv.GetKey("<KEY_ID>")
		if err != nil {
			t.Errorf("Method GetKey failed: %v", err)
		}
	})

	t.Run("Test UpdateKey", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My API Key",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "scopes": [],
    "secret": "919c2d18fb5d4...a2ae413da83346ad2",
    "accessedAt": "2020-10-15T06:38:00.000+00:00",
    "sdks": []
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

		_, err := srv.UpdateKey("<KEY_ID>", "<NAME>", []string{})
		if err != nil {
			t.Errorf("Method UpdateKey failed: %v", err)
		}
	})

	t.Run("Test DeleteKey", func(t *testing.T) {
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

		_, err := srv.DeleteKey("<KEY_ID>")
		if err != nil {
			t.Errorf("Method DeleteKey failed: %v", err)
		}
	})

	t.Run("Test ListProjects", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "projects": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "New Project",
            "teamId": "1592981250",
            "devKeys": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "name": "Dev API Key",
                    "expire": "2020-10-15T06:38:00.000+00:00",
                    "secret": "919c2d18fb5d4...a2ae413da83346ad2",
                    "accessedAt": "2020-10-15T06:38:00.000+00:00",
                    "sdks": []
                }
            ],
            "smtpEnabled": true,
            "smtpSenderName": "John Appwrite",
            "smtpSenderEmail": "john@appwrite.io",
            "smtpReplyToName": "Support Team",
            "smtpReplyToEmail": "support@appwrite.io",
            "smtpHost": "mail.appwrite.io",
            "smtpPort": 25,
            "smtpUsername": "emailuser",
            "smtpPassword": "string",
            "smtpSecure": "tls",
            "pingCount": 1,
            "pingedAt": "2020-10-15T06:38:00.000+00:00",
            "labels": [],
            "status": "active",
            "authMethods": [
                {
                    "$id": "email-password",
                    "enabled": true
                }
            ],
            "services": [
                {
                    "$id": "sites",
                    "enabled": true
                }
            ],
            "protocols": [
                {
                    "$id": "graphql",
                    "enabled": true
                }
            ],
            "region": "fra",
            "blocks": [
                {
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "resourceType": "project",
                    "resourceId": "5e5ea5c16897e",
                    "projectName": "My Project",
                    "region": "fra",
                    "organizationName": "Acme Inc.",
                    "organizationId": "5e5ea5c16897e",
                    "billingPlan": "pro"
                }
            ],
            "consoleAccessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.ListProjects()
		if err != nil {
			t.Errorf("Method ListProjects failed: %v", err)
		}
	})

	t.Run("Test CreateProject", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "teamId": "1592981250",
    "devKeys": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Dev API Key",
            "expire": "2020-10-15T06:38:00.000+00:00",
            "secret": "919c2d18fb5d4...a2ae413da83346ad2",
            "accessedAt": "2020-10-15T06:38:00.000+00:00",
            "sdks": []
        }
    ],
    "smtpEnabled": true,
    "smtpSenderName": "John Appwrite",
    "smtpSenderEmail": "john@appwrite.io",
    "smtpReplyToName": "Support Team",
    "smtpReplyToEmail": "support@appwrite.io",
    "smtpHost": "mail.appwrite.io",
    "smtpPort": 25,
    "smtpUsername": "emailuser",
    "smtpPassword": "string",
    "smtpSecure": "tls",
    "pingCount": 1,
    "pingedAt": "2020-10-15T06:38:00.000+00:00",
    "labels": [],
    "status": "active",
    "authMethods": [
        {
            "$id": "email-password",
            "enabled": true
        }
    ],
    "services": [
        {
            "$id": "sites",
            "enabled": true
        }
    ],
    "protocols": [
        {
            "$id": "graphql",
            "enabled": true
        }
    ],
    "region": "fra",
    "blocks": [
        {
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "resourceType": "project",
            "resourceId": "5e5ea5c16897e",
            "projectName": "My Project",
            "region": "fra",
            "organizationName": "Acme Inc.",
            "organizationId": "5e5ea5c16897e",
            "billingPlan": "pro"
        }
    ],
    "consoleAccessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateProject("", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateProject failed: %v", err)
		}
	})

	t.Run("Test GetProject", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "teamId": "1592981250",
    "devKeys": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Dev API Key",
            "expire": "2020-10-15T06:38:00.000+00:00",
            "secret": "919c2d18fb5d4...a2ae413da83346ad2",
            "accessedAt": "2020-10-15T06:38:00.000+00:00",
            "sdks": []
        }
    ],
    "smtpEnabled": true,
    "smtpSenderName": "John Appwrite",
    "smtpSenderEmail": "john@appwrite.io",
    "smtpReplyToName": "Support Team",
    "smtpReplyToEmail": "support@appwrite.io",
    "smtpHost": "mail.appwrite.io",
    "smtpPort": 25,
    "smtpUsername": "emailuser",
    "smtpPassword": "string",
    "smtpSecure": "tls",
    "pingCount": 1,
    "pingedAt": "2020-10-15T06:38:00.000+00:00",
    "labels": [],
    "status": "active",
    "authMethods": [
        {
            "$id": "email-password",
            "enabled": true
        }
    ],
    "services": [
        {
            "$id": "sites",
            "enabled": true
        }
    ],
    "protocols": [
        {
            "$id": "graphql",
            "enabled": true
        }
    ],
    "region": "fra",
    "blocks": [
        {
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "resourceType": "project",
            "resourceId": "5e5ea5c16897e",
            "projectName": "My Project",
            "region": "fra",
            "organizationName": "Acme Inc.",
            "organizationId": "5e5ea5c16897e",
            "billingPlan": "pro"
        }
    ],
    "consoleAccessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.GetProject("<PROJECT_ID>")
		if err != nil {
			t.Errorf("Method GetProject failed: %v", err)
		}
	})

	t.Run("Test UpdateProject", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "teamId": "1592981250",
    "devKeys": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Dev API Key",
            "expire": "2020-10-15T06:38:00.000+00:00",
            "secret": "919c2d18fb5d4...a2ae413da83346ad2",
            "accessedAt": "2020-10-15T06:38:00.000+00:00",
            "sdks": []
        }
    ],
    "smtpEnabled": true,
    "smtpSenderName": "John Appwrite",
    "smtpSenderEmail": "john@appwrite.io",
    "smtpReplyToName": "Support Team",
    "smtpReplyToEmail": "support@appwrite.io",
    "smtpHost": "mail.appwrite.io",
    "smtpPort": 25,
    "smtpUsername": "emailuser",
    "smtpPassword": "string",
    "smtpSecure": "tls",
    "pingCount": 1,
    "pingedAt": "2020-10-15T06:38:00.000+00:00",
    "labels": [],
    "status": "active",
    "authMethods": [
        {
            "$id": "email-password",
            "enabled": true
        }
    ],
    "services": [
        {
            "$id": "sites",
            "enabled": true
        }
    ],
    "protocols": [
        {
            "$id": "graphql",
            "enabled": true
        }
    ],
    "region": "fra",
    "blocks": [
        {
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "resourceType": "project",
            "resourceId": "5e5ea5c16897e",
            "projectName": "My Project",
            "region": "fra",
            "organizationName": "Acme Inc.",
            "organizationId": "5e5ea5c16897e",
            "billingPlan": "pro"
        }
    ],
    "consoleAccessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateProject("<PROJECT_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method UpdateProject failed: %v", err)
		}
	})

	t.Run("Test DeleteProject", func(t *testing.T) {
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

		_, err := srv.DeleteProject("<PROJECT_ID>")
		if err != nil {
			t.Errorf("Method DeleteProject failed: %v", err)
		}
	})
}
