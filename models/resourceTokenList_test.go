package models

import (
    "encoding/json"
    "testing"
)

func TestResourceTokenListModel(t *testing.T) {
    model := ResourceTokenList{        Total: 5,        Tokens: []ResourceToken{ResourceToken{        Id: "bb8ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        ResourceId: "5e5ea5c168bb8:5e5ea5c168bb8",        ResourceType: "files",        Expire: "2020-10-15T06:38:00.000+00:00",        Secret: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",        AccessedAt: "2020-10-15T06:38:00.000+00:00",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ResourceTokenList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
