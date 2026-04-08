package models

import (
    "encoding/json"
    "testing"
)

func TestIdentityModel(t *testing.T) {
    model := Identity{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        UserId: "5e5bb8c16897e",        Provider: "email",        ProviderUid: "5e5bb8c16897e",        ProviderEmail: "user@example.com",        ProviderAccessToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",        ProviderAccessTokenExpiry: "2020-10-15T06:38:00.000+00:00",        ProviderRefreshToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Identity
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.Provider != model.Provider {
        t.Errorf("Expected Provider %v, got %v", model.Provider, result.Provider)
    }
    if result.ProviderUid != model.ProviderUid {
        t.Errorf("Expected ProviderUid %v, got %v", model.ProviderUid, result.ProviderUid)
    }
    if result.ProviderEmail != model.ProviderEmail {
        t.Errorf("Expected ProviderEmail %v, got %v", model.ProviderEmail, result.ProviderEmail)
    }
    if result.ProviderAccessToken != model.ProviderAccessToken {
        t.Errorf("Expected ProviderAccessToken %v, got %v", model.ProviderAccessToken, result.ProviderAccessToken)
    }
    if result.ProviderAccessTokenExpiry != model.ProviderAccessTokenExpiry {
        t.Errorf("Expected ProviderAccessTokenExpiry %v, got %v", model.ProviderAccessTokenExpiry, result.ProviderAccessTokenExpiry)
    }
    if result.ProviderRefreshToken != model.ProviderRefreshToken {
        t.Errorf("Expected ProviderRefreshToken %v, got %v", model.ProviderRefreshToken, result.ProviderRefreshToken)
    }}
