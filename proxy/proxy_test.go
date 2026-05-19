package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v4/client"
)

func TestProxy(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test ListRules", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "rules": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "domain": "appwrite.company.com",
            "type": "deployment",
            "trigger": "manual",
            "redirectUrl": "https://appwrite.io/docs",
            "redirectStatusCode": 301,
            "deploymentId": "n3u9feiwmf",
            "deploymentResourceId": "n3u9feiwmf",
            "deploymentVcsProviderBranch": "main",
            "status": "verified",
            "logs": "Verification of DNS records failed with DNS resolver 8.8.8.8. Domain stage.myapp.com does not have DNS record.",
            "renewAt": "datetime"
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

		_, err := srv.ListRules()
		if err != nil {
			t.Errorf("Method ListRules failed: %v", err)
		}
	})

	t.Run("Test CreateAPIRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "appwrite.company.com",
    "type": "deployment",
    "trigger": "manual",
    "redirectUrl": "https://appwrite.io/docs",
    "redirectStatusCode": 301,
    "deploymentId": "n3u9feiwmf",
    "deploymentResourceId": "n3u9feiwmf",
    "deploymentVcsProviderBranch": "main",
    "status": "verified",
    "logs": "Verification of DNS records failed with DNS resolver 8.8.8.8. Domain stage.myapp.com does not have DNS record.",
    "renewAt": "datetime"
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

		_, err := srv.CreateAPIRule("")
		if err != nil {
			t.Errorf("Method CreateAPIRule failed: %v", err)
		}
	})

	t.Run("Test CreateFunctionRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "appwrite.company.com",
    "type": "deployment",
    "trigger": "manual",
    "redirectUrl": "https://appwrite.io/docs",
    "redirectStatusCode": 301,
    "deploymentId": "n3u9feiwmf",
    "deploymentResourceId": "n3u9feiwmf",
    "deploymentVcsProviderBranch": "main",
    "status": "verified",
    "logs": "Verification of DNS records failed with DNS resolver 8.8.8.8. Domain stage.myapp.com does not have DNS record.",
    "renewAt": "datetime"
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

		_, err := srv.CreateFunctionRule("", "<FUNCTION_ID>")
		if err != nil {
			t.Errorf("Method CreateFunctionRule failed: %v", err)
		}
	})

	t.Run("Test CreateRedirectRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "appwrite.company.com",
    "type": "deployment",
    "trigger": "manual",
    "redirectUrl": "https://appwrite.io/docs",
    "redirectStatusCode": 301,
    "deploymentId": "n3u9feiwmf",
    "deploymentResourceId": "n3u9feiwmf",
    "deploymentVcsProviderBranch": "main",
    "status": "verified",
    "logs": "Verification of DNS records failed with DNS resolver 8.8.8.8. Domain stage.myapp.com does not have DNS record.",
    "renewAt": "datetime"
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

		_, err := srv.CreateRedirectRule("", "https://example.com", "301", "<RESOURCE_ID>", "site")
		if err != nil {
			t.Errorf("Method CreateRedirectRule failed: %v", err)
		}
	})

	t.Run("Test CreateSiteRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "appwrite.company.com",
    "type": "deployment",
    "trigger": "manual",
    "redirectUrl": "https://appwrite.io/docs",
    "redirectStatusCode": 301,
    "deploymentId": "n3u9feiwmf",
    "deploymentResourceId": "n3u9feiwmf",
    "deploymentVcsProviderBranch": "main",
    "status": "verified",
    "logs": "Verification of DNS records failed with DNS resolver 8.8.8.8. Domain stage.myapp.com does not have DNS record.",
    "renewAt": "datetime"
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

		_, err := srv.CreateSiteRule("", "<SITE_ID>")
		if err != nil {
			t.Errorf("Method CreateSiteRule failed: %v", err)
		}
	})

	t.Run("Test GetRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "appwrite.company.com",
    "type": "deployment",
    "trigger": "manual",
    "redirectUrl": "https://appwrite.io/docs",
    "redirectStatusCode": 301,
    "deploymentId": "n3u9feiwmf",
    "deploymentResourceId": "n3u9feiwmf",
    "deploymentVcsProviderBranch": "main",
    "status": "verified",
    "logs": "Verification of DNS records failed with DNS resolver 8.8.8.8. Domain stage.myapp.com does not have DNS record.",
    "renewAt": "datetime"
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

		_, err := srv.GetRule("<RULE_ID>")
		if err != nil {
			t.Errorf("Method GetRule failed: %v", err)
		}
	})

	t.Run("Test DeleteRule", func(t *testing.T) {
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

		_, err := srv.DeleteRule("<RULE_ID>")
		if err != nil {
			t.Errorf("Method DeleteRule failed: %v", err)
		}
	})

	t.Run("Test UpdateRuleStatus", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "domain": "appwrite.company.com",
    "type": "deployment",
    "trigger": "manual",
    "redirectUrl": "https://appwrite.io/docs",
    "redirectStatusCode": 301,
    "deploymentId": "n3u9feiwmf",
    "deploymentResourceId": "n3u9feiwmf",
    "deploymentVcsProviderBranch": "main",
    "status": "verified",
    "logs": "Verification of DNS records failed with DNS resolver 8.8.8.8. Domain stage.myapp.com does not have DNS record.",
    "renewAt": "datetime"
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

		_, err := srv.UpdateRuleStatus("<RULE_ID>")
		if err != nil {
			t.Errorf("Method UpdateRuleStatus failed: %v", err)
		}
	})
}
