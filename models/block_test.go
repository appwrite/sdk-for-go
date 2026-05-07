package models

import (
    "encoding/json"
    "testing"
)

func TestBlockModel(t *testing.T) {
    model := Block{        CreatedAt: "2020-10-15T06:38:00.000+00:00",        ResourceType: "project",        ResourceId: "5e5ea5c16897e",        ProjectName: "My Project",        Region: "fra",        OrganizationName: "Acme Inc.",        OrganizationId: "5e5ea5c16897e",        BillingPlan: "pro",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Block
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.ResourceType != model.ResourceType {
        t.Errorf("Expected ResourceType %v, got %v", model.ResourceType, result.ResourceType)
    }
    if result.ResourceId != model.ResourceId {
        t.Errorf("Expected ResourceId %v, got %v", model.ResourceId, result.ResourceId)
    }
    if result.ProjectName != model.ProjectName {
        t.Errorf("Expected ProjectName %v, got %v", model.ProjectName, result.ProjectName)
    }
    if result.Region != model.Region {
        t.Errorf("Expected Region %v, got %v", model.Region, result.Region)
    }
    if result.OrganizationName != model.OrganizationName {
        t.Errorf("Expected OrganizationName %v, got %v", model.OrganizationName, result.OrganizationName)
    }
    if result.OrganizationId != model.OrganizationId {
        t.Errorf("Expected OrganizationId %v, got %v", model.OrganizationId, result.OrganizationId)
    }
    if result.BillingPlan != model.BillingPlan {
        t.Errorf("Expected BillingPlan %v, got %v", model.BillingPlan, result.BillingPlan)
    }}
