package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v2/client"
)

func TestUsers(t *testing.T) {
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
    "users": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "John Doe",
            "registration": "2020-10-15T06:38:00.000+00:00",
            "status": true,
            "labels": [],
            "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
            "email": "john@appwrite.io",
            "phone": "+4930901820",
            "emailVerification": true,
            "phoneVerification": true,
            "mfa": true,
            "prefs": {
            },
            "targets": [
                {
                    "$id": "259125845563242502",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "name": "Apple iPhone 12",
                    "userId": "259125845563242502",
                    "providerType": "email",
                    "identifier": "token",
                    "expired": true
                }
            ],
            "accessedAt": "2020-10-15T06:38:00.000+00:00"
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
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.Create("<USER_ID>")
		if err != nil {
			t.Errorf("Method Create failed: %v", err)
		}
	})

	t.Run("Test CreateArgon2User", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateArgon2User("<USER_ID>", "email@example.com", "password")
		if err != nil {
			t.Errorf("Method CreateArgon2User failed: %v", err)
		}
	})

	t.Run("Test CreateBcryptUser", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateBcryptUser("<USER_ID>", "email@example.com", "password")
		if err != nil {
			t.Errorf("Method CreateBcryptUser failed: %v", err)
		}
	})

	t.Run("Test ListIdentities", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "identities": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "userId": "5e5bb8c16897e",
            "provider": "email",
            "providerUid": "5e5bb8c16897e",
            "providerEmail": "user@example.com",
            "providerAccessToken": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",
            "providerAccessTokenExpiry": "2020-10-15T06:38:00.000+00:00",
            "providerRefreshToken": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3"
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

		_, err := srv.ListIdentities()
		if err != nil {
			t.Errorf("Method ListIdentities failed: %v", err)
		}
	})

	t.Run("Test DeleteIdentity", func(t *testing.T) {
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

		_, err := srv.DeleteIdentity("<IDENTITY_ID>")
		if err != nil {
			t.Errorf("Method DeleteIdentity failed: %v", err)
		}
	})

	t.Run("Test CreateMD5User", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateMD5User("<USER_ID>", "email@example.com", "password")
		if err != nil {
			t.Errorf("Method CreateMD5User failed: %v", err)
		}
	})

	t.Run("Test CreatePHPassUser", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreatePHPassUser("<USER_ID>", "email@example.com", "password")
		if err != nil {
			t.Errorf("Method CreatePHPassUser failed: %v", err)
		}
	})

	t.Run("Test CreateScryptUser", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateScryptUser("<USER_ID>", "email@example.com", "password", "<PASSWORD_SALT>", 1, 1, 1, 1)
		if err != nil {
			t.Errorf("Method CreateScryptUser failed: %v", err)
		}
	})

	t.Run("Test CreateScryptModifiedUser", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateScryptModifiedUser("<USER_ID>", "email@example.com", "password", "<PASSWORD_SALT>", "<PASSWORD_SALT_SEPARATOR>", "<PASSWORD_SIGNER_KEY>")
		if err != nil {
			t.Errorf("Method CreateScryptModifiedUser failed: %v", err)
		}
	})

	t.Run("Test CreateSHAUser", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateSHAUser("<USER_ID>", "email@example.com", "password")
		if err != nil {
			t.Errorf("Method CreateSHAUser failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.Get("<USER_ID>")
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

		_, err := srv.Delete("<USER_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})

	t.Run("Test UpdateEmail", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateEmail("<USER_ID>", "email@example.com")
		if err != nil {
			t.Errorf("Method UpdateEmail failed: %v", err)
		}
	})

	t.Run("Test UpdateImpersonator", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateImpersonator("<USER_ID>", true)
		if err != nil {
			t.Errorf("Method UpdateImpersonator failed: %v", err)
		}
	})

	t.Run("Test CreateJWT", func(t *testing.T) {
		mockResponse := `
{
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
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

		_, err := srv.CreateJWT("<USER_ID>")
		if err != nil {
			t.Errorf("Method CreateJWT failed: %v", err)
		}
	})

	t.Run("Test UpdateLabels", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateLabels("<USER_ID>", []string{})
		if err != nil {
			t.Errorf("Method UpdateLabels failed: %v", err)
		}
	})

	t.Run("Test ListLogs", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "logs": [
        {
            "event": "account.sessions.create",
            "userId": "610fc2f985ee0",
            "userEmail": "john@appwrite.io",
            "userName": "John Doe",
            "mode": "admin",
            "ip": "127.0.0.1",
            "time": "2020-10-15T06:38:00.000+00:00",
            "osCode": "Mac",
            "osName": "Mac",
            "osVersion": "Mac",
            "clientType": "browser",
            "clientCode": "CM",
            "clientName": "Chrome Mobile iOS",
            "clientVersion": "84.0",
            "clientEngine": "WebKit",
            "clientEngineVersion": "605.1.15",
            "deviceName": "smartphone",
            "deviceBrand": "Google",
            "deviceModel": "Nexus 5",
            "countryCode": "US",
            "countryName": "United States"
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

		_, err := srv.ListLogs("<USER_ID>")
		if err != nil {
			t.Errorf("Method ListLogs failed: %v", err)
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
            "teamId": "5e5ea5c16897e",
            "teamName": "VIP",
            "invited": "2020-10-15T06:38:00.000+00:00",
            "joined": "2020-10-15T06:38:00.000+00:00",
            "confirm": true,
            "mfa": true,
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

		_, err := srv.ListMemberships("<USER_ID>")
		if err != nil {
			t.Errorf("Method ListMemberships failed: %v", err)
		}
	})

	t.Run("Test UpdateMfa", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateMfa("<USER_ID>", true)
		if err != nil {
			t.Errorf("Method UpdateMfa failed: %v", err)
		}
	})

	t.Run("Test UpdateMFA", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateMFA("<USER_ID>", true)
		if err != nil {
			t.Errorf("Method UpdateMFA failed: %v", err)
		}
	})

	t.Run("Test DeleteMfaAuthenticator", func(t *testing.T) {
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

		_, err := srv.DeleteMfaAuthenticator("<USER_ID>", "totp")
		if err != nil {
			t.Errorf("Method DeleteMfaAuthenticator failed: %v", err)
		}
	})

	t.Run("Test DeleteMFAAuthenticator", func(t *testing.T) {
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

		_, err := srv.DeleteMFAAuthenticator("<USER_ID>", "totp")
		if err != nil {
			t.Errorf("Method DeleteMFAAuthenticator failed: %v", err)
		}
	})

	t.Run("Test ListMfaFactors", func(t *testing.T) {
		mockResponse := `
{
    "totp": true,
    "phone": true,
    "email": true,
    "recoveryCode": true
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

		_, err := srv.ListMfaFactors("<USER_ID>")
		if err != nil {
			t.Errorf("Method ListMfaFactors failed: %v", err)
		}
	})

	t.Run("Test ListMFAFactors", func(t *testing.T) {
		mockResponse := `
{
    "totp": true,
    "phone": true,
    "email": true,
    "recoveryCode": true
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

		_, err := srv.ListMFAFactors("<USER_ID>")
		if err != nil {
			t.Errorf("Method ListMFAFactors failed: %v", err)
		}
	})

	t.Run("Test GetMfaRecoveryCodes", func(t *testing.T) {
		mockResponse := `
{
    "recoveryCodes": []
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

		_, err := srv.GetMfaRecoveryCodes("<USER_ID>")
		if err != nil {
			t.Errorf("Method GetMfaRecoveryCodes failed: %v", err)
		}
	})

	t.Run("Test GetMFARecoveryCodes", func(t *testing.T) {
		mockResponse := `
{
    "recoveryCodes": []
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

		_, err := srv.GetMFARecoveryCodes("<USER_ID>")
		if err != nil {
			t.Errorf("Method GetMFARecoveryCodes failed: %v", err)
		}
	})

	t.Run("Test UpdateMfaRecoveryCodes", func(t *testing.T) {
		mockResponse := `
{
    "recoveryCodes": []
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

		_, err := srv.UpdateMfaRecoveryCodes("<USER_ID>")
		if err != nil {
			t.Errorf("Method UpdateMfaRecoveryCodes failed: %v", err)
		}
	})

	t.Run("Test UpdateMFARecoveryCodes", func(t *testing.T) {
		mockResponse := `
{
    "recoveryCodes": []
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

		_, err := srv.UpdateMFARecoveryCodes("<USER_ID>")
		if err != nil {
			t.Errorf("Method UpdateMFARecoveryCodes failed: %v", err)
		}
	})

	t.Run("Test CreateMfaRecoveryCodes", func(t *testing.T) {
		mockResponse := `
{
    "recoveryCodes": []
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

		_, err := srv.CreateMfaRecoveryCodes("<USER_ID>")
		if err != nil {
			t.Errorf("Method CreateMfaRecoveryCodes failed: %v", err)
		}
	})

	t.Run("Test CreateMFARecoveryCodes", func(t *testing.T) {
		mockResponse := `
{
    "recoveryCodes": []
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

		_, err := srv.CreateMFARecoveryCodes("<USER_ID>")
		if err != nil {
			t.Errorf("Method CreateMFARecoveryCodes failed: %v", err)
		}
	})

	t.Run("Test UpdateName", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateName("<USER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method UpdateName failed: %v", err)
		}
	})

	t.Run("Test UpdatePassword", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdatePassword("<USER_ID>", "")
		if err != nil {
			t.Errorf("Method UpdatePassword failed: %v", err)
		}
	})

	t.Run("Test UpdatePhone", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdatePhone("<USER_ID>", "+12065550100")
		if err != nil {
			t.Errorf("Method UpdatePhone failed: %v", err)
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

		_, err := srv.GetPrefs("<USER_ID>")
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
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdatePrefs("<USER_ID>", map[string]interface{}{})
		if err != nil {
			t.Errorf("Method UpdatePrefs failed: %v", err)
		}
	})

	t.Run("Test ListSessions", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "sessions": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "userId": "5e5bb8c16897e",
            "expire": "2020-10-15T06:38:00.000+00:00",
            "provider": "email",
            "providerUid": "user@example.com",
            "providerAccessToken": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",
            "providerAccessTokenExpiry": "2020-10-15T06:38:00.000+00:00",
            "providerRefreshToken": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",
            "ip": "127.0.0.1",
            "osCode": "Mac",
            "osName": "Mac",
            "osVersion": "Mac",
            "clientType": "browser",
            "clientCode": "CM",
            "clientName": "Chrome Mobile iOS",
            "clientVersion": "84.0",
            "clientEngine": "WebKit",
            "clientEngineVersion": "605.1.15",
            "deviceName": "smartphone",
            "deviceBrand": "Google",
            "deviceModel": "Nexus 5",
            "countryCode": "US",
            "countryName": "United States",
            "current": true,
            "factors": [],
            "secret": "5e5bb8c16897e",
            "mfaUpdatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.ListSessions("<USER_ID>")
		if err != nil {
			t.Errorf("Method ListSessions failed: %v", err)
		}
	})

	t.Run("Test CreateSession", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5bb8c16897e",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "provider": "email",
    "providerUid": "user@example.com",
    "providerAccessToken": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",
    "providerAccessTokenExpiry": "2020-10-15T06:38:00.000+00:00",
    "providerRefreshToken": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",
    "ip": "127.0.0.1",
    "osCode": "Mac",
    "osName": "Mac",
    "osVersion": "Mac",
    "clientType": "browser",
    "clientCode": "CM",
    "clientName": "Chrome Mobile iOS",
    "clientVersion": "84.0",
    "clientEngine": "WebKit",
    "clientEngineVersion": "605.1.15",
    "deviceName": "smartphone",
    "deviceBrand": "Google",
    "deviceModel": "Nexus 5",
    "countryCode": "US",
    "countryName": "United States",
    "current": true,
    "factors": [],
    "secret": "5e5bb8c16897e",
    "mfaUpdatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateSession("<USER_ID>")
		if err != nil {
			t.Errorf("Method CreateSession failed: %v", err)
		}
	})

	t.Run("Test DeleteSessions", func(t *testing.T) {
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

		_, err := srv.DeleteSessions("<USER_ID>")
		if err != nil {
			t.Errorf("Method DeleteSessions failed: %v", err)
		}
	})

	t.Run("Test DeleteSession", func(t *testing.T) {
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

		_, err := srv.DeleteSession("<USER_ID>", "<SESSION_ID>")
		if err != nil {
			t.Errorf("Method DeleteSession failed: %v", err)
		}
	})

	t.Run("Test UpdateStatus", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateStatus("<USER_ID>", true)
		if err != nil {
			t.Errorf("Method UpdateStatus failed: %v", err)
		}
	})

	t.Run("Test ListTargets", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
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

		_, err := srv.ListTargets("<USER_ID>")
		if err != nil {
			t.Errorf("Method ListTargets failed: %v", err)
		}
	})

	t.Run("Test CreateTarget", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Apple iPhone 12",
    "userId": "259125845563242502",
    "providerType": "email",
    "identifier": "token",
    "expired": true
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

		_, err := srv.CreateTarget("<USER_ID>", "<TARGET_ID>", "email", "<IDENTIFIER>")
		if err != nil {
			t.Errorf("Method CreateTarget failed: %v", err)
		}
	})

	t.Run("Test GetTarget", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Apple iPhone 12",
    "userId": "259125845563242502",
    "providerType": "email",
    "identifier": "token",
    "expired": true
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

		_, err := srv.GetTarget("<USER_ID>", "<TARGET_ID>")
		if err != nil {
			t.Errorf("Method GetTarget failed: %v", err)
		}
	})

	t.Run("Test UpdateTarget", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Apple iPhone 12",
    "userId": "259125845563242502",
    "providerType": "email",
    "identifier": "token",
    "expired": true
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

		_, err := srv.UpdateTarget("<USER_ID>", "<TARGET_ID>")
		if err != nil {
			t.Errorf("Method UpdateTarget failed: %v", err)
		}
	})

	t.Run("Test DeleteTarget", func(t *testing.T) {
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

		_, err := srv.DeleteTarget("<USER_ID>", "<TARGET_ID>")
		if err != nil {
			t.Errorf("Method DeleteTarget failed: %v", err)
		}
	})

	t.Run("Test CreateToken", func(t *testing.T) {
		mockResponse := `
{
    "$id": "bb8ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5ea5c168bb8",
    "secret": "string",
    "expire": "2020-10-15T06:38:00.000+00:00",
    "phrase": "Golden Fox"
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

		_, err := srv.CreateToken("<USER_ID>")
		if err != nil {
			t.Errorf("Method CreateToken failed: %v", err)
		}
	})

	t.Run("Test UpdateEmailVerification", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateEmailVerification("<USER_ID>", true)
		if err != nil {
			t.Errorf("Method UpdateEmailVerification failed: %v", err)
		}
	})

	t.Run("Test UpdatePhoneVerification", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "John Doe",
    "registration": "2020-10-15T06:38:00.000+00:00",
    "status": true,
    "labels": [],
    "passwordUpdate": "2020-10-15T06:38:00.000+00:00",
    "email": "john@appwrite.io",
    "phone": "+4930901820",
    "emailVerification": true,
    "phoneVerification": true,
    "mfa": true,
    "prefs": {
    },
    "targets": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Apple iPhone 12",
            "userId": "259125845563242502",
            "providerType": "email",
            "identifier": "token",
            "expired": true
        }
    ],
    "accessedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdatePhoneVerification("<USER_ID>", true)
		if err != nil {
			t.Errorf("Method UpdatePhoneVerification failed: %v", err)
		}
	})
}
