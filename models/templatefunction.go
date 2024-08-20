package models

import (
    "encoding/json"
    "errors"
)

// TemplateFunction Model
type TemplateFunction struct {
    // Function Template Icon.
    Icon string `json:"icon"`
    // Function Template ID.
    Id string `json:"id"`
    // Function Template Name.
    Name string `json:"name"`
    // Function Template Tagline.
    Tagline string `json:"tagline"`
    // Execution permissions.
    Permissions []string `json:"permissions"`
    // Function trigger events.
    Events []string `json:"events"`
    // Function execution schedult in CRON format.
    Cron string `json:"cron"`
    // Function execution timeout in seconds.
    Timeout int `json:"timeout"`
    // Function use cases.
    UseCases []string `json:"useCases"`
    // List of runtimes that can be used with this template.
    Runtimes []TemplateRuntime `json:"runtimes"`
    // Function Template Instructions.
    Instructions string `json:"instructions"`
    // VCS (Version Control System) Provider.
    VcsProvider string `json:"vcsProvider"`
    // VCS (Version Control System) Repository ID
    ProviderRepositoryId string `json:"providerRepositoryId"`
    // VCS (Version Control System) Owner.
    ProviderOwner string `json:"providerOwner"`
    // VCS (Version Control System) branch version (tag).
    ProviderVersion string `json:"providerVersion"`
    // Function variables.
    Variables []TemplateVariable `json:"variables"`
    // Function scopes.
    Scopes []string `json:"scopes"`

    // Used by Decode() method
    data []byte
}

func (model TemplateFunction) New(data []byte) *TemplateFunction {
    model.data = data
    return &model
}

func (model *TemplateFunction) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}