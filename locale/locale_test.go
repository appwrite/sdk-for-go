package locale

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v2/client"
)

func TestLocale(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "ip": "127.0.0.1",
    "countryCode": "US",
    "country": "United States",
    "continentCode": "NA",
    "continent": "North America",
    "eu": true,
    "currency": "USD"
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

	t.Run("Test ListCodes", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "localeCodes": [
        {
            "code": "en-us",
            "name": "US"
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

		_, err := srv.ListCodes()
		if err != nil {
			t.Errorf("Method ListCodes failed: %v", err)
		}
	})

	t.Run("Test ListContinents", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "continents": [
        {
            "name": "Europe",
            "code": "EU"
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

		_, err := srv.ListContinents()
		if err != nil {
			t.Errorf("Method ListContinents failed: %v", err)
		}
	})

	t.Run("Test ListCountries", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "countries": [
        {
            "name": "United States",
            "code": "US"
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

		_, err := srv.ListCountries()
		if err != nil {
			t.Errorf("Method ListCountries failed: %v", err)
		}
	})

	t.Run("Test ListCountriesEU", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "countries": [
        {
            "name": "United States",
            "code": "US"
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

		_, err := srv.ListCountriesEU()
		if err != nil {
			t.Errorf("Method ListCountriesEU failed: %v", err)
		}
	})

	t.Run("Test ListCountriesPhones", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "phones": [
        {
            "code": "+1",
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

		_, err := srv.ListCountriesPhones()
		if err != nil {
			t.Errorf("Method ListCountriesPhones failed: %v", err)
		}
	})

	t.Run("Test ListCurrencies", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "currencies": [
        {
            "symbol": "$",
            "name": "US dollar",
            "symbolNative": "$",
            "decimalDigits": 2,
            "rounding": 0,
            "code": "USD",
            "namePlural": "US dollars"
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

		_, err := srv.ListCurrencies()
		if err != nil {
			t.Errorf("Method ListCurrencies failed: %v", err)
		}
	})

	t.Run("Test ListLanguages", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "languages": [
        {
            "name": "Italian",
            "code": "it",
            "nativeName": "Italiano"
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

		_, err := srv.ListLanguages()
		if err != nil {
			t.Errorf("Method ListLanguages failed: %v", err)
		}
	})
}
