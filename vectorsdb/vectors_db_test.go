package vectorsdb

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestVectorsDB(t *testing.T) {
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
            "type": "legacy"
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
    "type": "legacy"
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

	t.Run("Test ListSpecifications", func(t *testing.T) {
		mockResponse := `
{
    "specifications": [
        {
            "slug": "s-2vcpu-2gb",
            "name": "Standard",
            "price": 20,
            "cpu": 2000,
            "memory": 2048,
            "maxConnections": 200,
            "includedStorage": 25,
            "includedBandwidth": 200,
            "enabled": true
        }
    ],
    "total": 9,
    "pricing": {
        "storageOverageRate": 0.125,
        "bandwidthOverageRate": 0.08,
        "replicaRate": 1,
        "pitrRate": 0.2
    }
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

		_, err := srv.ListSpecifications()
		if err != nil {
			t.Errorf("Method ListSpecifications failed: %v", err)
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
    "type": "legacy"
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
    "type": "legacy"
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

		_, err := srv.Update("<DATABASE_ID>", "<NAME>")
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
            "bytesUsed": 1500,
            "dimension": 1536
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
    "bytesUsed": 1500,
    "dimension": 1536
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

		_, err := srv.CreateCollection("<DATABASE_ID>", "<COLLECTION_ID>", "<NAME>", 1)
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
    "bytesUsed": 1500,
    "dimension": 1536
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
    "bytesUsed": 1500,
    "dimension": 1536
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

		_, err := srv.UpdateCollection("<DATABASE_ID>", "<COLLECTION_ID>", "<NAME>")
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

	t.Run("Test CreateQuery", func(t *testing.T) {
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

		_, err := srv.CreateQuery("<DATABASE_ID>", "<COLLECTION_ID>")
		if err != nil {
			t.Errorf("Method CreateQuery failed: %v", err)
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

		_, err := srv.CreateIndex("<DATABASE_ID>", "<COLLECTION_ID>", "", "hnsw_euclidean", []string{})
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

	t.Run("Test CreateFailover", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
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

		_, err := srv.CreateFailover("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method CreateFailover failed: %v", err)
		}
	})

	t.Run("Test ListOperations", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "operations": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "databaseId": "5e5ea5c16897e",
            "type": "update",
            "status": "completed",
            "attempts": 1,
            "errorCode": "Interrupted",
            "errorMessage": "string"
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

		_, err := srv.ListOperations("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method ListOperations failed: %v", err)
		}
	})

	t.Run("Test GetReplicas", func(t *testing.T) {
		mockResponse := `
{
    "replicas": 2,
    "syncMode": "async",
    "syncDegraded": true,
    "syncAcknowledgements": 1,
    "syncStandbyCount": 2,
    "members": [
        {
            "$id": "1",
            "role": "replica",
            "status": "active"
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

		_, err := srv.GetReplicas("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method GetReplicas failed: %v", err)
		}
	})

	t.Run("Test GetStatus", func(t *testing.T) {
		mockResponse := `
{
    "health": "healthy",
    "ready": true,
    "engine": "postgresql",
    "version": "17",
    "uptime": 86400,
    "connections": {
        "current": 12,
        "max": 100
    },
    "syncMode": "async",
    "syncDegraded": true,
    "syncAcknowledgements": 1,
    "syncStandbyCount": 2,
    "replicas": [
        {
            "index": 0,
            "role": "primary",
            "healthy": true
        }
    ],
    "volumes": [
        {
            "path": "/var/lib/postgresql/data",
            "usedPercent": "45%",
            "available": "55GB",
            "mounted": true
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

		_, err := srv.GetStatus("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method GetStatus failed: %v", err)
		}
	})
}
