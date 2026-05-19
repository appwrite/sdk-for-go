package advisor

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v4/client"
)

func TestAdvisor(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test ListReports", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "reports": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "appId": "5e5ea5c16897e",
            "type": "lighthouse",
            "title": "Lighthouse audit for https://appwrite.io/",
            "summary": "Performance score 78. 4 opportunities found.",
            "targetType": "urls",
            "target": "https://appwrite.io/",
            "categories": [],
            "insights": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "reportId": "5e5ea5c16897e",
                    "type": "tablesDBIndex",
                    "severity": "warning",
                    "status": "active",
                    "resourceType": "databases",
                    "resourceId": "main",
                    "parentResourceType": "tables",
                    "parentResourceId": "orders",
                    "title": "Missing index on collection orders",
                    "summary": "Queries against `orders.status` are scanning the full collection.",
                    "ctas": [
                        {
                            "label": "Create missing index",
                            "service": "tablesDB",
                            "method": "createIndex",
                            "params": {}
                        }
                    ]
                }
            ]
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

		_, err := srv.ListReports()
		if err != nil {
			t.Errorf("Method ListReports failed: %v", err)
		}
	})

	t.Run("Test GetReport", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "appId": "5e5ea5c16897e",
    "type": "lighthouse",
    "title": "Lighthouse audit for https://appwrite.io/",
    "summary": "Performance score 78. 4 opportunities found.",
    "targetType": "urls",
    "target": "https://appwrite.io/",
    "categories": [],
    "insights": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "reportId": "5e5ea5c16897e",
            "type": "tablesDBIndex",
            "severity": "warning",
            "status": "active",
            "resourceType": "databases",
            "resourceId": "main",
            "parentResourceType": "tables",
            "parentResourceId": "orders",
            "title": "Missing index on collection orders",
            "summary": "Queries against `orders.status` are scanning the full collection.",
            "ctas": [
                {
                    "label": "Create missing index",
                    "service": "tablesDB",
                    "method": "createIndex",
                    "params": {}
                }
            ]
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

		_, err := srv.GetReport("<REPORT_ID>")
		if err != nil {
			t.Errorf("Method GetReport failed: %v", err)
		}
	})

	t.Run("Test DeleteReport", func(t *testing.T) {
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

		_, err := srv.DeleteReport("<REPORT_ID>")
		if err != nil {
			t.Errorf("Method DeleteReport failed: %v", err)
		}
	})

	t.Run("Test ListInsights", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "insights": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "reportId": "5e5ea5c16897e",
            "type": "tablesDBIndex",
            "severity": "warning",
            "status": "active",
            "resourceType": "databases",
            "resourceId": "main",
            "parentResourceType": "tables",
            "parentResourceId": "orders",
            "title": "Missing index on collection orders",
            "summary": "Queries against `orders.status` are scanning the full collection.",
            "ctas": [
                {
                    "label": "Create missing index",
                    "service": "tablesDB",
                    "method": "createIndex",
                    "params": {}
                }
            ]
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

		_, err := srv.ListInsights("<REPORT_ID>")
		if err != nil {
			t.Errorf("Method ListInsights failed: %v", err)
		}
	})

	t.Run("Test GetInsight", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "reportId": "5e5ea5c16897e",
    "type": "tablesDBIndex",
    "severity": "warning",
    "status": "active",
    "resourceType": "databases",
    "resourceId": "main",
    "parentResourceType": "tables",
    "parentResourceId": "orders",
    "title": "Missing index on collection orders",
    "summary": "Queries against `orders.status` are scanning the full collection.",
    "ctas": [
        {
            "label": "Create missing index",
            "service": "tablesDB",
            "method": "createIndex",
            "params": {}
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

		_, err := srv.GetInsight("<REPORT_ID>", "<INSIGHT_ID>")
		if err != nil {
			t.Errorf("Method GetInsight failed: %v", err)
		}
	})
}
