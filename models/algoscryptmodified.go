package models


// AlgoScryptModified Model
type AlgoScryptModified struct {
    // Algo type.
    Type string `json:"xtype"`
    // Salt used to compute hash.
    Salt string `json:"salt"`
    // Separator used to compute hash.
    SaltSeparator string `json:"saltSeparator"`
    // Key used to compute hash.
    SignerKey string `json:"signerKey"`

}
