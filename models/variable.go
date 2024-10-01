package models

import (
    "encoding/json"
    "errors"
)

// Variable Model
type Variable struct {
    // Variable ID.
    Id string `json:"$id"`
    // Variable creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Variable creation date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Variable key.
    Key string `json:"key"`
    // Variable value.
    Value string `json:"value"`
    // Service to which the variable belongs. Possible values are "project",
    // "function"
    ResourceType string `json:"resourceType"`
    // ID of resource to which the variable belongs. If resourceType is "project",
    // it is empty. If resourceType is "function", it is ID of the function.
    ResourceId string `json:"resourceId"`

    // Used by Decode() method
    data []byte
}

func (model Variable) New(data []byte) *Variable {
    model.data = data
    return &model
}

func (model *Variable) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
