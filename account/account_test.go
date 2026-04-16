package account

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v3/client"
)

func TestAccount(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

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

		_, err := srv.Get()
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
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

		_, err := srv.Create("<USER_ID>", "email@example.com", "")
		if err != nil {
			t.Errorf("Method Create failed: %v", err)
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

		_, err := srv.UpdateEmail("email@example.com", "password")
		if err != nil {
			t.Errorf("Method UpdateEmail failed: %v", err)
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

		_, err := srv.CreateJWT()
		if err != nil {
			t.Errorf("Method CreateJWT failed: %v", err)
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
            "userType": "user",
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

		_, err := srv.ListLogs()
		if err != nil {
			t.Errorf("Method ListLogs failed: %v", err)
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

		_, err := srv.UpdateMFA(true)
		if err != nil {
			t.Errorf("Method UpdateMFA failed: %v", err)
		}
	})

	t.Run("Test CreateMfaAuthenticator", func(t *testing.T) {
		mockResponse := `
{
    "secret": "[SHARED_SECRET]",
    "uri": "otpauth://totp/appwrite:user@example.com?secret=[SHARED_SECRET]&amp;issuer=appwrite"
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

		_, err := srv.CreateMfaAuthenticator("totp")
		if err != nil {
			t.Errorf("Method CreateMfaAuthenticator failed: %v", err)
		}
	})

	t.Run("Test CreateMFAAuthenticator", func(t *testing.T) {
		mockResponse := `
{
    "secret": "[SHARED_SECRET]",
    "uri": "otpauth://totp/appwrite:user@example.com?secret=[SHARED_SECRET]&amp;issuer=appwrite"
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

		_, err := srv.CreateMFAAuthenticator("totp")
		if err != nil {
			t.Errorf("Method CreateMFAAuthenticator failed: %v", err)
		}
	})

	t.Run("Test UpdateMfaAuthenticator", func(t *testing.T) {
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

		_, err := srv.UpdateMfaAuthenticator("totp", "<OTP>")
		if err != nil {
			t.Errorf("Method UpdateMfaAuthenticator failed: %v", err)
		}
	})

	t.Run("Test UpdateMFAAuthenticator", func(t *testing.T) {
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

		_, err := srv.UpdateMFAAuthenticator("totp", "<OTP>")
		if err != nil {
			t.Errorf("Method UpdateMFAAuthenticator failed: %v", err)
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

		_, err := srv.DeleteMfaAuthenticator("totp")
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

		_, err := srv.DeleteMFAAuthenticator("totp")
		if err != nil {
			t.Errorf("Method DeleteMFAAuthenticator failed: %v", err)
		}
	})

	t.Run("Test CreateMfaChallenge", func(t *testing.T) {
		mockResponse := `
{
    "$id": "bb8ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5ea5c168bb8",
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

		_, err := srv.CreateMfaChallenge("email")
		if err != nil {
			t.Errorf("Method CreateMfaChallenge failed: %v", err)
		}
	})

	t.Run("Test CreateMFAChallenge", func(t *testing.T) {
		mockResponse := `
{
    "$id": "bb8ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "userId": "5e5ea5c168bb8",
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

		_, err := srv.CreateMFAChallenge("email")
		if err != nil {
			t.Errorf("Method CreateMFAChallenge failed: %v", err)
		}
	})

	t.Run("Test UpdateMfaChallenge", func(t *testing.T) {
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
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateMfaChallenge("<CHALLENGE_ID>", "<OTP>")
		if err != nil {
			t.Errorf("Method UpdateMfaChallenge failed: %v", err)
		}
	})

	t.Run("Test UpdateMFAChallenge", func(t *testing.T) {
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
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateMFAChallenge("<CHALLENGE_ID>", "<OTP>")
		if err != nil {
			t.Errorf("Method UpdateMFAChallenge failed: %v", err)
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

		_, err := srv.ListMfaFactors()
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

		_, err := srv.ListMFAFactors()
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

		_, err := srv.GetMfaRecoveryCodes()
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

		_, err := srv.GetMFARecoveryCodes()
		if err != nil {
			t.Errorf("Method GetMFARecoveryCodes failed: %v", err)
		}
	})

	t.Run("Test CreateMfaRecoveryCodes", func(t *testing.T) {
		mockResponse := `
{
    "recoveryCodes": []
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

		_, err := srv.CreateMfaRecoveryCodes()
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
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateMFARecoveryCodes()
		if err != nil {
			t.Errorf("Method CreateMFARecoveryCodes failed: %v", err)
		}
	})

	t.Run("Test UpdateMfaRecoveryCodes", func(t *testing.T) {
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

		_, err := srv.UpdateMfaRecoveryCodes()
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
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateMFARecoveryCodes()
		if err != nil {
			t.Errorf("Method UpdateMFARecoveryCodes failed: %v", err)
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

		_, err := srv.UpdateName("<NAME>")
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

		_, err := srv.UpdatePassword("")
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

		_, err := srv.UpdatePhone("+12065550100", "password")
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

		_, err := srv.GetPrefs()
		if err != nil {
			t.Errorf("Method GetPrefs failed: %v", err)
		}
	})

	t.Run("Test UpdatePrefs", func(t *testing.T) {
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

		_, err := srv.UpdatePrefs(map[string]interface{}{})
		if err != nil {
			t.Errorf("Method UpdatePrefs failed: %v", err)
		}
	})

	t.Run("Test CreateRecovery", func(t *testing.T) {
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

		_, err := srv.CreateRecovery("email@example.com", "https://example.com")
		if err != nil {
			t.Errorf("Method CreateRecovery failed: %v", err)
		}
	})

	t.Run("Test UpdateRecovery", func(t *testing.T) {
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
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateRecovery("<USER_ID>", "<SECRET>", "")
		if err != nil {
			t.Errorf("Method UpdateRecovery failed: %v", err)
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

		_, err := srv.ListSessions()
		if err != nil {
			t.Errorf("Method ListSessions failed: %v", err)
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

		_, err := srv.DeleteSessions()
		if err != nil {
			t.Errorf("Method DeleteSessions failed: %v", err)
		}
	})

	t.Run("Test CreateAnonymousSession", func(t *testing.T) {
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

		_, err := srv.CreateAnonymousSession()
		if err != nil {
			t.Errorf("Method CreateAnonymousSession failed: %v", err)
		}
	})

	t.Run("Test CreateEmailPasswordSession", func(t *testing.T) {
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

		_, err := srv.CreateEmailPasswordSession("email@example.com", "password")
		if err != nil {
			t.Errorf("Method CreateEmailPasswordSession failed: %v", err)
		}
	})

	t.Run("Test UpdateMagicURLSession", func(t *testing.T) {
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
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateMagicURLSession("<USER_ID>", "<SECRET>")
		if err != nil {
			t.Errorf("Method UpdateMagicURLSession failed: %v", err)
		}
	})

	t.Run("Test UpdatePhoneSession", func(t *testing.T) {
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
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdatePhoneSession("<USER_ID>", "<SECRET>")
		if err != nil {
			t.Errorf("Method UpdatePhoneSession failed: %v", err)
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

		_, err := srv.CreateSession("<USER_ID>", "<SECRET>")
		if err != nil {
			t.Errorf("Method CreateSession failed: %v", err)
		}
	})

	t.Run("Test GetSession", func(t *testing.T) {
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
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetSession("<SESSION_ID>")
		if err != nil {
			t.Errorf("Method GetSession failed: %v", err)
		}
	})

	t.Run("Test UpdateSession", func(t *testing.T) {
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
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateSession("<SESSION_ID>")
		if err != nil {
			t.Errorf("Method UpdateSession failed: %v", err)
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

		_, err := srv.DeleteSession("<SESSION_ID>")
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

		_, err := srv.UpdateStatus()
		if err != nil {
			t.Errorf("Method UpdateStatus failed: %v", err)
		}
	})

	t.Run("Test CreateEmailToken", func(t *testing.T) {
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

		_, err := srv.CreateEmailToken("<USER_ID>", "email@example.com")
		if err != nil {
			t.Errorf("Method CreateEmailToken failed: %v", err)
		}
	})

	t.Run("Test CreateMagicURLToken", func(t *testing.T) {
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

		_, err := srv.CreateMagicURLToken("<USER_ID>", "email@example.com")
		if err != nil {
			t.Errorf("Method CreateMagicURLToken failed: %v", err)
		}
	})

	t.Run("Test CreateOAuth2Token", func(t *testing.T) {
		mockResponse := `true`

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

		_, err := srv.CreateOAuth2Token("amazon")
		if err != nil {
			t.Errorf("Method CreateOAuth2Token failed: %v", err)
		}
	})

	t.Run("Test CreatePhoneToken", func(t *testing.T) {
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

		_, err := srv.CreatePhoneToken("<USER_ID>", "+12065550100")
		if err != nil {
			t.Errorf("Method CreatePhoneToken failed: %v", err)
		}
	})

	t.Run("Test CreateEmailVerification", func(t *testing.T) {
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

		_, err := srv.CreateEmailVerification("https://example.com")
		if err != nil {
			t.Errorf("Method CreateEmailVerification failed: %v", err)
		}
	})

	t.Run("Test CreateVerification", func(t *testing.T) {
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

		_, err := srv.CreateVerification("https://example.com")
		if err != nil {
			t.Errorf("Method CreateVerification failed: %v", err)
		}
	})

	t.Run("Test UpdateEmailVerification", func(t *testing.T) {
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
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateEmailVerification("<USER_ID>", "<SECRET>")
		if err != nil {
			t.Errorf("Method UpdateEmailVerification failed: %v", err)
		}
	})

	t.Run("Test UpdateVerification", func(t *testing.T) {
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
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateVerification("<USER_ID>", "<SECRET>")
		if err != nil {
			t.Errorf("Method UpdateVerification failed: %v", err)
		}
	})

	t.Run("Test CreatePhoneVerification", func(t *testing.T) {
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

		_, err := srv.CreatePhoneVerification()
		if err != nil {
			t.Errorf("Method CreatePhoneVerification failed: %v", err)
		}
	})

	t.Run("Test UpdatePhoneVerification", func(t *testing.T) {
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
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdatePhoneVerification("<USER_ID>", "<SECRET>")
		if err != nil {
			t.Errorf("Method UpdatePhoneVerification failed: %v", err)
		}
	})
}
