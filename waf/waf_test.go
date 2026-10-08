package waf

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestWaf(t *testing.T) {
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
            "$id": "wafRule1",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "name": "Block anonymous POST traffic",
            "description": "Blocks anonymous POST calls to /v1/graphql",
            "teamId": "5e5ea5c16897e",
            "projectId": "cloudConsole",
            "resourceType": "functions",
            "resourceId": "functionId",
            "action": "deny",
            "priority": 100,
            "enabled": true,
            "conditions": [],
            "config": {}
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

	t.Run("Test CreateBypassRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {}
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

		_, err := srv.CreateBypassRule("<RULE_ID>", "api", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateBypassRule failed: %v", err)
		}
	})

	t.Run("Test UpdateBypassRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {}
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

		_, err := srv.UpdateBypassRule("<RULE_ID>")
		if err != nil {
			t.Errorf("Method UpdateBypassRule failed: %v", err)
		}
	})

	t.Run("Test CreateChallengeRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {},
    "challengeType": "compute",
    "difficulty": 3,
    "ttl": 1800
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

		_, err := srv.CreateChallengeRule("<RULE_ID>", "api", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateChallengeRule failed: %v", err)
		}
	})

	t.Run("Test UpdateChallengeRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {},
    "challengeType": "compute",
    "difficulty": 3,
    "ttl": 1800
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

		_, err := srv.UpdateChallengeRule("<RULE_ID>")
		if err != nil {
			t.Errorf("Method UpdateChallengeRule failed: %v", err)
		}
	})

	t.Run("Test CreateDenyRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {}
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

		_, err := srv.CreateDenyRule("<RULE_ID>", "api", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateDenyRule failed: %v", err)
		}
	})

	t.Run("Test UpdateDenyRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {}
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

		_, err := srv.UpdateDenyRule("<RULE_ID>")
		if err != nil {
			t.Errorf("Method UpdateDenyRule failed: %v", err)
		}
	})

	t.Run("Test CreateRateLimitRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {},
    "limit": 1000,
    "interval": 60,
    "key": "userId",
    "strategy": "slidingWindow",
    "maxBucketSize": 200
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

		_, err := srv.CreateRateLimitRule("<RULE_ID>", "api", "<NAME>", 1, 1)
		if err != nil {
			t.Errorf("Method CreateRateLimitRule failed: %v", err)
		}
	})

	t.Run("Test UpdateRateLimitRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {},
    "limit": 1000,
    "interval": 60,
    "key": "userId",
    "strategy": "slidingWindow",
    "maxBucketSize": 200
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

		_, err := srv.UpdateRateLimitRule("<RULE_ID>")
		if err != nil {
			t.Errorf("Method UpdateRateLimitRule failed: %v", err)
		}
	})

	t.Run("Test CreateRedirectRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {},
    "location": "/maintenance",
    "statusCode": 301
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

		_, err := srv.CreateRedirectRule("<RULE_ID>", "api", "<NAME>", "<LOCATION>", 1)
		if err != nil {
			t.Errorf("Method CreateRedirectRule failed: %v", err)
		}
	})

	t.Run("Test UpdateRedirectRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {},
    "location": "/maintenance",
    "statusCode": 301
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

		_, err := srv.UpdateRedirectRule("<RULE_ID>")
		if err != nil {
			t.Errorf("Method UpdateRedirectRule failed: %v", err)
		}
	})

	t.Run("Test GetRule", func(t *testing.T) {
		mockResponse := `
{
    "$id": "wafRule1",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "name": "Block anonymous POST traffic",
    "description": "Blocks anonymous POST calls to /v1/graphql",
    "teamId": "5e5ea5c16897e",
    "projectId": "cloudConsole",
    "resourceType": "functions",
    "resourceId": "functionId",
    "action": "deny",
    "priority": 100,
    "enabled": true,
    "conditions": [],
    "config": {}
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
}
