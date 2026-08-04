package models

import (
    "encoding/json"
    "errors"
)

// Connections Model
type DatabaseStatusConnections struct {
    // Current number of active connections.
    Current int `json:"current"`
    // The engine's own max_connections. On a pooled database this is the backend
    // limit the pooler multiplexes onto, not the ceiling a client pool may reach
    // — that is networkMaxConnections on the database resource.
    Max int `json:"max"`

    // Used by Decode() method
    data []byte
}

func (model DatabaseStatusConnections) New(data []byte) *DatabaseStatusConnections {
    model.data = data
    return &model
}

func (model *DatabaseStatusConnections) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}