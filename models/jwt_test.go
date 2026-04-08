package models

import (
    "encoding/json"
    "testing"
)

func TestJwtModel(t *testing.T) {
    model := Jwt{        Jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Jwt
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Jwt != model.Jwt {
        t.Errorf("Expected Jwt %v, got %v", model.Jwt, result.Jwt)
    }}
