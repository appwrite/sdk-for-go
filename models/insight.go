package models

import (
    "encoding/json"
    "errors"
)

// Insight Model
type Insight struct {
    // Insight ID.
    Id string `json:"$id"`
    // Insight creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Insight update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Parent report ID. Insights always belong to a report.
    ReportId string `json:"reportId"`
    // Insight type. One of databaseIndex (legacy), tablesDBIndex,
    // documentsDBIndex, vectorsDBIndex, databasePerformance, sitePerformance,
    // siteAccessibility, siteSeo, functionPerformance. The index types are
    // engine-specific so each CTA can pair the right service+method
    // (databases.createIndex, tablesDB.createIndex, documentsDB.createIndex, or
    // vectorsDB.createIndex).
    Type string `json:"type"`
    // Insight severity. One of info, warning, critical.
    Severity string `json:"severity"`
    // Insight status. One of active, dismissed.
    Status string `json:"status"`
    // Type of the resource the insight is about. Plural noun, e.g. databases,
    // sites, functions.
    ResourceType string `json:"resourceType"`
    // ID of the resource the insight is about.
    ResourceId string `json:"resourceId"`
    // Plural noun for the parent resource that contains the insight's resource,
    // e.g. an insight about a column index on a table → resourceType=indexes,
    // parentResourceType=tables. Empty when the resource has no parent.
    ParentResourceType string `json:"parentResourceType"`
    // ID of the parent resource. Empty when the resource has no parent.
    ParentResourceId string `json:"parentResourceId"`
    // Insight title.
    Title string `json:"title"`
    // Short markdown summary describing the insight.
    Summary string `json:"summary"`
    // List of call-to-action buttons attached to this insight.
    Ctas []InsightCTA `json:"ctas"`
    // Time the insight was analyzed in ISO 8601 format.
    AnalyzedAt string `json:"analyzedAt"`
    // Time the insight was dismissed in ISO 8601 format. Empty when not
    // dismissed.
    DismissedAt string `json:"dismissedAt"`
    // User ID that dismissed the insight. Empty when not dismissed.
    DismissedBy string `json:"dismissedBy"`

    // Used by Decode() method
    data []byte
}

func (model Insight) New(data []byte) *Insight {
    model.data = data
    return &model
}

func (model *Insight) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}