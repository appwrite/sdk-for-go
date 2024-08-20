package models

import (
    "encoding/json"
    "errors"
)

// TemplateRuntime Model
type TemplateRuntime struct {
    // Runtime Name.
    Name string `json:"name"`
    // The build command used to build the deployment.
    Commands string `json:"commands"`
    // The entrypoint file used to execute the deployment.
    Entrypoint string `json:"entrypoint"`
    // Path to function in VCS (Version Control System) repository
    ProviderRootDirectory string `json:"providerRootDirectory"`

    // Used by Decode() method
    data []byte
}

func (model TemplateRuntime) New(data []byte) *TemplateRuntime {
    model.data = data
    return &model
}

func (model *TemplateRuntime) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}