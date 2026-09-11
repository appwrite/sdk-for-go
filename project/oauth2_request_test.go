package project

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestOAuth2ProviderRequests(t *testing.T) {
	for _, provider := range []string{"kakao", "tiktok"} {
		for _, enabled := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/enabled=%t", provider, enabled), func(t *testing.T) {
				requests := 0
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					requests++
					if r.Method != http.MethodPatch || r.URL.Path != "/project/oauth2/"+provider {
						t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
					}
					if r.Header.Get("Accept") != "application/json" {
						t.Errorf("unexpected Accept header: %q", r.Header.Get("Accept"))
					}
					var body map[string]any
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Errorf("decode request: %v", err)
						w.WriteHeader(http.StatusBadRequest)
						return
					}
					if body["clientId"] != "test-client" || body["clientSecret"] != "test-secret" || body["enabled"] != enabled {
						t.Errorf("unexpected OAuth configuration: %#v", body)
					}
					w.Header().Set("Content-Type", "application/json")
					// The server's answer intentionally differs from the request.
					_, _ = w.Write([]byte(`{"$id":"` + provider + `","enabled":false,"clientId":"saved-client"}`))
				}))
				defer ts.Close()

				c := client.New()
				c.Endpoint = ts.URL
				c.Client = ts.Client()
				srv := New(c)
				switch provider {
				case "kakao":
					result, err := srv.UpdateOAuth2Kakao(srv.WithUpdateOAuth2KakaoClientId("test-client"), srv.WithUpdateOAuth2KakaoClientSecret("test-secret"), srv.WithUpdateOAuth2KakaoEnabled(enabled))
					if err != nil {
						t.Fatal(err)
					}
					if result.Id != "kakao" || result.Enabled || result.ClientId != "saved-client" {
						t.Errorf("unexpected saved configuration: %+v", result)
					}
				case "tiktok":
					result, err := srv.UpdateOAuth2TikTok(srv.WithUpdateOAuth2TikTokClientId("test-client"), srv.WithUpdateOAuth2TikTokClientSecret("test-secret"), srv.WithUpdateOAuth2TikTokEnabled(enabled))
					if err != nil {
						t.Fatal(err)
					}
					if result.Id != "tiktok" || result.Enabled || result.ClientId != "saved-client" {
						t.Errorf("unexpected saved configuration: %+v", result)
					}
				}
				if requests != 1 {
					t.Errorf("expected one request, got %d", requests)
				}
			})
		}
	}
}
