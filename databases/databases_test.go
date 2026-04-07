package databases

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v2/client"
)

func TestDatabases(t *testing.T) {
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

	t.Run("Test ListCollections", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "collections": [
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

		_, err := srv.ListCollections("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method ListCollections failed: %v", err)
		}
	})

	t.Run("Test CreateCollection", func(t *testing.T) {
		mockResponse := `
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

		_, err := srv.CreateCollection("<DATABASE_ID>", "<COLLECTION_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method CreateCollection failed: %v", err)
		}
	})

	t.Run("Test GetCollection", func(t *testing.T) {
		mockResponse := `
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

		_, err := srv.GetCollection("<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method GetCollection failed: %v", err)
		}
	})

	t.Run("Test UpdateCollection", func(t *testing.T) {
		mockResponse := `
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

		_, err := srv.UpdateCollection("<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method UpdateCollection failed: %v", err)
		}
	})

	t.Run("Test DeleteCollection", func(t *testing.T) {
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

		_, err := srv.DeleteCollection("<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method DeleteCollection failed: %v", err)
		}
	})

	t.Run("Test ListAttributes", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "attributes": []
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

		_, err := srv.ListAttributes("<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method ListAttributes failed: %v", err)
		}
	})

	t.Run("Test CreateBooleanAttribute", func(t *testing.T) {
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

		_, err := srv.CreateBooleanAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateBooleanAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateBooleanAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateBooleanAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, true)
		if err != nil {
			t.Errorf("Method UpdateBooleanAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateDatetimeAttribute", func(t *testing.T) {
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

		_, err := srv.CreateDatetimeAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateDatetimeAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateDatetimeAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateDatetimeAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, "2020-10-15T06:38:00.000+00:00")
		if err != nil {
			t.Errorf("Method UpdateDatetimeAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateEmailAttribute", func(t *testing.T) {
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

		_, err := srv.CreateEmailAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateEmailAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateEmailAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateEmailAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, "email@example.com")
		if err != nil {
			t.Errorf("Method UpdateEmailAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateEnumAttribute", func(t *testing.T) {
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

		_, err := srv.CreateEnumAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", []string{}, true)
		if err != nil {
			t.Errorf("Method CreateEnumAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateEnumAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateEnumAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", []string{}, true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateEnumAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateFloatAttribute", func(t *testing.T) {
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

		_, err := srv.CreateFloatAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateFloatAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateFloatAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateFloatAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, 1.0)
		if err != nil {
			t.Errorf("Method UpdateFloatAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateIntegerAttribute", func(t *testing.T) {
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

		_, err := srv.CreateIntegerAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateIntegerAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateIntegerAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateIntegerAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, 1)
		if err != nil {
			t.Errorf("Method UpdateIntegerAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateIpAttribute", func(t *testing.T) {
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

		_, err := srv.CreateIpAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateIpAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateIpAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateIpAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, "")
		if err != nil {
			t.Errorf("Method UpdateIpAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateLineAttribute", func(t *testing.T) {
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

		_, err := srv.CreateLineAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateLineAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateLineAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateLineAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method UpdateLineAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateLongtextAttribute", func(t *testing.T) {
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

		_, err := srv.CreateLongtextAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateLongtextAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateLongtextAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateLongtextAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateLongtextAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateMediumtextAttribute", func(t *testing.T) {
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

		_, err := srv.CreateMediumtextAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateMediumtextAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateMediumtextAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateMediumtextAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateMediumtextAttribute failed: %v", err)
		}
	})

	t.Run("Test CreatePointAttribute", func(t *testing.T) {
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

		_, err := srv.CreatePointAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreatePointAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdatePointAttribute", func(t *testing.T) {
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

		_, err := srv.UpdatePointAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method UpdatePointAttribute failed: %v", err)
		}
	})

	t.Run("Test CreatePolygonAttribute", func(t *testing.T) {
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

		_, err := srv.CreatePolygonAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreatePolygonAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdatePolygonAttribute", func(t *testing.T) {
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

		_, err := srv.UpdatePolygonAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method UpdatePolygonAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateRelationshipAttribute", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "relatedCollection": "collection",
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

		_, err := srv.CreateRelationshipAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "<RELATED_COLLECTION_ID>", "oneToOne")
		if err != nil {
			t.Errorf("Method CreateRelationshipAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateRelationshipAttribute", func(t *testing.T) {
		mockResponse := `
{
    "key": "fullName",
    "type": "string",
    "status": "available",
    "error": "string",
    "required": true,
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "relatedCollection": "collection",
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

		_, err := srv.UpdateRelationshipAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "")
		if err != nil {
			t.Errorf("Method UpdateRelationshipAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateStringAttribute", func(t *testing.T) {
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

		_, err := srv.CreateStringAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", 1, true)
		if err != nil {
			t.Errorf("Method CreateStringAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateStringAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateStringAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateStringAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateTextAttribute", func(t *testing.T) {
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

		_, err := srv.CreateTextAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateTextAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateTextAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateTextAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateTextAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateUrlAttribute", func(t *testing.T) {
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

		_, err := srv.CreateUrlAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true)
		if err != nil {
			t.Errorf("Method CreateUrlAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateUrlAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateUrlAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, "https://example.com")
		if err != nil {
			t.Errorf("Method UpdateUrlAttribute failed: %v", err)
		}
	})

	t.Run("Test CreateVarcharAttribute", func(t *testing.T) {
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

		_, err := srv.CreateVarcharAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", 1, true)
		if err != nil {
			t.Errorf("Method CreateVarcharAttribute failed: %v", err)
		}
	})

	t.Run("Test UpdateVarcharAttribute", func(t *testing.T) {
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

		_, err := srv.UpdateVarcharAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "", true, "<DEFAULT>")
		if err != nil {
			t.Errorf("Method UpdateVarcharAttribute failed: %v", err)
		}
	})

	t.Run("Test GetAttribute", func(t *testing.T) {
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

		_, err := srv.GetAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "")
		if err != nil {
			t.Errorf("Method GetAttribute failed: %v", err)
		}
	})

	t.Run("Test DeleteAttribute", func(t *testing.T) {
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

		_, err := srv.DeleteAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "")
		if err != nil {
			t.Errorf("Method DeleteAttribute failed: %v", err)
		}
	})

	t.Run("Test ListDocuments", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "documents": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.ListDocuments("<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method ListDocuments failed: %v", err)
		}
	})

	t.Run("Test CreateDocument", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.CreateDocument("<DATABASE_ID>", "<COLLECTION_ID>", "<DOCUMENT_ID>", map[string]interface{}{})
		if err != nil {
			t.Errorf("Method CreateDocument failed: %v", err)
		}
	})

	t.Run("Test CreateDocuments", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "documents": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.CreateDocuments("<DATABASE_ID>", "<COLLECTION_ID>", []interface{}{})
		if err != nil {
			t.Errorf("Method CreateDocuments failed: %v", err)
		}
	})

	t.Run("Test UpsertDocuments", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "documents": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.UpsertDocuments("<DATABASE_ID>", "<COLLECTION_ID>", []interface{}{})
		if err != nil {
			t.Errorf("Method UpsertDocuments failed: %v", err)
		}
	})

	t.Run("Test UpdateDocuments", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "documents": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.UpdateDocuments("<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method UpdateDocuments failed: %v", err)
		}
	})

	t.Run("Test DeleteDocuments", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "documents": [
        {
            "$id": "5e5ea5c16897e",
            "$sequence": "1",
            "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.DeleteDocuments("<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method DeleteDocuments failed: %v", err)
		}
	})

	t.Run("Test GetDocument", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.GetDocument("<DATABASE_ID>", "<COLLECTION_ID>", "<DOCUMENT_ID>")
		if err != nil {
			t.Errorf("Method GetDocument failed: %v", err)
		}
	})

	t.Run("Test UpsertDocument", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.UpsertDocument("<DATABASE_ID>", "<COLLECTION_ID>", "<DOCUMENT_ID>")
		if err != nil {
			t.Errorf("Method UpsertDocument failed: %v", err)
		}
	})

	t.Run("Test UpdateDocument", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.UpdateDocument("<DATABASE_ID>", "<COLLECTION_ID>", "<DOCUMENT_ID>")
		if err != nil {
			t.Errorf("Method UpdateDocument failed: %v", err)
		}
	})

	t.Run("Test DeleteDocument", func(t *testing.T) {
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

		_, err := srv.DeleteDocument("<DATABASE_ID>", "<COLLECTION_ID>", "<DOCUMENT_ID>")
		if err != nil {
			t.Errorf("Method DeleteDocument failed: %v", err)
		}
	})

	t.Run("Test DecrementDocumentAttribute", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.DecrementDocumentAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "<DOCUMENT_ID>", "")
		if err != nil {
			t.Errorf("Method DecrementDocumentAttribute failed: %v", err)
		}
	})

	t.Run("Test IncrementDocumentAttribute", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$sequence": "1",
    "$collectionId": "5e5ea5c15117e",
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

		_, err := srv.IncrementDocumentAttribute("<DATABASE_ID>", "<COLLECTION_ID>", "<DOCUMENT_ID>", "")
		if err != nil {
			t.Errorf("Method IncrementDocumentAttribute failed: %v", err)
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
            "attributes": [],
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

		_, err := srv.ListIndexes("<DATABASE_ID>", "<COLLECTION_ID>")
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
    "attributes": [],
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

		_, err := srv.CreateIndex("<DATABASE_ID>", "<COLLECTION_ID>", "", "key", []string{})
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
    "attributes": [],
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

		_, err := srv.GetIndex("<DATABASE_ID>", "<COLLECTION_ID>", "")
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

		_, err := srv.DeleteIndex("<DATABASE_ID>", "<COLLECTION_ID>", "")
		if err != nil {
			t.Errorf("Method DeleteIndex failed: %v", err)
		}
	})
}
