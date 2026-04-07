package models

import (
    "encoding/json"
    "testing"
)

func TestMfaTypeModel(t *testing.T) {
    model := MfaType{        Secret: "[SHARED_SECRET]",        Uri: "otpauth://totp/appwrite:user@example.com?secret=[SHARED_SECRET]&amp;issuer=appwrite",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result MfaType
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
    }
    if result.Uri != model.Uri {
        t.Errorf("Expected Uri %v, got %v", model.Uri, result.Uri)
    }}
