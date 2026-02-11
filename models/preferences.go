package models

import (
    "encoding/json"
    "errors"
)

// Preferences Model
type Preferences struct {

    // Used by Decode() method
    data []byte
}

func (model Preferences) New(data []byte) *Preferences {
    model.data = data
    return &model
}

func (p *Preferences) UnmarshalJSON(b []byte) error {
    p.data = make([]byte, len(b))
    copy(p.data, b)
    return nil
}

// Use this method to get response in desired type
func (model *Preferences) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}