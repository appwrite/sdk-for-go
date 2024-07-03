package models


// User Model
type User struct {
    // User ID.
    Id string `json:"$id"`
    // User creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // User update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // User name.
    Name string `json:"name"`
    // Hashed user password.
    Password string `json:"password"`
    // Password hashing algorithm.
    Hash string `json:"hash"`
    // Password hashing algorithm configuration.
    HashOptions interface{} `json:"hashOptions"`
    // User registration date in ISO 8601 format.
    Registration string `json:"registration"`
    // User status. Pass `true` for enabled and `false` for disabled.
    Status bool `json:"status"`
    // Labels for the user.
    Labels []interface{} `json:"labels"`
    // Password update time in ISO 8601 format.
    PasswordUpdate string `json:"passwordUpdate"`
    // User email address.
    Email string `json:"email"`
    // User phone number in E.164 format.
    Phone string `json:"phone"`
    // Email verification status.
    EmailVerification bool `json:"emailVerification"`
    // Phone verification status.
    PhoneVerification bool `json:"phoneVerification"`
    // Multi factor authentication status.
    Mfa bool `json:"mfa"`
    // User preferences as a key-value object
    Prefs interface{} `json:"prefs"`
    // A user-owned message receiver. A single user may have multiple e.g. emails,
    // phones, and a browser. Each target is registered with a single provider.
    Targets []interface{} `json:"targets"`
    // Most recent access date in ISO 8601 format. This attribute is only updated
    // again after 24 hours.
    AccessedAt string `json:"accessedAt"`

}
