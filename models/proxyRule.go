package models

import (
    "encoding/json"
    "errors"
)

// Rule Model
type ProxyRule struct {
    // Rule ID.
    Id string `json:"$id"`
    // Rule creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Rule update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Domain name.
    Domain string `json:"domain"`
    // Action definition for the rule. Possible values are "api", "deployment", or
    // "redirect"
    Type string `json:"type"`
    // Defines how the rule was created. Possible values are "manual" or
    // "deployment"
    Trigger string `json:"trigger"`
    // URL to redirect to. Used if type is "redirect"
    RedirectUrl string `json:"redirectUrl"`
    // Status code to apply during redirect. Used if type is "redirect"
    RedirectStatusCode int `json:"redirectStatusCode"`
    // ID of deployment. Used if type is "deployment"
    DeploymentId string `json:"deploymentId"`
    // Type of deployment. Possible values are "function", "site". Used if rule's
    // type is "deployment".
    DeploymentResourceType string `json:"deploymentResourceType"`
    // ID of deployment's resource (site or function ID). Used if type is
    // "deployment"
    DeploymentResourceId string `json:"deploymentResourceId"`
    // Name of Git branch that updates rule. Used if type is "deployment"
    DeploymentVcsProviderBranch string `json:"deploymentVcsProviderBranch"`
    // Domain verification status. Possible values are "unverified", "verifying",
    // "verified"
    Status string `json:"status"`
    // Logs from rule verification or certificate generation. Certificate
    // generation logs are prioritized if both are available.
    Logs string `json:"logs"`
    // Certificate auto-renewal date in ISO 8601 format.
    RenewAt string `json:"renewAt"`

    // Used by Decode() method
    data []byte
}

func (model ProxyRule) New(data []byte) *ProxyRule {
    model.data = data
    return &model
}

func (model *ProxyRule) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}