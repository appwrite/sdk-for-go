package teams

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestTeams(t *testing.T) {
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
    "teams": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "VIP",
            "total": 7,
            "prefs": {
            }
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
    "name": "VIP",
    "total": 7,
    "prefs": {
    }
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

		_, err := srv.Create("<TEAM_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method Create failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "VIP",
    "total": 7,
    "prefs": {
    }
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

		_, err := srv.Get("<TEAM_ID>")
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
		}
	})

	t.Run("Test UpdateName", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "VIP",
    "total": 7,
    "prefs": {
    }
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

		_, err := srv.UpdateName("<TEAM_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method UpdateName failed: %v", err)
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

		_, err := srv.Delete("<TEAM_ID>")
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
            "authorizationDetails": [],
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

		_, err := srv.ListInstallations("<TEAM_ID>")
		if err != nil {
			t.Errorf("Method ListInstallations failed: %v", err)
		}
	})

	t.Run("Test CreateInstallation", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "appId": "5e5ea5c16897e",
    "teamId": "5e5ea5c16897e",
    "scopes": [],
    "authorizationDetails": [],
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

		_, err := srv.CreateInstallation("<TEAM_ID>", "<APP_ID>")
		if err != nil {
			t.Errorf("Method CreateInstallation failed: %v", err)
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
    "authorizationDetails": [],
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

		_, err := srv.GetInstallation("<TEAM_ID>", "<INSTALLATION_ID>")
		if err != nil {
			t.Errorf("Method GetInstallation failed: %v", err)
		}
	})

	t.Run("Test UpdateInstallation", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "appId": "5e5ea5c16897e",
    "teamId": "5e5ea5c16897e",
    "scopes": [],
    "authorizationDetails": [],
    "createdById": "5e5ea5c16897e",
    "createdByName": "Walter White"
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

		_, err := srv.UpdateInstallation("<TEAM_ID>", "<INSTALLATION_ID>")
		if err != nil {
			t.Errorf("Method UpdateInstallation failed: %v", err)
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

		_, err := srv.DeleteInstallation("<TEAM_ID>", "<INSTALLATION_ID>")
		if err != nil {
			t.Errorf("Method DeleteInstallation failed: %v", err)
		}
	})

	t.Run("Test ListMemberships", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "memberships": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "userId": "5e5ea5c16897e",
            "userName": "John Doe",
            "userEmail": "john@appwrite.io",
            "userPhone": "+1 555 555 5555",
            "teamId": "5e5ea5c16897e",
            "teamName": "VIP",
            "invited": "2020-10-15T06:38:00.000+00:00",
            "joined": "2020-10-15T06:38:00.000+00:00",
            "confirm": true,
            "mfa": true,
            "userAccessedAt": "2020-10-15T06:38:00.000+00:00",
            "roles": []
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

		_, err := srv.ListMemberships("<TEAM_ID>")
		if err != nil {
			t.Errorf("Method ListMemberships failed: %v", err)
		}
	})

	t.Run("Test CreateMembership", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5ea5c16897e",
    "userName": "John Doe",
    "userEmail": "john@appwrite.io",
    "userPhone": "+1 555 555 5555",
    "teamId": "5e5ea5c16897e",
    "teamName": "VIP",
    "invited": "2020-10-15T06:38:00.000+00:00",
    "joined": "2020-10-15T06:38:00.000+00:00",
    "confirm": true,
    "mfa": true,
    "userAccessedAt": "2020-10-15T06:38:00.000+00:00",
    "roles": []
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

		_, err := srv.CreateMembership("<TEAM_ID>", []string{})
		if err != nil {
			t.Errorf("Method CreateMembership failed: %v", err)
		}
	})

	t.Run("Test GetMembership", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5ea5c16897e",
    "userName": "John Doe",
    "userEmail": "john@appwrite.io",
    "userPhone": "+1 555 555 5555",
    "teamId": "5e5ea5c16897e",
    "teamName": "VIP",
    "invited": "2020-10-15T06:38:00.000+00:00",
    "joined": "2020-10-15T06:38:00.000+00:00",
    "confirm": true,
    "mfa": true,
    "userAccessedAt": "2020-10-15T06:38:00.000+00:00",
    "roles": []
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

		_, err := srv.GetMembership("<TEAM_ID>", "<MEMBERSHIP_ID>")
		if err != nil {
			t.Errorf("Method GetMembership failed: %v", err)
		}
	})

	t.Run("Test UpdateMembership", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5ea5c16897e",
    "userName": "John Doe",
    "userEmail": "john@appwrite.io",
    "userPhone": "+1 555 555 5555",
    "teamId": "5e5ea5c16897e",
    "teamName": "VIP",
    "invited": "2020-10-15T06:38:00.000+00:00",
    "joined": "2020-10-15T06:38:00.000+00:00",
    "confirm": true,
    "mfa": true,
    "userAccessedAt": "2020-10-15T06:38:00.000+00:00",
    "roles": []
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

		_, err := srv.UpdateMembership("<TEAM_ID>", "<MEMBERSHIP_ID>", []string{})
		if err != nil {
			t.Errorf("Method UpdateMembership failed: %v", err)
		}
	})

	t.Run("Test DeleteMembership", func(t *testing.T) {
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

		_, err := srv.DeleteMembership("<TEAM_ID>", "<MEMBERSHIP_ID>")
		if err != nil {
			t.Errorf("Method DeleteMembership failed: %v", err)
		}
	})

	t.Run("Test UpdateMembershipStatus", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5ea5c16897e",
    "userName": "John Doe",
    "userEmail": "john@appwrite.io",
    "userPhone": "+1 555 555 5555",
    "teamId": "5e5ea5c16897e",
    "teamName": "VIP",
    "invited": "2020-10-15T06:38:00.000+00:00",
    "joined": "2020-10-15T06:38:00.000+00:00",
    "confirm": true,
    "mfa": true,
    "userAccessedAt": "2020-10-15T06:38:00.000+00:00",
    "roles": []
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

		_, err := srv.UpdateMembershipStatus("<TEAM_ID>", "<MEMBERSHIP_ID>", "<USER_ID>", "<SECRET>")
		if err != nil {
			t.Errorf("Method UpdateMembershipStatus failed: %v", err)
		}
	})

	t.Run("Test GetPrefs", func(t *testing.T) {
		mockResponse := `
{
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

		_, err := srv.GetPrefs("<TEAM_ID>")
		if err != nil {
			t.Errorf("Method GetPrefs failed: %v", err)
		}
	})

	t.Run("Test UpdatePrefs", func(t *testing.T) {
		mockResponse := `
{
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

		_, err := srv.UpdatePrefs("<TEAM_ID>", map[string]interface{}{})
		if err != nil {
			t.Errorf("Method UpdatePrefs failed: %v", err)
		}
	})
}
