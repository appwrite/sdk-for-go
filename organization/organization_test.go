package organization

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestOrganization(t *testing.T) {
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
    "name": "VIP",
    "total": 7,
    "prefs": {
    },
    "budgetAlerts": [],
    "billingPlan": "tier-1",
    "billingPlanId": "tier-1",
    "billingPlanDetails": {
        "$id": "tier-0",
        "name": "Hobby",
        "desc": "Hobby plan",
        "order": 0,
        "price": 25,
        "trial": 14,
        "bandwidth": 25,
        "storage": 25,
        "imageTransformations": 100,
        "screenshotsGenerated": 50,
        "webhooks": 25,
        "wafRules": 2,
        "projects": 2,
        "platforms": 3,
        "users": 25,
        "teams": 25,
        "databases": 25,
        "databasesReads": 500000,
        "databasesWrites": 250000,
        "databasesBatchSize": 100,
        "buckets": 25,
        "fileSize": 25,
        "functions": 25,
        "sites": 1,
        "executions": 25,
        "executionsRetentionCount": 10000,
        "GBHours": 100,
        "realtime": 25,
        "realtimeMessages": 100000,
        "messages": 1000,
        "topics": 1,
        "authPhone": 10,
        "domains": 5,
        "usageLogs": 30,
        "projectInactivityDays": 7,
        "alertLimit": 80,
        "usage": {
            "bandwidth": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "executions": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "realtime": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "realtimeMessages": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "storage": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "users": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "GBHours": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "imageTransformations": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            }
        },
        "addons": {
        },
        "budgetCapEnabled": true,
        "customSmtp": true,
        "emailBranding": true,
        "requiresPaymentMethod": true,
        "requiresBillingAddress": true,
        "isAvailable": true,
        "selfService": true,
        "premiumSupport": true,
        "budgeting": true,
        "supportsMockNumbers": true,
        "supportsOrganizationRoles": true,
        "supportsCredits": true,
        "supportsDisposableEmailValidation": true,
        "supportsCanonicalEmailValidation": true,
        "supportsFreeEmailValidation": true,
        "supportsCorporateEmailValidation": true,
        "supportsProjectSpecificRoles": true,
        "usagePerProject": true,
        "supportedAddons": {
            "baa": true,
            "premiumGeoDB": true,
            "premiumGeoDBOrg": true
        },
        "deploymentSize": 30,
        "buildSize": 2000,
        "databasesAllowEncrypt": true,
        "group": "pro"
    },
    "billingEmail": "billing@org.example",
    "billingStartDate": "2020-10-15T06:38:00.000+00:00",
    "billingCurrentInvoiceDate": "2020-10-15T06:38:00.000+00:00",
    "billingNextInvoiceDate": "2020-10-15T06:38:00.000+00:00",
    "billingTrialDays": 14,
    "billingAggregationId": "adbc3de4rddfsd",
    "billingInvoiceId": "adbc3de4rddfsd",
    "paymentMethodId": "adbc3de4rddfsd",
    "status": "active",
    "markedForDeletion": true,
    "platform": "imagine",
    "projects": []
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

	t.Run("Test Update", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "VIP",
    "total": 7,
    "prefs": {
    },
    "budgetAlerts": [],
    "billingPlan": "tier-1",
    "billingPlanId": "tier-1",
    "billingPlanDetails": {
        "$id": "tier-0",
        "name": "Hobby",
        "desc": "Hobby plan",
        "order": 0,
        "price": 25,
        "trial": 14,
        "bandwidth": 25,
        "storage": 25,
        "imageTransformations": 100,
        "screenshotsGenerated": 50,
        "webhooks": 25,
        "wafRules": 2,
        "projects": 2,
        "platforms": 3,
        "users": 25,
        "teams": 25,
        "databases": 25,
        "databasesReads": 500000,
        "databasesWrites": 250000,
        "databasesBatchSize": 100,
        "buckets": 25,
        "fileSize": 25,
        "functions": 25,
        "sites": 1,
        "executions": 25,
        "executionsRetentionCount": 10000,
        "GBHours": 100,
        "realtime": 25,
        "realtimeMessages": 100000,
        "messages": 1000,
        "topics": 1,
        "authPhone": 10,
        "domains": 5,
        "usageLogs": 30,
        "projectInactivityDays": 7,
        "alertLimit": 80,
        "usage": {
            "bandwidth": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "executions": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "realtime": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "realtimeMessages": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "storage": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "users": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "GBHours": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            },
            "imageTransformations": {
                "name": "string",
                "unit": "GB",
                "currency": "USD",
                "price": 5,
                "value": 25,
                "invoiceDesc": "string"
            }
        },
        "addons": {
        },
        "budgetCapEnabled": true,
        "customSmtp": true,
        "emailBranding": true,
        "requiresPaymentMethod": true,
        "requiresBillingAddress": true,
        "isAvailable": true,
        "selfService": true,
        "premiumSupport": true,
        "budgeting": true,
        "supportsMockNumbers": true,
        "supportsOrganizationRoles": true,
        "supportsCredits": true,
        "supportsDisposableEmailValidation": true,
        "supportsCanonicalEmailValidation": true,
        "supportsFreeEmailValidation": true,
        "supportsCorporateEmailValidation": true,
        "supportsProjectSpecificRoles": true,
        "usagePerProject": true,
        "supportedAddons": {
            "baa": true,
            "premiumGeoDB": true,
            "premiumGeoDBOrg": true
        },
        "deploymentSize": 30,
        "buildSize": 2000,
        "databasesAllowEncrypt": true,
        "group": "pro"
    },
    "billingEmail": "billing@org.example",
    "billingStartDate": "2020-10-15T06:38:00.000+00:00",
    "billingCurrentInvoiceDate": "2020-10-15T06:38:00.000+00:00",
    "billingNextInvoiceDate": "2020-10-15T06:38:00.000+00:00",
    "billingTrialDays": 14,
    "billingAggregationId": "adbc3de4rddfsd",
    "billingInvoiceId": "adbc3de4rddfsd",
    "paymentMethodId": "adbc3de4rddfsd",
    "status": "active",
    "markedForDeletion": true,
    "platform": "imagine",
    "projects": []
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

		_, err := srv.Update("<NAME>")
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

		_, err := srv.Delete()
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

		_, err := srv.ListInstallations()
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
    "authorizationDetails": {},
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

		_, err := srv.CreateInstallation("<APP_ID>")
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

		_, err := srv.GetInstallation("<INSTALLATION_ID>")
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
    "authorizationDetails": {},
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

		_, err := srv.UpdateInstallation("<INSTALLATION_ID>")
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

		_, err := srv.DeleteInstallation("<INSTALLATION_ID>")
		if err != nil {
			t.Errorf("Method DeleteInstallation failed: %v", err)
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

		_, err := srv.ListMemberships()
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

		_, err := srv.CreateMembership([]string{})
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

		_, err := srv.GetMembership("<MEMBERSHIP_ID>")
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

		_, err := srv.UpdateMembership("<MEMBERSHIP_ID>", []string{})
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

		_, err := srv.DeleteMembership("<MEMBERSHIP_ID>")
		if err != nil {
			t.Errorf("Method DeleteMembership failed: %v", err)
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
            "region": "fra",
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
            "smtpPassword": "smtp-password",
            "smtpSecure": "tls",
            "pingCount": 1,
            "pingedAt": "2020-10-15T06:38:00.000+00:00",
            "labels": [],
            "status": "active",
            "onboarding": {},
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
            "blocks": [
                {
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "resourceType": "project",
                    "resourceId": "5e5ea5c16897e",
                    "mode": "readOnly",
                    "projectName": "My Project",
                    "region": "fra",
                    "organizationName": "Acme Inc.",
                    "organizationId": "5e5ea5c16897e",
                    "billingPlan": "pro"
                }
            ],
            "consoleAccessedAt": "2020-10-15T06:38:00.000+00:00",
            "wafEnabled": true
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
    "region": "fra",
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
    "smtpPassword": "smtp-password",
    "smtpSecure": "tls",
    "pingCount": 1,
    "pingedAt": "2020-10-15T06:38:00.000+00:00",
    "labels": [],
    "status": "active",
    "onboarding": {},
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
    "blocks": [
        {
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "resourceType": "project",
            "resourceId": "5e5ea5c16897e",
            "mode": "readOnly",
            "projectName": "My Project",
            "region": "fra",
            "organizationName": "Acme Inc.",
            "organizationId": "5e5ea5c16897e",
            "billingPlan": "pro"
        }
    ],
    "consoleAccessedAt": "2020-10-15T06:38:00.000+00:00",
    "wafEnabled": true
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
    "region": "fra",
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
    "smtpPassword": "smtp-password",
    "smtpSecure": "tls",
    "pingCount": 1,
    "pingedAt": "2020-10-15T06:38:00.000+00:00",
    "labels": [],
    "status": "active",
    "onboarding": {},
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
    "blocks": [
        {
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "resourceType": "project",
            "resourceId": "5e5ea5c16897e",
            "mode": "readOnly",
            "projectName": "My Project",
            "region": "fra",
            "organizationName": "Acme Inc.",
            "organizationId": "5e5ea5c16897e",
            "billingPlan": "pro"
        }
    ],
    "consoleAccessedAt": "2020-10-15T06:38:00.000+00:00",
    "wafEnabled": true
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
    "region": "fra",
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
    "smtpPassword": "smtp-password",
    "smtpSecure": "tls",
    "pingCount": 1,
    "pingedAt": "2020-10-15T06:38:00.000+00:00",
    "labels": [],
    "status": "active",
    "onboarding": {},
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
    "blocks": [
        {
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "resourceType": "project",
            "resourceId": "5e5ea5c16897e",
            "mode": "readOnly",
            "projectName": "My Project",
            "region": "fra",
            "organizationName": "Acme Inc.",
            "organizationId": "5e5ea5c16897e",
            "billingPlan": "pro"
        }
    ],
    "consoleAccessedAt": "2020-10-15T06:38:00.000+00:00",
    "wafEnabled": true
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
