package models

import (
    "encoding/json"
    "testing"
)

func TestMfaRecoveryCodesModel(t *testing.T) {
    model := MfaRecoveryCodes{        RecoveryCodes: []string{"test"},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result MfaRecoveryCodes
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }}
