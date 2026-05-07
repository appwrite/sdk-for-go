package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2ProviderListModel(t *testing.T) {
    model := OAuth2ProviderList{        Total: 5,        Providers: []interface{}{},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2ProviderList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
