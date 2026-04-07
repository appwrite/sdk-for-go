package models

import (
    "encoding/json"
    "testing"
)

func TestTokenModel(t *testing.T) {
    model := Token{        Id: "bb8ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UserId: "5e5ea5c168bb8",        Secret: "string",        Expire: "2020-10-15T06:38:00.000+00:00",        Phrase: "Golden Fox",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Token
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
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
    }
    if result.Expire != model.Expire {
        t.Errorf("Expected Expire %v, got %v", model.Expire, result.Expire)
    }
    if result.Phrase != model.Phrase {
        t.Errorf("Expected Phrase %v, got %v", model.Phrase, result.Phrase)
    }}
