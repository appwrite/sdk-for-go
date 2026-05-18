package models

import (
    "encoding/json"
    "errors"
)

// InsightCTA Model
type InsightCTA struct {
    // Human-readable label for the CTA, used in UI.
    Label string `json:"label"`
    // Public API service (SDK namespace) the client should invoke. Must match the
    // engine that owns the resource — for index suggestions: databases
    // (legacy), tablesDB, documentsDB, or vectorsDB.
    Service string `json:"service"`
    // Public API method on the chosen service the client should invoke when this
    // CTA is triggered.
    Method string `json:"method"`
    // Parameter map the client should pass to the service method when this CTA is
    // triggered. Keys match the target API's parameter names (e.g.
    // databaseId/tableId/columns for tablesDB, databaseId/collectionId/attributes
    // for the legacy Databases API).
    Params interface{} `json:"params"`

    // Used by Decode() method
    data []byte
}

func (model InsightCTA) New(data []byte) *InsightCTA {
    model.data = data
    return &model
}

func (model *InsightCTA) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}