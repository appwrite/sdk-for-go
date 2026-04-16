package messaging

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v3/client"
)

func TestMessaging(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test ListMessages", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "messages": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "providerType": "email",
            "topics": [],
            "users": [],
            "targets": [],
            "deliveredTotal": 1,
            "data": {},
            "status": "processing"
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

		_, err := srv.ListMessages()
		if err != nil {
			t.Errorf("Method ListMessages failed: %v", err)
		}
	})

	t.Run("Test CreateEmail", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "providerType": "email",
    "topics": [],
    "users": [],
    "targets": [],
    "deliveredTotal": 1,
    "data": {},
    "status": "processing"
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

		_, err := srv.CreateEmail("<MESSAGE_ID>", "<SUBJECT>", "<CONTENT>")
		if err != nil {
			t.Errorf("Method CreateEmail failed: %v", err)
		}
	})

	t.Run("Test UpdateEmail", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "providerType": "email",
    "topics": [],
    "users": [],
    "targets": [],
    "deliveredTotal": 1,
    "data": {},
    "status": "processing"
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

		_, err := srv.UpdateEmail("<MESSAGE_ID>")
		if err != nil {
			t.Errorf("Method UpdateEmail failed: %v", err)
		}
	})

	t.Run("Test CreatePush", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "providerType": "email",
    "topics": [],
    "users": [],
    "targets": [],
    "deliveredTotal": 1,
    "data": {},
    "status": "processing"
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

		_, err := srv.CreatePush("<MESSAGE_ID>")
		if err != nil {
			t.Errorf("Method CreatePush failed: %v", err)
		}
	})

	t.Run("Test UpdatePush", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "providerType": "email",
    "topics": [],
    "users": [],
    "targets": [],
    "deliveredTotal": 1,
    "data": {},
    "status": "processing"
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

		_, err := srv.UpdatePush("<MESSAGE_ID>")
		if err != nil {
			t.Errorf("Method UpdatePush failed: %v", err)
		}
	})

	t.Run("Test CreateSms", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "providerType": "email",
    "topics": [],
    "users": [],
    "targets": [],
    "deliveredTotal": 1,
    "data": {},
    "status": "processing"
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

		_, err := srv.CreateSms("<MESSAGE_ID>", "<CONTENT>")
		if err != nil {
			t.Errorf("Method CreateSms failed: %v", err)
		}
	})

	t.Run("Test CreateSMS", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "providerType": "email",
    "topics": [],
    "users": [],
    "targets": [],
    "deliveredTotal": 1,
    "data": {},
    "status": "processing"
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

		_, err := srv.CreateSMS("<MESSAGE_ID>", "<CONTENT>")
		if err != nil {
			t.Errorf("Method CreateSMS failed: %v", err)
		}
	})

	t.Run("Test UpdateSms", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "providerType": "email",
    "topics": [],
    "users": [],
    "targets": [],
    "deliveredTotal": 1,
    "data": {},
    "status": "processing"
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

		_, err := srv.UpdateSms("<MESSAGE_ID>")
		if err != nil {
			t.Errorf("Method UpdateSms failed: %v", err)
		}
	})

	t.Run("Test UpdateSMS", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "providerType": "email",
    "topics": [],
    "users": [],
    "targets": [],
    "deliveredTotal": 1,
    "data": {},
    "status": "processing"
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

		_, err := srv.UpdateSMS("<MESSAGE_ID>")
		if err != nil {
			t.Errorf("Method UpdateSMS failed: %v", err)
		}
	})

	t.Run("Test GetMessage", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "providerType": "email",
    "topics": [],
    "users": [],
    "targets": [],
    "deliveredTotal": 1,
    "data": {},
    "status": "processing"
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

		_, err := srv.GetMessage("<MESSAGE_ID>")
		if err != nil {
			t.Errorf("Method GetMessage failed: %v", err)
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

		_, err := srv.Delete("<MESSAGE_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})

	t.Run("Test ListMessageLogs", func(t *testing.T) {
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

		_, err := srv.ListMessageLogs("<MESSAGE_ID>")
		if err != nil {
			t.Errorf("Method ListMessageLogs failed: %v", err)
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

		_, err := srv.ListTargets("<MESSAGE_ID>")
		if err != nil {
			t.Errorf("Method ListTargets failed: %v", err)
		}
	})

	t.Run("Test ListProviders", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "providers": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Mailgun",
            "provider": "mailgun",
            "enabled": true,
            "type": "sms",
            "credentials": {}
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

		_, err := srv.ListProviders()
		if err != nil {
			t.Errorf("Method ListProviders failed: %v", err)
		}
	})

	t.Run("Test CreateApnsProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateApnsProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateApnsProvider failed: %v", err)
		}
	})

	t.Run("Test CreateAPNSProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateAPNSProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateAPNSProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateApnsProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateApnsProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateApnsProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateAPNSProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateAPNSProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateAPNSProvider failed: %v", err)
		}
	})

	t.Run("Test CreateFcmProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateFcmProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateFcmProvider failed: %v", err)
		}
	})

	t.Run("Test CreateFCMProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateFCMProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateFCMProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateFcmProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateFcmProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateFcmProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateFCMProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateFCMProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateFCMProvider failed: %v", err)
		}
	})

	t.Run("Test CreateMailgunProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateMailgunProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateMailgunProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateMailgunProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateMailgunProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateMailgunProvider failed: %v", err)
		}
	})

	t.Run("Test CreateMsg91Provider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateMsg91Provider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateMsg91Provider failed: %v", err)
		}
	})

	t.Run("Test UpdateMsg91Provider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateMsg91Provider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateMsg91Provider failed: %v", err)
		}
	})

	t.Run("Test CreateResendProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateResendProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateResendProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateResendProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateResendProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateResendProvider failed: %v", err)
		}
	})

	t.Run("Test CreateSendgridProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateSendgridProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateSendgridProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateSendgridProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateSendgridProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateSendgridProvider failed: %v", err)
		}
	})

	t.Run("Test CreateSmtpProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateSmtpProvider("<PROVIDER_ID>", "<NAME>", "<HOST>")
		if err != nil {
			t.Errorf("Method CreateSmtpProvider failed: %v", err)
		}
	})

	t.Run("Test CreateSMTPProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateSMTPProvider("<PROVIDER_ID>", "<NAME>", "<HOST>")
		if err != nil {
			t.Errorf("Method CreateSMTPProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateSmtpProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateSmtpProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateSmtpProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateSMTPProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateSMTPProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateSMTPProvider failed: %v", err)
		}
	})

	t.Run("Test CreateTelesignProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateTelesignProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateTelesignProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateTelesignProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateTelesignProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateTelesignProvider failed: %v", err)
		}
	})

	t.Run("Test CreateTextmagicProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateTextmagicProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateTextmagicProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateTextmagicProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateTextmagicProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateTextmagicProvider failed: %v", err)
		}
	})

	t.Run("Test CreateTwilioProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateTwilioProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateTwilioProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateTwilioProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateTwilioProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateTwilioProvider failed: %v", err)
		}
	})

	t.Run("Test CreateVonageProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.CreateVonageProvider("<PROVIDER_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateVonageProvider failed: %v", err)
		}
	})

	t.Run("Test UpdateVonageProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.UpdateVonageProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method UpdateVonageProvider failed: %v", err)
		}
	})

	t.Run("Test GetProvider", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Mailgun",
    "provider": "mailgun",
    "enabled": true,
    "type": "sms",
    "credentials": {}
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

		_, err := srv.GetProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method GetProvider failed: %v", err)
		}
	})

	t.Run("Test DeleteProvider", func(t *testing.T) {
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

		_, err := srv.DeleteProvider("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method DeleteProvider failed: %v", err)
		}
	})

	t.Run("Test ListProviderLogs", func(t *testing.T) {
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

		_, err := srv.ListProviderLogs("<PROVIDER_ID>")
		if err != nil {
			t.Errorf("Method ListProviderLogs failed: %v", err)
		}
	})

	t.Run("Test ListSubscriberLogs", func(t *testing.T) {
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

		_, err := srv.ListSubscriberLogs("<SUBSCRIBER_ID>")
		if err != nil {
			t.Errorf("Method ListSubscriberLogs failed: %v", err)
		}
	})

	t.Run("Test ListTopics", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "topics": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "events",
            "emailTotal": 100,
            "smsTotal": 100,
            "pushTotal": 100,
            "subscribe": []
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

		_, err := srv.ListTopics()
		if err != nil {
			t.Errorf("Method ListTopics failed: %v", err)
		}
	})

	t.Run("Test CreateTopic", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "events",
    "emailTotal": 100,
    "smsTotal": 100,
    "pushTotal": 100,
    "subscribe": []
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

		_, err := srv.CreateTopic("<TOPIC_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateTopic failed: %v", err)
		}
	})

	t.Run("Test GetTopic", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "events",
    "emailTotal": 100,
    "smsTotal": 100,
    "pushTotal": 100,
    "subscribe": []
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

		_, err := srv.GetTopic("<TOPIC_ID>")
		if err != nil {
			t.Errorf("Method GetTopic failed: %v", err)
		}
	})

	t.Run("Test UpdateTopic", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "events",
    "emailTotal": 100,
    "smsTotal": 100,
    "pushTotal": 100,
    "subscribe": []
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

		_, err := srv.UpdateTopic("<TOPIC_ID>")
		if err != nil {
			t.Errorf("Method UpdateTopic failed: %v", err)
		}
	})

	t.Run("Test DeleteTopic", func(t *testing.T) {
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

		_, err := srv.DeleteTopic("<TOPIC_ID>")
		if err != nil {
			t.Errorf("Method DeleteTopic failed: %v", err)
		}
	})

	t.Run("Test ListTopicLogs", func(t *testing.T) {
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

		_, err := srv.ListTopicLogs("<TOPIC_ID>")
		if err != nil {
			t.Errorf("Method ListTopicLogs failed: %v", err)
		}
	})

	t.Run("Test ListSubscribers", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "subscribers": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "targetId": "259125845563242502",
            "target": {
                "$id": "259125845563242502",
                "$createdAt": "2020-10-15T06:38:00.000+00:00",
                "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                "name": "Apple iPhone 12",
                "userId": "259125845563242502",
                "providerType": "email",
                "identifier": "token",
                "expired": true
            },
            "userId": "5e5ea5c16897e",
            "userName": "Aegon Targaryen",
            "topicId": "259125845563242502",
            "providerType": "email"
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

		_, err := srv.ListSubscribers("<TOPIC_ID>")
		if err != nil {
			t.Errorf("Method ListSubscribers failed: %v", err)
		}
	})

	t.Run("Test CreateSubscriber", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "targetId": "259125845563242502",
    "target": {
        "$id": "259125845563242502",
        "$createdAt": "2020-10-15T06:38:00.000+00:00",
        "$updatedAt": "2020-10-15T06:38:00.000+00:00",
        "name": "Apple iPhone 12",
        "userId": "259125845563242502",
        "providerType": "email",
        "identifier": "token",
        "expired": true
    },
    "userId": "5e5ea5c16897e",
    "userName": "Aegon Targaryen",
    "topicId": "259125845563242502",
    "providerType": "email"
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

		_, err := srv.CreateSubscriber("<TOPIC_ID>", "<SUBSCRIBER_ID>", "<TARGET_ID>")
		if err != nil {
			t.Errorf("Method CreateSubscriber failed: %v", err)
		}
	})

	t.Run("Test GetSubscriber", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "targetId": "259125845563242502",
    "target": {
        "$id": "259125845563242502",
        "$createdAt": "2020-10-15T06:38:00.000+00:00",
        "$updatedAt": "2020-10-15T06:38:00.000+00:00",
        "name": "Apple iPhone 12",
        "userId": "259125845563242502",
        "providerType": "email",
        "identifier": "token",
        "expired": true
    },
    "userId": "5e5ea5c16897e",
    "userName": "Aegon Targaryen",
    "topicId": "259125845563242502",
    "providerType": "email"
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

		_, err := srv.GetSubscriber("<TOPIC_ID>", "<SUBSCRIBER_ID>")
		if err != nil {
			t.Errorf("Method GetSubscriber failed: %v", err)
		}
	})

	t.Run("Test DeleteSubscriber", func(t *testing.T) {
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

		_, err := srv.DeleteSubscriber("<TOPIC_ID>", "<SUBSCRIBER_ID>")
		if err != nil {
			t.Errorf("Method DeleteSubscriber failed: %v", err)
		}
	})
}
