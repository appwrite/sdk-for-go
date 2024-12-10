package models

import (
    "encoding/json"
    "errors"
)

// UsersList Model
type UserList struct {
    // Total number of users documents that matched your query.
    Total int `json:"total"`
    // List of users.
    Users []User `json:"users"`

    // Used by Decode() method
    data []byte
}

func (model UserList) New(data []byte) *UserList {
    model.data = data
    return &model
}

func (model *UserList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}