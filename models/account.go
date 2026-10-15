package models

import (
	"encoding/json"
	"errors"
)

// Account Model
type Account struct {
	// User ID.
	Id string `json:"$id"`
	// User creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// User update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// User name.
	Name string `json:"name"`
	// User registration date in ISO 8601 format.
	Registration string `json:"registration"`
	// User status. Pass `true` for enabled and `false` for disabled.
	Status bool `json:"status"`
	// Labels for the user.
	Labels []string `json:"labels"`
	// Password update time in ISO 8601 format.
	PasswordUpdate string `json:"passwordUpdate"`
	// User email address.
	Email string `json:"email"`
	// User phone number in E.164 format.
	Phone string `json:"phone"`
	// Email verification status.
	EmailVerification bool `json:"emailVerification"`
	// Canonical form of the user email address.
	EmailCanonical *string `json:"emailCanonical"`
	// Whether the user email is from a free email provider.
	EmailIsFree *bool `json:"emailIsFree"`
	// Whether the user email is from a disposable email provider.
	EmailIsDisposable *bool `json:"emailIsDisposable"`
	// Whether the user email is from a corporate domain.
	EmailIsCorporate *bool `json:"emailIsCorporate"`
	// Whether the user email is in its canonical form.
	EmailIsCanonical *bool `json:"emailIsCanonical"`
	// Whether the password was found in a known data breach the last time it was
	// checked. Null when the password has never been checked.
	PasswordPwned *bool `json:"passwordPwned"`
	// Phone verification status.
	PhoneVerification bool `json:"phoneVerification"`
	// Multi factor authentication status.
	Mfa bool `json:"mfa"`
	// User preferences as a key-value object
	Prefs Preferences `json:"prefs"`
	// A user-owned message receiver. A single user may have multiple e.g. emails,
	// phones, and a browser. Each target is registered with a single provider.
	Targets []Target `json:"targets"`
	// Most recent access date in ISO 8601 format. This attribute is only updated
	// again after 24 hours.
	AccessedAt string `json:"accessedAt"`
	// Whether the user can impersonate other users.
	Impersonator *bool `json:"impersonator"`
	// ID of the original actor performing the impersonation. Present only when
	// the current request is impersonating another user. Internal audit logs
	// attribute the action to this user, while the impersonated target is
	// recorded only in internal audit payload data.
	ImpersonatorUserId *string `json:"impersonatorUserId"`

	// Used by Decode() method
	data []byte
}

func (model Account) New(data []byte) *Account {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Account) UnmarshalJSON(data []byte) error {
	type alias Account
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                json.RawMessage `json:"$id"`
		CreatedAt         json.RawMessage `json:"$createdAt"`
		UpdatedAt         json.RawMessage `json:"$updatedAt"`
		Name              json.RawMessage `json:"name"`
		Registration      json.RawMessage `json:"registration"`
		Status            json.RawMessage `json:"status"`
		Labels            json.RawMessage `json:"labels"`
		PasswordUpdate    json.RawMessage `json:"passwordUpdate"`
		Email             json.RawMessage `json:"email"`
		Phone             json.RawMessage `json:"phone"`
		EmailVerification json.RawMessage `json:"emailVerification"`
		PhoneVerification json.RawMessage `json:"phoneVerification"`
		Mfa               json.RawMessage `json:"mfa"`
		AccessedAt        json.RawMessage `json:"accessedAt"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Registration == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Registration); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Labels == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Labels); err != nil {
			return err
		}
	}
	if present.PasswordUpdate == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PasswordUpdate); err != nil {
			return err
		}
	}
	if present.Email == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Email); err != nil {
			return err
		}
	}
	if present.Phone == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Phone); err != nil {
			return err
		}
	}
	if present.EmailVerification == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.EmailVerification); err != nil {
			return err
		}
	}
	if present.PhoneVerification == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.PhoneVerification); err != nil {
			return err
		}
	}
	if present.Mfa == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Mfa); err != nil {
			return err
		}
	}
	if present.AccessedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AccessedAt); err != nil {
			return err
		}
	}

	*model = Account(decoded)
	return nil
}

func (model *Account) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
