package models

import (
    "encoding/json"
    "testing"
)

func TestProviderListModel(t *testing.T) {
    model := ProviderList{        Total: 5,        Providers: []Provider{Provider{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "Mailgun",        Provider: "mailgun",        Enabled: true,        Type: "sms",        Credentials: map[string]interface{}{},    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ProviderList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
