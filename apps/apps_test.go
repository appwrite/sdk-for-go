package apps

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestApps(t *testing.T) {
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
    "apps": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Application",
            "description": "Connect your workspace to My Application.",
            "clientUri": "https://example.com",
            "logoUri": "https://example.com/logo.png",
            "privacyPolicyUrl": "https://example.com/privacy",
            "termsUrl": "https://example.com/terms",
            "contacts": [],
            "tagline": "Automate your workspace.",
            "tags": [],
            "labels": [],
            "images": [],
            "supportUrl": "https://example.com/support",
            "dataDeletionUrl": "https://example.com/data-deletion",
            "redirectUris": [],
            "postLogoutRedirectUris": [],
            "enabled": true,
            "type": "confidential",
            "deviceFlow": true,
            "teamId": "5e5ea5c16897e",
            "userId": "5e5ea5c16897e",
            "installationScopes": [],
            "installationRedirectUrl": "https://example.com/setup",
            "secrets": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "appId": "5e5ea5c16897e",
                    "secret": "string",
                    "hint": "f5c6c7",
                    "createdById": "5e5ea5c16897e",
                    "createdByName": "Walter White"
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
    "name": "My Application",
    "description": "Connect your workspace to My Application.",
    "clientUri": "https://example.com",
    "logoUri": "https://example.com/logo.png",
    "privacyPolicyUrl": "https://example.com/privacy",
    "termsUrl": "https://example.com/terms",
    "contacts": [],
    "tagline": "Automate your workspace.",
    "tags": [],
    "labels": [],
    "images": [],
    "supportUrl": "https://example.com/support",
    "dataDeletionUrl": "https://example.com/data-deletion",
    "redirectUris": [],
    "postLogoutRedirectUris": [],
    "enabled": true,
    "type": "confidential",
    "deviceFlow": true,
    "teamId": "5e5ea5c16897e",
    "userId": "5e5ea5c16897e",
    "installationScopes": [],
    "installationRedirectUrl": "https://example.com/setup",
    "secrets": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "appId": "5e5ea5c16897e",
            "secret": "string",
            "hint": "f5c6c7",
            "createdById": "5e5ea5c16897e",
            "createdByName": "Walter White"
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

		_, err := srv.Create("<APP_ID>", "<NAME>", []string{})
		if err != nil {
			t.Errorf("Method Create failed: %v", err)
		}
	})

	t.Run("Test ListInstallationScopes", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "scopes": [
        {
            "value": "organization:organization.read",
            "description": "Access to read the organization",
            "type": "organization",
            "category": "Organization",
            "deprecated": true
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

		_, err := srv.ListInstallationScopes()
		if err != nil {
			t.Errorf("Method ListInstallationScopes failed: %v", err)
		}
	})

	t.Run("Test ListOAuth2Scopes", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "scopes": [
        {
            "value": "organization:organization.read",
            "description": "Access to read the organization",
            "type": "organization",
            "category": "Organization",
            "deprecated": true
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

		_, err := srv.ListOAuth2Scopes()
		if err != nil {
			t.Errorf("Method ListOAuth2Scopes failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Application",
    "description": "Connect your workspace to My Application.",
    "clientUri": "https://example.com",
    "logoUri": "https://example.com/logo.png",
    "privacyPolicyUrl": "https://example.com/privacy",
    "termsUrl": "https://example.com/terms",
    "contacts": [],
    "tagline": "Automate your workspace.",
    "tags": [],
    "labels": [],
    "images": [],
    "supportUrl": "https://example.com/support",
    "dataDeletionUrl": "https://example.com/data-deletion",
    "redirectUris": [],
    "postLogoutRedirectUris": [],
    "enabled": true,
    "type": "confidential",
    "deviceFlow": true,
    "teamId": "5e5ea5c16897e",
    "userId": "5e5ea5c16897e",
    "installationScopes": [],
    "installationRedirectUrl": "https://example.com/setup",
    "secrets": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "appId": "5e5ea5c16897e",
            "secret": "string",
            "hint": "f5c6c7",
            "createdById": "5e5ea5c16897e",
            "createdByName": "Walter White"
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

		_, err := srv.Get("<APP_ID>")
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
    "name": "My Application",
    "description": "Connect your workspace to My Application.",
    "clientUri": "https://example.com",
    "logoUri": "https://example.com/logo.png",
    "privacyPolicyUrl": "https://example.com/privacy",
    "termsUrl": "https://example.com/terms",
    "contacts": [],
    "tagline": "Automate your workspace.",
    "tags": [],
    "labels": [],
    "images": [],
    "supportUrl": "https://example.com/support",
    "dataDeletionUrl": "https://example.com/data-deletion",
    "redirectUris": [],
    "postLogoutRedirectUris": [],
    "enabled": true,
    "type": "confidential",
    "deviceFlow": true,
    "teamId": "5e5ea5c16897e",
    "userId": "5e5ea5c16897e",
    "installationScopes": [],
    "installationRedirectUrl": "https://example.com/setup",
    "secrets": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "appId": "5e5ea5c16897e",
            "secret": "string",
            "hint": "f5c6c7",
            "createdById": "5e5ea5c16897e",
            "createdByName": "Walter White"
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

		_, err := srv.Update("<APP_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method Update failed: %v", err)
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

		_, err := srv.Delete("<APP_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})

	t.Run("Test ListInstallations", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "installations": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "appId": "5e5ea5c16897e",
            "teamId": "5e5ea5c16897e",
            "scopes": [],
            "authorizationDetails": {},
            "createdById": "5e5ea5c16897e",
            "createdByName": "Walter White"
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

		_, err := srv.ListInstallations("<APP_ID>")
		if err != nil {
			t.Errorf("Method ListInstallations failed: %v", err)
		}
	})

	t.Run("Test GetInstallation", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "appId": "5e5ea5c16897e",
    "teamId": "5e5ea5c16897e",
    "scopes": [],
    "authorizationDetails": {},
    "createdById": "5e5ea5c16897e",
    "createdByName": "Walter White"
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

		_, err := srv.GetInstallation("<APP_ID>", "<INSTALLATION_ID>")
		if err != nil {
			t.Errorf("Method GetInstallation failed: %v", err)
		}
	})

	t.Run("Test DeleteInstallation", func(t *testing.T) {
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

		_, err := srv.DeleteInstallation("<APP_ID>", "<INSTALLATION_ID>")
		if err != nil {
			t.Errorf("Method DeleteInstallation failed: %v", err)
		}
	})

	t.Run("Test CreateInstallationToken", func(t *testing.T) {
		mockResponse := `
{
    "access_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9...",
    "token_type": "Bearer",
    "expires_in": 3600,
    "refresh_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9...",
    "scope": "openid email profile"
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

		_, err := srv.CreateInstallationToken("<APP_ID>", "<INSTALLATION_ID>")
		if err != nil {
			t.Errorf("Method CreateInstallationToken failed: %v", err)
		}
	})

	t.Run("Test ListKeys", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "keys": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "appId": "5e5ea5c16897e",
            "secret": "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a",
            "hint": "f5c6c7",
            "createdById": "5e5ea5c16897e",
            "createdByName": "Walter White"
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

		_, err := srv.ListKeys("<APP_ID>")
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
    "appId": "5e5ea5c16897e",
    "secret": "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a",
    "hint": "f5c6c7",
    "createdById": "5e5ea5c16897e",
    "createdByName": "Walter White"
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

		_, err := srv.CreateKey("<APP_ID>")
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
    "appId": "5e5ea5c16897e",
    "secret": "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a",
    "hint": "f5c6c7",
    "createdById": "5e5ea5c16897e",
    "createdByName": "Walter White"
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

		_, err := srv.GetKey("<APP_ID>", "<KEY_ID>")
		if err != nil {
			t.Errorf("Method GetKey failed: %v", err)
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

		_, err := srv.DeleteKey("<APP_ID>", "<KEY_ID>")
		if err != nil {
			t.Errorf("Method DeleteKey failed: %v", err)
		}
	})

	t.Run("Test UpdateLabels", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Application",
    "description": "Connect your workspace to My Application.",
    "clientUri": "https://example.com",
    "logoUri": "https://example.com/logo.png",
    "privacyPolicyUrl": "https://example.com/privacy",
    "termsUrl": "https://example.com/terms",
    "contacts": [],
    "tagline": "Automate your workspace.",
    "tags": [],
    "labels": [],
    "images": [],
    "supportUrl": "https://example.com/support",
    "dataDeletionUrl": "https://example.com/data-deletion",
    "redirectUris": [],
    "postLogoutRedirectUris": [],
    "enabled": true,
    "type": "confidential",
    "deviceFlow": true,
    "teamId": "5e5ea5c16897e",
    "userId": "5e5ea5c16897e",
    "installationScopes": [],
    "installationRedirectUrl": "https://example.com/setup",
    "secrets": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "appId": "5e5ea5c16897e",
            "secret": "string",
            "hint": "f5c6c7",
            "createdById": "5e5ea5c16897e",
            "createdByName": "Walter White"
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

		_, err := srv.UpdateLabels("<APP_ID>", []string{})
		if err != nil {
			t.Errorf("Method UpdateLabels failed: %v", err)
		}
	})

	t.Run("Test ListSecrets", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "secrets": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "appId": "5e5ea5c16897e",
            "secret": "string",
            "hint": "f5c6c7",
            "createdById": "5e5ea5c16897e",
            "createdByName": "Walter White"
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

		_, err := srv.ListSecrets("<APP_ID>")
		if err != nil {
			t.Errorf("Method ListSecrets failed: %v", err)
		}
	})

	t.Run("Test CreateSecret", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "appId": "5e5ea5c16897e",
    "secret": "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a",
    "hint": "f5c6c7",
    "createdById": "5e5ea5c16897e",
    "createdByName": "Walter White"
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

		_, err := srv.CreateSecret("<APP_ID>")
		if err != nil {
			t.Errorf("Method CreateSecret failed: %v", err)
		}
	})

	t.Run("Test GetSecret", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "appId": "5e5ea5c16897e",
    "secret": "string",
    "hint": "f5c6c7",
    "createdById": "5e5ea5c16897e",
    "createdByName": "Walter White"
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

		_, err := srv.GetSecret("<APP_ID>", "<SECRET_ID>")
		if err != nil {
			t.Errorf("Method GetSecret failed: %v", err)
		}
	})

	t.Run("Test DeleteSecret", func(t *testing.T) {
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

		_, err := srv.DeleteSecret("<APP_ID>", "<SECRET_ID>")
		if err != nil {
			t.Errorf("Method DeleteSecret failed: %v", err)
		}
	})

	t.Run("Test UpdateTeam", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Application",
    "description": "Connect your workspace to My Application.",
    "clientUri": "https://example.com",
    "logoUri": "https://example.com/logo.png",
    "privacyPolicyUrl": "https://example.com/privacy",
    "termsUrl": "https://example.com/terms",
    "contacts": [],
    "tagline": "Automate your workspace.",
    "tags": [],
    "labels": [],
    "images": [],
    "supportUrl": "https://example.com/support",
    "dataDeletionUrl": "https://example.com/data-deletion",
    "redirectUris": [],
    "postLogoutRedirectUris": [],
    "enabled": true,
    "type": "confidential",
    "deviceFlow": true,
    "teamId": "5e5ea5c16897e",
    "userId": "5e5ea5c16897e",
    "installationScopes": [],
    "installationRedirectUrl": "https://example.com/setup",
    "secrets": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "appId": "5e5ea5c16897e",
            "secret": "string",
            "hint": "f5c6c7",
            "createdById": "5e5ea5c16897e",
            "createdByName": "Walter White"
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

		_, err := srv.UpdateTeam("<APP_ID>", "<TEAM_ID>")
		if err != nil {
			t.Errorf("Method UpdateTeam failed: %v", err)
		}
	})

	t.Run("Test DeleteTokens", func(t *testing.T) {
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

		_, err := srv.DeleteTokens("<APP_ID>")
		if err != nil {
			t.Errorf("Method DeleteTokens failed: %v", err)
		}
	})
}
