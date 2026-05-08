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

    // Used by Decode() method
    data []byte
}

func (model PolicyMembershipPrivacy) New(data []byte) *PolicyMembershipPrivacy {
    model.data = data
    return &model
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