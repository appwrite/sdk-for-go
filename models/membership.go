package models

import (
	"encoding/json"
	"errors"
)

// Membership Model
type Membership struct {
	// Membership ID.
	Id string `json:"$id"`
	// Membership creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Membership update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// User ID.
	UserId string `json:"userId"`
	// User name. Hide this attribute by toggling membership privacy in the
	// Console.
	UserName string `json:"userName"`
	// User email address. Hide this attribute by toggling membership privacy in
	// the Console.
	UserEmail string `json:"userEmail"`
	// User phone number. Hide this attribute by toggling membership privacy in
	// the Console.
	UserPhone string `json:"userPhone"`
	// Team ID.
	TeamId string `json:"teamId"`
	// Team name.
	TeamName string `json:"teamName"`
	// Date, the user has been invited to join the team in ISO 8601 format.
	Invited string `json:"invited"`
	// Date, the user has accepted the invitation to join the team in ISO 8601
	// format.
	Joined string `json:"joined"`
	// User confirmation status, true if the user has joined the team or false
	// otherwise.
	Confirm bool `json:"confirm"`
	// Multi factor authentication status, true if the user has MFA enabled or
	// false otherwise. Hide this attribute by toggling membership privacy in the
	// Console.
	Mfa bool `json:"mfa"`
	// Most recent access date in ISO 8601 format. Show this attribute by toggling
	// membership privacy in the Console.
	UserAccessedAt string `json:"userAccessedAt"`
	// User list of roles
	Roles []string `json:"roles"`

	// Used by Decode() method
	data []byte
}

func (model Membership) New(data []byte) *Membership {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Membership) UnmarshalJSON(data []byte) error {
	type alias Membership
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id             json.RawMessage `json:"$id"`
		CreatedAt      json.RawMessage `json:"$createdAt"`
		UpdatedAt      json.RawMessage `json:"$updatedAt"`
		UserId         json.RawMessage `json:"userId"`
		UserName       json.RawMessage `json:"userName"`
		UserEmail      json.RawMessage `json:"userEmail"`
		UserPhone      json.RawMessage `json:"userPhone"`
		TeamId         json.RawMessage `json:"teamId"`
		TeamName       json.RawMessage `json:"teamName"`
		Invited        json.RawMessage `json:"invited"`
		Joined         json.RawMessage `json:"joined"`
		Confirm        json.RawMessage `json:"confirm"`
		Mfa            json.RawMessage `json:"mfa"`
		UserAccessedAt json.RawMessage `json:"userAccessedAt"`
		Roles          json.RawMessage `json:"roles"`
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
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.UserName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserName); err != nil {
			return err
		}
	}
	if present.UserEmail == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserEmail); err != nil {
			return err
		}
	}
	if present.UserPhone == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserPhone); err != nil {
			return err
		}
	}
	if present.TeamId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TeamId); err != nil {
			return err
		}
	}
	if present.TeamName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TeamName); err != nil {
			return err
		}
	}
	if present.Invited == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Invited); err != nil {
			return err
		}
	}
	if present.Joined == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Joined); err != nil {
			return err
		}
	}
	if present.Confirm == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Confirm); err != nil {
			return err
		}
	}
	if present.Mfa == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Mfa); err != nil {
			return err
		}
	}
	if present.UserAccessedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserAccessedAt); err != nil {
			return err
		}
	}
	if present.Roles == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Roles); err != nil {
			return err
		}
	}

	*model = Membership(decoded)
	return nil
}

func (model *Membership) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
