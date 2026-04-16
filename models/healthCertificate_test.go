package models

import (
    "encoding/json"
    "testing"
)

func TestHealthCertificateModel(t *testing.T) {
    model := HealthCertificate{        Name: "/CN=www.google.com",        SubjectSN: "string",        IssuerOrganisation: "string",        ValidFrom: "1704200998",        ValidTo: "1711458597",        SignatureTypeSN: "RSA-SHA256",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result HealthCertificate
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.SubjectSN != model.SubjectSN {
        t.Errorf("Expected SubjectSN %v, got %v", model.SubjectSN, result.SubjectSN)
    }
    if result.IssuerOrganisation != model.IssuerOrganisation {
        t.Errorf("Expected IssuerOrganisation %v, got %v", model.IssuerOrganisation, result.IssuerOrganisation)
    }
    if result.ValidFrom != model.ValidFrom {
        t.Errorf("Expected ValidFrom %v, got %v", model.ValidFrom, result.ValidFrom)
    }
    if result.ValidTo != model.ValidTo {
        t.Errorf("Expected ValidTo %v, got %v", model.ValidTo, result.ValidTo)
    }
    if result.SignatureTypeSN != model.SignatureTypeSN {
        t.Errorf("Expected SignatureTypeSN %v, got %v", model.SignatureTypeSN, result.SignatureTypeSN)
    }}
