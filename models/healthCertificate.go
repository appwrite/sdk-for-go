package models

import (
    "encoding/json"
    "errors"
)

// HealthCertificate Model
type HealthCertificate struct {
    // Certificate name
    Name string `json:"name"`
    // Subject SN
    SubjectSN string `json:"subjectSN"`
    // Issuer organisation
    IssuerOrganisation string `json:"issuerOrganisation"`
    // Valid from
    ValidFrom string `json:"validFrom"`
    // Valid to
    ValidTo string `json:"validTo"`
    // Signature type SN
    SignatureTypeSN string `json:"signatureTypeSN"`

    // Used by Decode() method
    data []byte
}

func (model HealthCertificate) New(data []byte) *HealthCertificate {
    model.data = data
    return &model
}

func (model *HealthCertificate) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}