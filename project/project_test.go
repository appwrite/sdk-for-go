package project

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v3/client"
	"github.com/appwrite/sdk-for-go/v3/models"
)

func TestProject(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

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

		_, err := srv.Delete()
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})

	t.Run("Test UpdateAuthMethod", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateAuthMethod("email-password", true)
		if err != nil {
			t.Errorf("Method UpdateAuthMethod failed: %v", err)
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

	t.Run("Test CreateEphemeralKey", func(t *testing.T) {
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

		_, err := srv.CreateEphemeralKey([]string{}, 1)
		if err != nil {
			t.Errorf("Method CreateEphemeralKey failed: %v", err)
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

	t.Run("Test UpdateLabels", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateLabels([]string{})
		if err != nil {
			t.Errorf("Method UpdateLabels failed: %v", err)
		}
	})

	t.Run("Test ListMockPhones", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "mockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.ListMockPhones()
		if err != nil {
			t.Errorf("Method ListMockPhones failed: %v", err)
		}
	})

	t.Run("Test CreateMockPhone", func(t *testing.T) {
		mockResponse := `
{
    "number": "+1612842323",
    "otp": "123456",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateMockPhone("+12065550100", "<OTP>")
		if err != nil {
			t.Errorf("Method CreateMockPhone failed: %v", err)
		}
	})

	t.Run("Test GetMockPhone", func(t *testing.T) {
		mockResponse := `
{
    "number": "+1612842323",
    "otp": "123456",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.GetMockPhone("+12065550100")
		if err != nil {
			t.Errorf("Method GetMockPhone failed: %v", err)
		}
	})

	t.Run("Test UpdateMockPhone", func(t *testing.T) {
		mockResponse := `
{
    "number": "+1612842323",
    "otp": "123456",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateMockPhone("+12065550100", "<OTP>")
		if err != nil {
			t.Errorf("Method UpdateMockPhone failed: %v", err)
		}
	})

	t.Run("Test DeleteMockPhone", func(t *testing.T) {
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

		_, err := srv.DeleteMockPhone("+12065550100")
		if err != nil {
			t.Errorf("Method DeleteMockPhone failed: %v", err)
		}
	})

	t.Run("Test ListOAuth2Providers", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "providers": []
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

		_, err := srv.ListOAuth2Providers()
		if err != nil {
			t.Errorf("Method ListOAuth2Providers failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Amazon", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "amzn1.application-oa2-client.87400c00000000000000000000063d5b2",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Amazon()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Amazon failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Apple", func(t *testing.T) {
		mockResponse := `
{
    "$id": "apple",
    "enabled": true,
    "serviceId": "ip.appwrite.app.web",
    "keyId": "P4000000N8",
    "teamId": "D4000000R6",
    "p8File": "-----BEGIN PRIVATE KEY-----MIGTAg...jy2Xbna-----END PRIVATE KEY-----"
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

		_, err := srv.UpdateOAuth2Apple()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Apple failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Auth0", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "OaOkIA000000000000000000005KLSYq",
    "clientSecret": "&lt;CLIENT_SECRET&gt;",
    "endpoint": "example.us.auth0.com"
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

		_, err := srv.UpdateOAuth2Auth0()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Auth0 failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Authentik", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "dTKOPa0000000000000000000000000000e7G8hv",
    "clientSecret": "&lt;CLIENT_SECRET&gt;",
    "endpoint": "example.authentik.com"
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

		_, err := srv.UpdateOAuth2Authentik()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Authentik failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Autodesk", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "5zw90v00000000000000000000kVYXN7",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Autodesk()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Autodesk failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Bitbucket", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "key": "Knt70000000000ByRc",
    "secret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Bitbucket()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Bitbucket failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Bitly", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "d95151000000000000000000000000000067af9b",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Bitly()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Bitly failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Box", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "deglcs00000000000000000000x2og6y",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Box()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Box failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Dailymotion", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "apiKey": "07a9000000000000067f",
    "apiSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Dailymotion()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Dailymotion failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Discord", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "950722000000343754",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Discord()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Discord failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Disqus", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "publicKey": "cgegH70000000000000000000000000000000000000000000000000000Hr1nYX",
    "secretKey": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Disqus()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Disqus failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Dropbox", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "appKey": "jl000000000009t",
    "appSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Dropbox()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Dropbox failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Etsy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "keyString": "nsgzxh0000000000008j85a2",
    "sharedSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Etsy()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Etsy failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Facebook", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "appId": "260600000007694",
    "appSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Facebook()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Facebook failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Figma", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "byay5H0000000000VtiI40",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Figma()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Figma failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2FusionAuth", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "b2222c00-0000-0000-0000-000000862097",
    "clientSecret": "&lt;CLIENT_SECRET&gt;",
    "endpoint": "example.fusionauth.io"
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

		_, err := srv.UpdateOAuth2FusionAuth()
		if err != nil {
			t.Errorf("Method UpdateOAuth2FusionAuth failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2GitHub", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "e4d87900000000540733",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2GitHub()
		if err != nil {
			t.Errorf("Method UpdateOAuth2GitHub failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Gitlab", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "applicationId": "d41ffe0000000000000000000000000000000000000000000000000000d5e252",
    "secret": "&lt;CLIENT_SECRET&gt;",
    "endpoint": "https://gitlab.com"
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

		_, err := srv.UpdateOAuth2Gitlab()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Gitlab failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Google", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "120000000095-92ifjb00000000000000000000g7ijfb.apps.googleusercontent.com",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Google()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Google failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Keycloak", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "appwrite-o0000000st-app",
    "clientSecret": "&lt;CLIENT_SECRET&gt;",
    "endpoint": "keycloak.example.com",
    "realmName": "appwrite-realm"
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

		_, err := srv.UpdateOAuth2Keycloak()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Keycloak failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Kick", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "01KQ7C00000000000001MFHS32",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Kick()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Kick failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Linkedin", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "770000000000dv",
    "primaryClientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Linkedin()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Linkedin failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Microsoft", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "applicationId": "00001111-aaaa-2222-bbbb-3333cccc4444",
    "applicationSecret": "&lt;CLIENT_SECRET&gt;",
    "tenant": "common"
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

		_, err := srv.UpdateOAuth2Microsoft()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Microsoft failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Notion", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "oauthClientId": "341d8700-0000-0000-0000-000000446ee3",
    "oauthClientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Notion()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Notion failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Oidc", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "qibI2x0000000000000000000000000006L2YFoG",
    "clientSecret": "&lt;CLIENT_SECRET&gt;",
    "wellKnownURL": "https://myoauth.com/.well-known/openid-configuration",
    "authorizationURL": "https://myoauth.com/oauth2/authorize",
    "tokenURL": "https://myoauth.com/oauth2/token",
    "userInfoURL": "https://myoauth.com/oauth2/userinfo"
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

		_, err := srv.UpdateOAuth2Oidc()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Oidc failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Okta", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "0oa00000000000000698",
    "clientSecret": "&lt;CLIENT_SECRET&gt;",
    "domain": "trial-6400025.okta.com",
    "authorizationServerId": "aus000000000000000h7z"
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

		_, err := srv.UpdateOAuth2Okta()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Okta failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Paypal", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "AdhIEG7-000000000000-0000000000000000000000000000000-0000000000000000000000-2pyB",
    "secretKey": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Paypal()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Paypal failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2PaypalSandbox", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "AdhIEG7-000000000000-0000000000000000000000000000000-0000000000000000000000-2pyB",
    "secretKey": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2PaypalSandbox()
		if err != nil {
			t.Errorf("Method UpdateOAuth2PaypalSandbox failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Podio", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "appwrite-oauth-test-app",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Podio()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Podio failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Salesforce", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "customerKey": "3MVG9I0000000000000000000000000000000000000000000000000000000000000000000000000C5Aejq",
    "customerSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Salesforce()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Salesforce failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Slack", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "23000000089.15000000000023",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Slack()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Slack failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Spotify", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "6ec271000000000000000000009beace",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Spotify()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Spotify failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Stripe", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "ca_UKibXX0000000000000000000006byvR",
    "apiSecretKey": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Stripe()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Stripe failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Tradeshift", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "oauth2ClientId": "appwrite-test-org.appwrite-test-app",
    "oauth2ClientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Tradeshift()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Tradeshift failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2TradeshiftSandbox", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "oauth2ClientId": "appwrite-test-org.appwrite-test-app",
    "oauth2ClientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2TradeshiftSandbox()
		if err != nil {
			t.Errorf("Method UpdateOAuth2TradeshiftSandbox failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Twitch", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "vvi0in000000000000000000ikmt9p",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Twitch()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Twitch failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2WordPress", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "130005",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2WordPress()
		if err != nil {
			t.Errorf("Method UpdateOAuth2WordPress failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2X", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "customerKey": "slzZV0000000000000NFLaWT",
    "secretKey": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2X()
		if err != nil {
			t.Errorf("Method UpdateOAuth2X failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Yahoo", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "dj0yJm000000000000000000000000000000000000000000000000000000000000000000000000000000000000Z4PWRm",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Yahoo()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Yahoo failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Yandex", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "6a8a6a0000000000000000000091483c",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Yandex()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Yandex failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Zoho", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "1000.83C178000000000000000000RPNX0B",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Zoho()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Zoho failed: %v", err)
		}
	})

	t.Run("Test UpdateOAuth2Zoom", func(t *testing.T) {
		mockResponse := `
{
    "$id": "github",
    "enabled": true,
    "clientId": "QMAC00000000000000w0AQ",
    "clientSecret": "&lt;CLIENT_SECRET&gt;"
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

		_, err := srv.UpdateOAuth2Zoom()
		if err != nil {
			t.Errorf("Method UpdateOAuth2Zoom failed: %v", err)
		}
	})

	t.Run("Test GetOAuth2Provider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "microsoft",
    "enabled": true,
    "applicationId": "00001111-aaaa-2222-bbbb-3333cccc4444",
    "applicationSecret": "&lt;CLIENT_SECRET&gt;",
    "tenant": "common"
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

		response, err := srv.GetOAuth2Provider("amazon")
		if err != nil {
			t.Errorf("Method GetOAuth2Provider failed: %v", err)
		}
		if _, ok := response.(*models.OAuth2Microsoft); !ok {
			t.Errorf("Expected response type *models.OAuth2Microsoft, got %T", response)
		}
	})

	t.Run("Test ListPlatforms", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "platforms": []
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

		_, err := srv.ListPlatforms()
		if err != nil {
			t.Errorf("Method ListPlatforms failed: %v", err)
		}
	})

	t.Run("Test CreateAndroidPlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "applicationId": "com.company.appname"
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

		_, err := srv.CreateAndroidPlatform("<PLATFORM_ID>", "<NAME>", "<APPLICATION_ID>")
		if err != nil {
			t.Errorf("Method CreateAndroidPlatform failed: %v", err)
		}
	})

	t.Run("Test UpdateAndroidPlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "applicationId": "com.company.appname"
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

		_, err := srv.UpdateAndroidPlatform("<PLATFORM_ID>", "<NAME>", "<APPLICATION_ID>")
		if err != nil {
			t.Errorf("Method UpdateAndroidPlatform failed: %v", err)
		}
	})

	t.Run("Test CreateApplePlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "bundleIdentifier": "com.company.appname"
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

		_, err := srv.CreateApplePlatform("<PLATFORM_ID>", "<NAME>", "<BUNDLE_IDENTIFIER>")
		if err != nil {
			t.Errorf("Method CreateApplePlatform failed: %v", err)
		}
	})

	t.Run("Test UpdateApplePlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "bundleIdentifier": "com.company.appname"
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

		_, err := srv.UpdateApplePlatform("<PLATFORM_ID>", "<NAME>", "<BUNDLE_IDENTIFIER>")
		if err != nil {
			t.Errorf("Method UpdateApplePlatform failed: %v", err)
		}
	})

	t.Run("Test CreateLinuxPlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "packageName": "com.company.appname"
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

		_, err := srv.CreateLinuxPlatform("<PLATFORM_ID>", "<NAME>", "<PACKAGE_NAME>")
		if err != nil {
			t.Errorf("Method CreateLinuxPlatform failed: %v", err)
		}
	})

	t.Run("Test UpdateLinuxPlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "packageName": "com.company.appname"
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

		_, err := srv.UpdateLinuxPlatform("<PLATFORM_ID>", "<NAME>", "<PACKAGE_NAME>")
		if err != nil {
			t.Errorf("Method UpdateLinuxPlatform failed: %v", err)
		}
	})

	t.Run("Test CreateWebPlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "hostname": "app.example.com"
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

		_, err := srv.CreateWebPlatform("<PLATFORM_ID>", "<NAME>", "app.example.com")
		if err != nil {
			t.Errorf("Method CreateWebPlatform failed: %v", err)
		}
	})

	t.Run("Test UpdateWebPlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "hostname": "app.example.com"
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

		_, err := srv.UpdateWebPlatform("<PLATFORM_ID>", "<NAME>", "app.example.com")
		if err != nil {
			t.Errorf("Method UpdateWebPlatform failed: %v", err)
		}
	})

	t.Run("Test CreateWindowsPlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "packageIdentifierName": "com.company.appname"
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

		_, err := srv.CreateWindowsPlatform("<PLATFORM_ID>", "<NAME>", "<PACKAGE_IDENTIFIER_NAME>")
		if err != nil {
			t.Errorf("Method CreateWindowsPlatform failed: %v", err)
		}
	})

	t.Run("Test UpdateWindowsPlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "web",
    "packageIdentifierName": "com.company.appname"
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

		_, err := srv.UpdateWindowsPlatform("<PLATFORM_ID>", "<NAME>", "<PACKAGE_IDENTIFIER_NAME>")
		if err != nil {
			t.Errorf("Method UpdateWindowsPlatform failed: %v", err)
		}
	})

	t.Run("Test GetPlatform", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "My Web App",
    "type": "linux",
    "packageName": "com.company.appname"
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

		response, err := srv.GetPlatform("<PLATFORM_ID>")
		if err != nil {
			t.Errorf("Method GetPlatform failed: %v", err)
		}
		if _, ok := response.(*models.PlatformLinux); !ok {
			t.Errorf("Expected response type *models.PlatformLinux, got %T", response)
		}
	})

	t.Run("Test DeletePlatform", func(t *testing.T) {
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

		_, err := srv.DeletePlatform("<PLATFORM_ID>")
		if err != nil {
			t.Errorf("Method DeletePlatform failed: %v", err)
		}
	})

	t.Run("Test ListPolicies", func(t *testing.T) {
		mockResponse := `
{
    "total": 9,
    "policies": []
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

		_, err := srv.ListPolicies()
		if err != nil {
			t.Errorf("Method ListPolicies failed: %v", err)
		}
	})

	t.Run("Test UpdateMembershipPrivacyPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateMembershipPrivacyPolicy()
		if err != nil {
			t.Errorf("Method UpdateMembershipPrivacyPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdatePasswordDictionaryPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdatePasswordDictionaryPolicy(true)
		if err != nil {
			t.Errorf("Method UpdatePasswordDictionaryPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdatePasswordHistoryPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdatePasswordHistoryPolicy(1)
		if err != nil {
			t.Errorf("Method UpdatePasswordHistoryPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdatePasswordPersonalDataPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdatePasswordPersonalDataPolicy(true)
		if err != nil {
			t.Errorf("Method UpdatePasswordPersonalDataPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdateSessionAlertPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateSessionAlertPolicy(true)
		if err != nil {
			t.Errorf("Method UpdateSessionAlertPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdateSessionDurationPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateSessionDurationPolicy(1)
		if err != nil {
			t.Errorf("Method UpdateSessionDurationPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdateSessionInvalidationPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateSessionInvalidationPolicy(true)
		if err != nil {
			t.Errorf("Method UpdateSessionInvalidationPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdateSessionLimitPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateSessionLimitPolicy(1)
		if err != nil {
			t.Errorf("Method UpdateSessionLimitPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdateUserLimitPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateUserLimitPolicy(1)
		if err != nil {
			t.Errorf("Method UpdateUserLimitPolicy failed: %v", err)
		}
	})

	t.Run("Test GetPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "membership-privacy",
    "userId": true,
    "userEmail": true,
    "userPhone": true,
    "userName": true,
    "userMFA": true
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

		response, err := srv.GetPolicy("password-dictionary")
		if err != nil {
			t.Errorf("Method GetPolicy failed: %v", err)
		}
		if _, ok := response.(*models.PolicyMembershipPrivacy); !ok {
			t.Errorf("Expected response type *models.PolicyMembershipPrivacy, got %T", response)
		}
	})

	t.Run("Test UpdateProtocol", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateProtocol("rest", true)
		if err != nil {
			t.Errorf("Method UpdateProtocol failed: %v", err)
		}
	})

	t.Run("Test UpdateService", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateService("account", true)
		if err != nil {
			t.Errorf("Method UpdateService failed: %v", err)
		}
	})

	t.Run("Test UpdateSMTP", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "New Project",
    "description": "This is a new project.",
    "teamId": "1592981250",
    "logo": "5f5c451b403cb",
    "url": "5f5c451b403cb",
    "legalName": "Company LTD.",
    "legalCountry": "US",
    "legalState": "New York",
    "legalCity": "New York City.",
    "legalAddress": "620 Eighth Avenue, New York, NY 10018",
    "legalTaxId": "131102020",
    "authDuration": 60,
    "authLimit": 100,
    "authSessionsLimit": 10,
    "authPasswordHistory": 5,
    "authPasswordDictionary": true,
    "authPersonalDataCheck": true,
    "authDisposableEmails": true,
    "authCanonicalEmails": true,
    "authFreeEmails": true,
    "authMockNumbers": [
        {
            "number": "+1612842323",
            "otp": "123456",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authMembershipsUserId": true,
    "authMembershipsUserPhone": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "string",
            "enabled": true
        }
    ],
    "platforms": [],
    "webhooks": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "My Webhook",
            "url": "https://example.com/webhook",
            "events": [],
            "tls": true,
            "authUsername": "username",
            "authPassword": "password",
            "secret": "ad3d581ca230e2b7059c545e5a",
            "enabled": true,
            "logs": "Failed to connect to remote server.",
            "attempts": 10
        }
    ],
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
    ],
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
    "authEmailPassword": true,
    "authUsersAuthMagicURL": true,
    "authEmailOtp": true,
    "authAnonymous": true,
    "authInvites": true,
    "authJWT": true,
    "authPhone": true,
    "serviceStatusForAccount": true,
    "serviceStatusForAvatars": true,
    "serviceStatusForDatabases": true,
    "serviceStatusForTablesdb": true,
    "serviceStatusForLocale": true,
    "serviceStatusForHealth": true,
    "serviceStatusForProject": true,
    "serviceStatusForStorage": true,
    "serviceStatusForTeams": true,
    "serviceStatusForUsers": true,
    "serviceStatusForVcs": true,
    "serviceStatusForSites": true,
    "serviceStatusForFunctions": true,
    "serviceStatusForProxy": true,
    "serviceStatusForGraphql": true,
    "serviceStatusForMigrations": true,
    "serviceStatusForMessaging": true,
    "protocolStatusForRest": true,
    "protocolStatusForGraphql": true,
    "protocolStatusForWebsocket": true,
    "region": "fra",
    "billingLimits": {
        "bandwidth": 5,
        "storage": 150,
        "users": 200000,
        "executions": 750000,
        "GBHours": 100,
        "imageTransformations": 100,
        "authPhone": 10,
        "budgetLimit": 100
    },
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

		_, err := srv.UpdateSMTP()
		if err != nil {
			t.Errorf("Method UpdateSMTP failed: %v", err)
		}
	})

	t.Run("Test CreateSMTPTest", func(t *testing.T) {
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

		_, err := srv.CreateSMTPTest([]string{})
		if err != nil {
			t.Errorf("Method CreateSMTPTest failed: %v", err)
		}
	})

	t.Run("Test ListEmailTemplates", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "templates": [
        {
            "templateId": "verification",
            "locale": "en_us",
            "message": "Click on the link to verify your account.",
            "senderName": "My User",
            "senderEmail": "mail@appwrite.io",
            "replyToEmail": "emails@appwrite.io",
            "replyToName": "Support Team",
            "subject": "Please verify your email address"
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

		_, err := srv.ListEmailTemplates()
		if err != nil {
			t.Errorf("Method ListEmailTemplates failed: %v", err)
		}
	})

	t.Run("Test UpdateEmailTemplate", func(t *testing.T) {
		mockResponse := `
{
    "templateId": "verification",
    "locale": "en_us",
    "message": "Click on the link to verify your account.",
    "senderName": "My User",
    "senderEmail": "mail@appwrite.io",
    "replyToEmail": "emails@appwrite.io",
    "replyToName": "Support Team",
    "subject": "Please verify your email address"
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

		_, err := srv.UpdateEmailTemplate("verification")
		if err != nil {
			t.Errorf("Method UpdateEmailTemplate failed: %v", err)
		}
	})

	t.Run("Test GetEmailTemplate", func(t *testing.T) {
		mockResponse := `
{
    "templateId": "verification",
    "locale": "en_us",
    "message": "Click on the link to verify your account.",
    "senderName": "My User",
    "senderEmail": "mail@appwrite.io",
    "replyToEmail": "emails@appwrite.io",
    "replyToName": "Support Team",
    "subject": "Please verify your email address"
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

		_, err := srv.GetEmailTemplate("verification")
		if err != nil {
			t.Errorf("Method GetEmailTemplate failed: %v", err)
		}
	})

	t.Run("Test ListVariables", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "variables": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "API_KEY",
            "value": "myPa$$word1",
            "secret": true,
            "resourceType": "function",
            "resourceId": "myAwesomeFunction"
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

		_, err := srv.ListVariables()
		if err != nil {
			t.Errorf("Method ListVariables failed: %v", err)
		}
	})

	t.Run("Test CreateVariable", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "key": "API_KEY",
    "value": "myPa$$word1",
    "secret": true,
    "resourceType": "function",
    "resourceId": "myAwesomeFunction"
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

		_, err := srv.CreateVariable("<VARIABLE_ID>", "<KEY>", "<VALUE>")
		if err != nil {
			t.Errorf("Method CreateVariable failed: %v", err)
		}
	})

	t.Run("Test GetVariable", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "key": "API_KEY",
    "value": "myPa$$word1",
    "secret": true,
    "resourceType": "function",
    "resourceId": "myAwesomeFunction"
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

		_, err := srv.GetVariable("<VARIABLE_ID>")
		if err != nil {
			t.Errorf("Method GetVariable failed: %v", err)
		}
	})

	t.Run("Test UpdateVariable", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "key": "API_KEY",
    "value": "myPa$$word1",
    "secret": true,
    "resourceType": "function",
    "resourceId": "myAwesomeFunction"
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

		_, err := srv.UpdateVariable("<VARIABLE_ID>")
		if err != nil {
			t.Errorf("Method UpdateVariable failed: %v", err)
		}
	})

	t.Run("Test DeleteVariable", func(t *testing.T) {
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

		_, err := srv.DeleteVariable("<VARIABLE_ID>")
		if err != nil {
			t.Errorf("Method DeleteVariable failed: %v", err)
		}
	})
}
