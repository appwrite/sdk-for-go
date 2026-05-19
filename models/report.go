package models

import (
    "encoding/json"
    "errors"
)

// Report Model
type Report struct {
    // Report ID.
    Id string `json:"$id"`
    // Report creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Report update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // ID of the third-party app that submitted the report.
    AppId string `json:"appId"`
    // Analyzer that produced this report. e.g. lighthouse, audit,
    // databaseAnalyzer.
    Type string `json:"type"`
    // Short, human-readable title for the report.
    Title string `json:"title"`
    // Markdown summary describing the report.
    Summary string `json:"summary"`
    // Plural noun describing what the report analyzes, e.g. databases, sites,
    // urls.
    TargetType string `json:"targetType"`
    // Free-form target identifier (URL for lighthouse, resource ID for db).
    Target string `json:"target"`
    // Categories covered by the report, e.g. performance, accessibility.
    Categories []string `json:"categories"`
    // Insights nested under this report.
    Insights []Insight `json:"insights"`
    // Time the report was analyzed in ISO 8601 format.
    AnalyzedAt string `json:"analyzedAt"`

    // Used by Decode() method
    data []byte
}

func (model Report) New(data []byte) *Report {
    model.data = data
    return &model
}

func (model *Report) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}