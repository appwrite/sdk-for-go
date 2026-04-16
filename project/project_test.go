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
            "phone": "+1612842323",
            "otp": "123456"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "Bpw_g9c2TGXxfgLshDbSaL8tsCcqgczQ",
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
    "smtpReplyTo": "support@appwrite.io",
    "smtpHost": "mail.appwrite.io",
    "smtpPort": 25,
    "smtpUsername": "emailuser",
    "smtpPassword": "securepassword",
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
            "resourceId": "5e5ea5c16897e"
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

	t.Run("Test UpdateProtocolStatus", func(t *testing.T) {
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
            "phone": "+1612842323",
            "otp": "123456"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "Bpw_g9c2TGXxfgLshDbSaL8tsCcqgczQ",
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
    "smtpReplyTo": "support@appwrite.io",
    "smtpHost": "mail.appwrite.io",
    "smtpPort": 25,
    "smtpUsername": "emailuser",
    "smtpPassword": "securepassword",
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
            "resourceId": "5e5ea5c16897e"
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

		_, err := srv.UpdateProtocolStatus("rest", true)
		if err != nil {
			t.Errorf("Method UpdateProtocolStatus failed: %v", err)
		}
	})

	t.Run("Test UpdateServiceStatus", func(t *testing.T) {
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
            "phone": "+1612842323",
            "otp": "123456"
        }
    ],
    "authSessionAlerts": true,
    "authMembershipsUserName": true,
    "authMembershipsUserEmail": true,
    "authMembershipsMfa": true,
    "authInvalidateSessions": true,
    "oAuthProviders": [
        {
            "key": "github",
            "name": "GitHub",
            "appId": "259125845563242502",
            "secret": "Bpw_g9c2TGXxfgLshDbSaL8tsCcqgczQ",
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
    "smtpReplyTo": "support@appwrite.io",
    "smtpHost": "mail.appwrite.io",
    "smtpPort": 25,
    "smtpUsername": "emailuser",
    "smtpPassword": "securepassword",
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
            "resourceId": "5e5ea5c16897e"
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

		_, err := srv.UpdateServiceStatus("account", true)
		if err != nil {
			t.Errorf("Method UpdateServiceStatus failed: %v", err)
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
