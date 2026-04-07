package tablesdb

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v2/client"
)

func TestTablesDB(t *testing.T) {
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
    "databases": [
        {
            "$id": "5e5ea5c16897e",
            "name": "My Database",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "enabled": true,
            "type": "legacy",
            "policies": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "key": "index1",
                    "type": "primary",
                    "status": "available",
                    "error": "string",
                    "attributes": [],
                    "lengths": []
                }
            ],
            "archives": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "$permissions": [],
                    "databaseId": "5e5ea5c16897e",
                    "name": "My Collection",
                    "enabled": true,
                    "documentSecurity": true,
                    "attributes": [],
                    "indexes": [
                        {
                            "$id": "5e5ea5c16897e",
                            "$createdAt": "2020-10-15T06:38:00.000+00:00",
                            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                            "key": "index1",
                            "type": "primary",
                            "status": "available",
                            "error": "string",
                            "attributes": [],
                            "lengths": []
                        }
                    ],
                    "bytesMax": 65535,
                    "bytesUsed": 1500
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

		_, err := srv.List()
		if err != nil {
			t.Errorf("Method List failed: %v", err)
		}
	})

	t.Run("Test Create", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "name": "My Database",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "enabled": true,
    "type": "legacy",
    "policies": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "index1",
            "type": "primary",
            "status": "available",
            "error": "string",
            "attributes": [],
            "lengths": []
        }
    ],
    "archives": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": [],
            "databaseId": "5e5ea5c16897e",
            "name": "My Collection",
            "enabled": true,
            "documentSecurity": true,
            "attributes": [],
            "indexes": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "key": "index1",
                    "type": "primary",
                    "status": "available",
                    "error": "string",
                    "attributes": [],
                    "lengths": []
                }
            ],
            "bytesMax": 65535,
            "bytesUsed": 1500
        }
    ]
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

		_, err := srv.Create("<DATABASE_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method Create failed: %v", err)
		}
	})

	t.Run("Test ListTransactions", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "transactions": [
        {
            "$id": "259125845563242502",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "status": "pending",
            "operations": 5,
            "expiresAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.ListTransactions()
		if err != nil {
			t.Errorf("Method ListTransactions failed: %v", err)
		}
	})

	t.Run("Test CreateTransaction", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "operations": 5,
    "expiresAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateTransaction()
		if err != nil {
			t.Errorf("Method CreateTransaction failed: %v", err)
		}
	})

	t.Run("Test GetTransaction", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "operations": 5,
    "expiresAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.GetTransaction("<TRANSACTION_ID>")
		if err != nil {
			t.Errorf("Method GetTransaction failed: %v", err)
		}
	})

	t.Run("Test UpdateTransaction", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "operations": 5,
    "expiresAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateTransaction("<TRANSACTION_ID>")
		if err != nil {
			t.Errorf("Method UpdateTransaction failed: %v", err)
		}
	})

	t.Run("Test DeleteTransaction", func(t *testing.T) {
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

		_, err := srv.DeleteTransaction("<TRANSACTION_ID>")
		if err != nil {
			t.Errorf("Method DeleteTransaction failed: %v", err)
		}
	})

	t.Run("Test CreateOperations", func(t *testing.T) {
		mockResponse := `
{
    "$id": "259125845563242502",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "status": "pending",
    "operations": 5,
    "expiresAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.CreateOperations("<TRANSACTION_ID>")
		if err != nil {
			t.Errorf("Method CreateOperations failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "name": "My Database",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "enabled": true,
    "type": "legacy",
    "policies": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "index1",
            "type": "primary",
            "status": "available",
            "error": "string",
            "attributes": [],
            "lengths": []
        }
    ],
    "archives": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": [],
            "databaseId": "5e5ea5c16897e",
            "name": "My Collection",
            "enabled": true,
            "documentSecurity": true,
            "attributes": [],
            "indexes": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "key": "index1",
                    "type": "primary",
                    "status": "available",
                    "error": "string",
                    "attributes": [],
                    "lengths": []
                }
            ],
            "bytesMax": 65535,
            "bytesUsed": 1500
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

		_, err := srv.Get("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
		}
	})

	t.Run("Test Update", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "name": "My Database",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "enabled": true,
    "type": "legacy",
    "policies": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "index1",
            "type": "primary",
            "status": "available",
            "error": "string",
            "attributes": [],
            "lengths": []
        }
    ],
    "archives": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": [],
            "databaseId": "5e5ea5c16897e",
            "name": "My Collection",
            "enabled": true,
            "documentSecurity": true,
            "attributes": [],
            "indexes": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "key": "index1",
                    "type": "primary",
                    "status": "available",
                    "error": "string",
                    "attributes": [],
                    "lengths": []
                }
            ],
            "bytesMax": 65535,
            "bytesUsed": 1500
        }
    ]
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

		_, err := srv.Update("<DATABASE_ID>")
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

		_, err := srv.Delete("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})

	t.Run("Test ListTables", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "tables": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": [],
            "databaseId": "5e5ea5c16897e",
            "name": "My Table",
            "enabled": true,
            "rowSecurity": true,
            "columns": [],
            "indexes": [
                {
                    "$id": "5e5ea5c16897e",
                    "$createdAt": "2020-10-15T06:38:00.000+00:00",
                    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
                    "key": "index1",
                    "type": "primary",
                    "status": "available",
                    "error": "string",
                    "columns": [],
                    "lengths": []
                }
            ],
            "bytesMax": 65535,
            "bytesUsed": 1500
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

		_, err := srv.ListTables("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method ListTables failed: %v", err)
		}
	})

	t.Run("Test CreateTable", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "databaseId": "5e5ea5c16897e",
    "name": "My Table",
    "enabled": true,
    "rowSecurity": true,
    "columns": [],
    "indexes": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "index1",
            "type": "primary",
            "status": "available",
            "error": "string",
            "columns": [],
            "lengths": []
        }
    ],
    "bytesMax": 65535,
    "bytesUsed": 1500
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

		_, err := srv.CreateTable("<DATABASE_ID>", "<TABLE_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateTable failed: %v", err)
		}
	})

	t.Run("Test GetTable", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "databaseId": "5e5ea5c16897e",
    "name": "My Table",
    "enabled": true,
    "rowSecurity": true,
    "columns": [],
    "indexes": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "index1",
            "type": "primary",
            "status": "available",
            "error": "string",
            "columns": [],
            "lengths": []
        }
    ],
    "bytesMax": 65535,
    "bytesUsed": 1500
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

		_, err := srv.GetTable("<DATABASE_ID>", "<TABLE_ID>")
		if err != nil {
			t.Errorf("Method GetTable failed: %v", err)
		}
	})

	t.Run("Test UpdateTable", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": [],
    "databaseId": "5e5ea5c16897e",
    "name": "My Table",
    "enabled": true,
    "rowSecurity": true,
    "columns": [],
    "indexes": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "index1",
            "type": "primary",
            "status": "available",
            "error": "string",
            "columns": [],
            "lengths": []
        }
    ],
    "bytesMax": 65535,
    "bytesUsed": 1500
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

		_, err := srv.UpdateTable("<DATABASE_ID>", "<TABLE_ID>")
		if err != nil {
			t.Errorf("Method UpdateTable failed: %v", err)
		}
	})

	t.Run("Test DeleteTable", func(t *testing.T) {
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

		_, err := srv.DeleteTable("<DATABASE_ID>", "<TABLE_ID>")
		if err != nil {
			t.Errorf("Method DeleteTable failed: %v", err)
		}
	})

	t.Run("Test ListColumns", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "columns": []
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

		_, err := srv.ListColumns("<DATABASE_ID>", "<TABLE_ID>")
		if err != nil {
			t.Errorf("Method ListColumns failed: %v", err)
		}
	})

	t.Run("Test CreateBooleanColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "isEnabled",
    "type": "boolean",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.CreateBooleanColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateBooleanColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateBooleanColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "isEnabled",
    "type": "boolean",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateBooleanColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, true)
		if err != nil {
			t.Errorf("Method UpdateBooleanColumn failed: %v", err)
		}
	})

	t.Run("Test CreateDatetimeColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "birthDay",
    "type": "datetime",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "format": "datetime"
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

		_, err := srv.CreateDatetimeColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateDatetimeColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateDatetimeColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "birthDay",
    "type": "datetime",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "format": "datetime"
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

		_, err := srv.UpdateDatetimeColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, "2020-10-15T06:38:00.000+00:00")
		if err != nil {
			t.Errorf("Method UpdateDatetimeColumn failed: %v", err)
		}
	})

	t.Run("Test CreateEmailColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "userEmail",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "format": "email"
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

		_, err := srv.CreateEmailColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateEmailColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateEmailColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "userEmail",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "format": "email"
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

		_, err := srv.UpdateEmailColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, "email@example.com")
		if err != nil {
			t.Errorf("Method UpdateEmailColumn failed: %v", err)
		}
	})

	t.Run("Test CreateEnumColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "status",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "elements": [],
    "format": "enum"
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

		_, err := srv.CreateEnumColumn("<DATABASE_ID>", "<TABLE_ID>", "", []string{}, true)
		if err != nil {
			t.Errorf("Method CreateEnumColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateEnumColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "status",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "elements": [],
    "format": "enum"
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

		_, err := srv.UpdateEnumColumn("<DATABASE_ID>", "<TABLE_ID>", "", []string{}, true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateEnumColumn failed: %v", err)
		}
	})

	t.Run("Test CreateFloatColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "percentageCompleted",
    "type": "double",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.CreateFloatColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateFloatColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateFloatColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "percentageCompleted",
    "type": "double",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateFloatColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, 1.0)
		if err != nil {
			t.Errorf("Method UpdateFloatColumn failed: %v", err)
		}
	})

	t.Run("Test CreateIntegerColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "count",
    "type": "integer",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.CreateIntegerColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateIntegerColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateIntegerColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "count",
    "type": "integer",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateIntegerColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, 1)
		if err != nil {
			t.Errorf("Method UpdateIntegerColumn failed: %v", err)
		}
	})

	t.Run("Test CreateIpColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "ipAddress",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "format": "ip"
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

		_, err := srv.CreateIpColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateIpColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateIpColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "ipAddress",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "format": "ip"
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

		_, err := srv.UpdateIpColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, "")
		if err != nil {
			t.Errorf("Method UpdateIpColumn failed: %v", err)
		}
	})

	t.Run("Test CreateLineColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.CreateLineColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateLineColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateLineColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateLineColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method UpdateLineColumn failed: %v", err)
		}
	})

	t.Run("Test CreateLongtextColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.CreateLongtextColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateLongtextColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateLongtextColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateLongtextColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateLongtextColumn failed: %v", err)
		}
	})

	t.Run("Test CreateMediumtextColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.CreateMediumtextColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateMediumtextColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateMediumtextColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateMediumtextColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateMediumtextColumn failed: %v", err)
		}
	})

	t.Run("Test CreatePointColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.CreatePointColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreatePointColumn failed: %v", err)
		}
	})

	t.Run("Test UpdatePointColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdatePointColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method UpdatePointColumn failed: %v", err)
		}
	})

	t.Run("Test CreatePolygonColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.CreatePolygonColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreatePolygonColumn failed: %v", err)
		}
	})

	t.Run("Test UpdatePolygonColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdatePolygonColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method UpdatePolygonColumn failed: %v", err)
		}
	})

	t.Run("Test CreateRelationshipColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "relatedTable": "table",
    "relationType": "oneToOne|oneToMany|manyToOne|manyToMany",
    "twoWay": true,
    "twoWayKey": "string",
    "onDelete": "restrict|cascade|setNull",
    "side": "parent|child"
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

		_, err := srv.CreateRelationshipColumn("<DATABASE_ID>", "<TABLE_ID>", "<RELATED_TABLE_ID>", "oneToOne")
		if err != nil {
			t.Errorf("Method CreateRelationshipColumn failed: %v", err)
		}
	})

	t.Run("Test CreateStringColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "size": 128
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

		_, err := srv.CreateStringColumn("<DATABASE_ID>", "<TABLE_ID>", "", 1, true)
		if err != nil {
			t.Errorf("Method CreateStringColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateStringColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "size": 128
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

		_, err := srv.UpdateStringColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateStringColumn failed: %v", err)
		}
	})

	t.Run("Test CreateTextColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.CreateTextColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateTextColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateTextColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00"
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

		_, err := srv.UpdateTextColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateTextColumn failed: %v", err)
		}
	})

	t.Run("Test CreateUrlColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "githubUrl",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "format": "url"
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

		_, err := srv.CreateUrlColumn("<DATABASE_ID>", "<TABLE_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateUrlColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateUrlColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "githubUrl",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "format": "url"
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

		_, err := srv.UpdateUrlColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, "https://example.com")
		if err != nil {
			t.Errorf("Method UpdateUrlColumn failed: %v", err)
		}
	})

	t.Run("Test CreateVarcharColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "size": 128
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

		_, err := srv.CreateVarcharColumn("<DATABASE_ID>", "<TABLE_ID>", "", 1, true)
		if err != nil {
			t.Errorf("Method CreateVarcharColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateVarcharColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "size": 128
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

		_, err := srv.UpdateVarcharColumn("<DATABASE_ID>", "<TABLE_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateVarcharColumn failed: %v", err)
		}
	})

	t.Run("Test GetColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "isEnabled",
    "type": "boolean",
    "status": "available",
    "error": "string",
    "required": true,
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

		_, err := srv.GetColumn("<DATABASE_ID>", "<TABLE_ID>", "")
		if err != nil {
			t.Errorf("Method GetColumn failed: %v", err)
		}
	})

	t.Run("Test DeleteColumn", func(t *testing.T) {
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

		_, err := srv.DeleteColumn("<DATABASE_ID>", "<TABLE_ID>", "")
		if err != nil {
			t.Errorf("Method DeleteColumn failed: %v", err)
		}
	})

	t.Run("Test UpdateRelationshipColumn", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "relatedTable": "table",
    "relationType": "oneToOne|oneToMany|manyToOne|manyToMany",
    "twoWay": true,
    "twoWayKey": "string",
    "onDelete": "restrict|cascade|setNull",
    "side": "parent|child"
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

		_, err := srv.UpdateRelationshipColumn("<DATABASE_ID>", "<TABLE_ID>", "")
		if err != nil {
			t.Errorf("Method UpdateRelationshipColumn failed: %v", err)
		}
	})

	t.Run("Test ListIndexes", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "indexes": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "key": "index1",
            "type": "primary",
            "status": "available",
            "error": "string",
            "columns": [],
            "lengths": []
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

		_, err := srv.ListIndexes("<DATABASE_ID>", "<TABLE_ID>")
		if err != nil {
			t.Errorf("Method ListIndexes failed: %v", err)
		}
	})

	t.Run("Test CreateIndex", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "key": "index1",
    "type": "primary",
    "status": "available",
    "error": "string",
    "columns": [],
    "lengths": []
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

		_, err := srv.CreateIndex("<DATABASE_ID>", "<TABLE_ID>", "", "key", []string{})
		if err != nil {
			t.Errorf("Method CreateIndex failed: %v", err)
		}
	})

	t.Run("Test GetIndex", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "key": "index1",
    "type": "primary",
    "status": "available",
    "error": "string",
    "columns": [],
    "lengths": []
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

		_, err := srv.GetIndex("<DATABASE_ID>", "<TABLE_ID>", "")
		if err != nil {
			t.Errorf("Method GetIndex failed: %v", err)
		}
	})

	t.Run("Test DeleteIndex", func(t *testing.T) {
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

		_, err := srv.DeleteIndex("<DATABASE_ID>", "<TABLE_ID>", "")
		if err != nil {
			t.Errorf("Method DeleteIndex failed: %v", err)
		}
	})

	t.Run("Test ListRows", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "rows": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$tableId": "5e5ea5c15117e",
            "$databaseId": "5e5ea5c15117e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": []
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

		_, err := srv.ListRows("<DATABASE_ID>", "<TABLE_ID>")
		if err != nil {
			t.Errorf("Method ListRows failed: %v", err)
		}
	})

	t.Run("Test CreateRow", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$tableId": "5e5ea5c15117e",
    "$databaseId": "5e5ea5c15117e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": []
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

		_, err := srv.CreateRow("<DATABASE_ID>", "<TABLE_ID>", "<ROW_ID>", map[string]interface{}{})
		if err != nil {
			t.Errorf("Method CreateRow failed: %v", err)
		}
	})

	t.Run("Test CreateRows", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "rows": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$tableId": "5e5ea5c15117e",
            "$databaseId": "5e5ea5c15117e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": []
        }
    ]
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

		_, err := srv.CreateRows("<DATABASE_ID>", "<TABLE_ID>", []interface{}{})
		if err != nil {
			t.Errorf("Method CreateRows failed: %v", err)
		}
	})

	t.Run("Test UpsertRows", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "rows": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$tableId": "5e5ea5c15117e",
            "$databaseId": "5e5ea5c15117e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": []
        }
    ]
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

		_, err := srv.UpsertRows("<DATABASE_ID>", "<TABLE_ID>", []interface{}{})
		if err != nil {
			t.Errorf("Method UpsertRows failed: %v", err)
		}
	})

	t.Run("Test UpdateRows", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "rows": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$tableId": "5e5ea5c15117e",
            "$databaseId": "5e5ea5c15117e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": []
        }
    ]
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

		_, err := srv.UpdateRows("<DATABASE_ID>", "<TABLE_ID>")
		if err != nil {
			t.Errorf("Method UpdateRows failed: %v", err)
		}
	})

	t.Run("Test DeleteRows", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "rows": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$tableId": "5e5ea5c15117e",
            "$databaseId": "5e5ea5c15117e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "$permissions": []
        }
    ]
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

		_, err := srv.DeleteRows("<DATABASE_ID>", "<TABLE_ID>")
		if err != nil {
			t.Errorf("Method DeleteRows failed: %v", err)
		}
	})

	t.Run("Test GetRow", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$tableId": "5e5ea5c15117e",
    "$databaseId": "5e5ea5c15117e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": []
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

		_, err := srv.GetRow("<DATABASE_ID>", "<TABLE_ID>", "<ROW_ID>")
		if err != nil {
			t.Errorf("Method GetRow failed: %v", err)
		}
	})

	t.Run("Test UpsertRow", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$tableId": "5e5ea5c15117e",
    "$databaseId": "5e5ea5c15117e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": []
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

		_, err := srv.UpsertRow("<DATABASE_ID>", "<TABLE_ID>", "<ROW_ID>")
		if err != nil {
			t.Errorf("Method UpsertRow failed: %v", err)
		}
	})

	t.Run("Test UpdateRow", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$tableId": "5e5ea5c15117e",
    "$databaseId": "5e5ea5c15117e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": []
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

		_, err := srv.UpdateRow("<DATABASE_ID>", "<TABLE_ID>", "<ROW_ID>")
		if err != nil {
			t.Errorf("Method UpdateRow failed: %v", err)
		}
	})

	t.Run("Test DeleteRow", func(t *testing.T) {
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

		_, err := srv.DeleteRow("<DATABASE_ID>", "<TABLE_ID>", "<ROW_ID>")
		if err != nil {
			t.Errorf("Method DeleteRow failed: %v", err)
		}
	})

	t.Run("Test DecrementRowColumn", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$tableId": "5e5ea5c15117e",
    "$databaseId": "5e5ea5c15117e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": []
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

		_, err := srv.DecrementRowColumn("<DATABASE_ID>", "<TABLE_ID>", "<ROW_ID>", "")
		if err != nil {
			t.Errorf("Method DecrementRowColumn failed: %v", err)
		}
	})

	t.Run("Test IncrementRowColumn", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$tableId": "5e5ea5c15117e",
    "$databaseId": "5e5ea5c15117e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "$permissions": []
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

		_, err := srv.IncrementRowColumn("<DATABASE_ID>", "<TABLE_ID>", "<ROW_ID>", "")
		if err != nil {
			t.Errorf("Method IncrementRowColumn failed: %v", err)
		}
	})
}
