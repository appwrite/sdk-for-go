package models


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

}
