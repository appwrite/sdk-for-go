package oauth2

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v6/client"
)

func TestOauth2(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test Approve", func(t *testing.T) {
		mockResponse := `
{
    "redirectUrl": "https://example.com/callback?code=abcde&state=fghij"
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

		_, err := srv.Approve("<GRANT_ID>")
		if err != nil {
			t.Errorf("Method Approve failed: %v", err)
		}
	})

	t.Run("Test Authorize", func(t *testing.T) {
		mockResponse := `
{
    "grantId": "5e5ea5c16897e",
    "redirectUrl": "https://example.com/callback?code=abcde&state=fghij"
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

		_, err := srv.Authorize()
		if err != nil {
			t.Errorf("Method Authorize failed: %v", err)
		}
	})

	t.Run("Test AuthorizePost", func(t *testing.T) {
		mockResponse := `
{
    "grantId": "5e5ea5c16897e",
    "redirectUrl": "https://example.com/callback?code=abcde&state=fghij"
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

		_, err := srv.AuthorizePost()
		if err != nil {
			t.Errorf("Method AuthorizePost failed: %v", err)
		}
	})

	t.Run("Test CreateDeviceAuthorization", func(t *testing.T) {
		mockResponse := `
{
    "device_code": "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a",
    "user_code": "ABCD-EFGH",
    "verification_uri": "https://cloud.appwrite.io/console/oauth2/device",
    "verification_uri_complete": "https://cloud.appwrite.io/console/oauth2/device?user_code=ABCD-EFGH",
    "expires_in": 900,
    "interval": 5
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

		_, err := srv.CreateDeviceAuthorization()
		if err != nil {
			t.Errorf("Method CreateDeviceAuthorization failed: %v", err)
		}
	})

	t.Run("Test CreateGrant", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5ea5c16897e",
    "appId": "5e5ea5c16897e",
    "scopes": [],
    "resources": [],
    "authorizationDetails": "[{\"type\":\"calendar\",\"identifier\":\"primary\",\"actions\":[\"read_events\",\"create_event\"]}]",
    "prompt": "login",
    "redirectUri": "https://example.com/callback",
    "authTime": 1592981250,
    "expire": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateGrant("<USER_CODE>")
		if err != nil {
			t.Errorf("Method CreateGrant failed: %v", err)
		}
	})

	t.Run("Test GetGrant", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5ea5c16897e",
    "appId": "5e5ea5c16897e",
    "scopes": [],
    "resources": [],
    "authorizationDetails": "[{\"type\":\"calendar\",\"identifier\":\"primary\",\"actions\":[\"read_events\",\"create_event\"]}]",
    "prompt": "login",
    "redirectUri": "https://example.com/callback",
    "authTime": 1592981250,
    "expire": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.GetGrant("<GRANT_ID>")
		if err != nil {
			t.Errorf("Method GetGrant failed: %v", err)
		}
	})

	t.Run("Test ListOrganizations", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "organizations": [
        {
            "$id": "5e5ea5c16897e"
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

		_, err := srv.ListOrganizations()
		if err != nil {
			t.Errorf("Method ListOrganizations failed: %v", err)
		}
	})

	t.Run("Test CreatePAR", func(t *testing.T) {
		mockResponse := `
{
    "request_uri": "urn:appwrite:oauth2:request:5e5ea5c16897e",
    "expires_in": 600
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

		_, err := srv.CreatePAR("<CLIENT_ID>", "https://example.com", "code")
		if err != nil {
			t.Errorf("Method CreatePAR failed: %v", err)
		}
	})

	t.Run("Test ListProjects", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "projects": [
        {
            "$id": "5e5ea5c16897e",
            "region": "fra",
            "endpoint": "https://fra.cloud.appwrite.io/v1"
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

	t.Run("Test Reject", func(t *testing.T) {
		mockResponse := `
{
    "redirectUrl": "https://example.com/callback?error=access_denied&state=fghij"
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

		_, err := srv.Reject("<GRANT_ID>")
		if err != nil {
			t.Errorf("Method Reject failed: %v", err)
		}
	})

	t.Run("Test Revoke", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
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

		_, err := srv.Revoke("<TOKEN>")
		if err != nil {
			t.Errorf("Method Revoke failed: %v", err)
		}
	})

	t.Run("Test CreateToken", func(t *testing.T) {
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

		_, err := srv.CreateToken("<GRANT_TYPE>")
		if err != nil {
			t.Errorf("Method CreateToken failed: %v", err)
		}
	})
}
