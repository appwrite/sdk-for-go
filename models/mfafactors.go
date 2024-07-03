package models


// MFAFactors Model
type MfaFactors struct {
    // Can TOTP be used for MFA challenge for this account.
    Totp bool `json:"totp"`
    // Can phone (SMS) be used for MFA challenge for this account.
    Phone bool `json:"phone"`
    // Can email be used for MFA challenge for this account.
    Email bool `json:"email"`
    // Can recovery code be used for MFA challenge for this account.
    RecoveryCode bool `json:"recoveryCode"`

}
