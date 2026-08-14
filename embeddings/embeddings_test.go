package embeddings

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestEmbeddings(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test CreateTextEmbeddings", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "embeddings": [
        {
            "model": "nomic-embed-text",
            "dimension": 768,
            "embedding": [],
            "error": "Error message"
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

		_, err := srv.CreateTextEmbeddings([]string{})
		if err != nil {
			t.Errorf("Method CreateTextEmbeddings failed: %v", err)
		}
	})
}
