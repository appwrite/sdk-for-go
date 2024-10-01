package models

import (
    "encoding/json"
    "errors"
)

// Runtime Model
type Runtime struct {
    // Runtime ID.
    Id string `json:"$id"`
    // Parent runtime key.
    Key string `json:"key"`
    // Runtime Name.
    Name string `json:"name"`
    // Runtime version.
    Version string `json:"version"`
    // Base Docker image used to build the runtime.
    Base string `json:"base"`
    // Image name of Docker Hub.
    Image string `json:"image"`
    // Name of the logo image.
    Logo string `json:"logo"`
    // List of supported architectures.
    Supports []string `json:"supports"`

    // Used by Decode() method
    data []byte
}

func (model Runtime) New(data []byte) *Runtime {
    model.data = data
    return &model
}

func (model *Runtime) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
