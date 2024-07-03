package models


// MFAType Model
type MfaType struct {
    // Secret token used for TOTP factor.
    Secret string `json:"secret"`
    // URI for authenticator apps.
    Uri string `json:"uri"`

}
