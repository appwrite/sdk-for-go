package models

import (
	"encoding/json"
	"errors"
)

// PolicyMembershipPrivacy Model
type PolicyMembershipPrivacy struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether user ID is visible in memberships.
	UserId bool `json:"userId"`
	// Whether user email is visible in memberships.
	UserEmail bool `json:"userEmail"`
	// Whether user phone is visible in memberships.
	UserPhone bool `json:"userPhone"`
	// Whether user name is visible in memberships.
	UserName bool `json:"userName"`
	// Whether user MFA status is visible in memberships.
	UserMFA bool `json:"userMFA"`
	// Whether user last access time is visible in memberships.
	UserAccessedAt bool `json:"userAccessedAt"`

	// Used by Decode() method
	data []byte
}

func (model PolicyMembershipPrivacy) New(data []byte) *PolicyMembershipPrivacy {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyMembershipPrivacy) UnmarshalJSON(data []byte) error {
	type alias PolicyMembershipPrivacy
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id             json.RawMessage `json:"$id"`
		UserId         json.RawMessage `json:"userId"`
		UserEmail      json.RawMessage `json:"userEmail"`
		UserPhone      json.RawMessage `json:"userPhone"`
		UserName       json.RawMessage `json:"userName"`
		UserMFA        json.RawMessage `json:"userMFA"`
		UserAccessedAt json.RawMessage `json:"userAccessedAt"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.UserEmail == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.UserEmail); err != nil {
			return err
		}
	}
	if present.UserPhone == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.UserPhone); err != nil {
			return err
		}
	}
	if present.UserName == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.UserName); err != nil {
			return err
		}
	}
	if present.UserMFA == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.UserMFA); err != nil {
			return err
		}
	}
	if present.UserAccessedAt == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.UserAccessedAt); err != nil {
			return err
		}
	}

	*model = PolicyMembershipPrivacy(decoded)
	return nil
}

func (model *PolicyMembershipPrivacy) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
