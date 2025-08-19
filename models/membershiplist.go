package models

import (
    "encoding/json"
    "errors"
)

// MembershipsList Model
type MembershipList struct {
    // Total number of memberships rows that matched your query.
    Total int `json:"total"`
    // List of memberships.
    Memberships []Membership `json:"memberships"`

    // Used by Decode() method
    data []byte
}

func (model MembershipList) New(data []byte) *MembershipList {
    model.data = data
    return &model
}

func (model *MembershipList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}