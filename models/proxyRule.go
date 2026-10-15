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
	// Type of deployment resource: function, site, or dedicatedDatabase.
	DeploymentResourceType *string `json:"deploymentResourceType"`
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
	// Ingress protocol for the rule.
	Protocol string `json:"protocol"`

	// Used by Decode() method
	data []byte
}

func (model ProxyRule) New(data []byte) *ProxyRule {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *ProxyRule) UnmarshalJSON(data []byte) error {
	type alias ProxyRule
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                          json.RawMessage `json:"$id"`
		CreatedAt                   json.RawMessage `json:"$createdAt"`
		UpdatedAt                   json.RawMessage `json:"$updatedAt"`
		Domain                      json.RawMessage `json:"domain"`
		Type                        json.RawMessage `json:"type"`
		Trigger                     json.RawMessage `json:"trigger"`
		RedirectUrl                 json.RawMessage `json:"redirectUrl"`
		RedirectStatusCode          json.RawMessage `json:"redirectStatusCode"`
		DeploymentId                json.RawMessage `json:"deploymentId"`
		DeploymentResourceId        json.RawMessage `json:"deploymentResourceId"`
		DeploymentVcsProviderBranch json.RawMessage `json:"deploymentVcsProviderBranch"`
		Status                      json.RawMessage `json:"status"`
		Logs                        json.RawMessage `json:"logs"`
		Protocol                    json.RawMessage `json:"protocol"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.Domain == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Domain); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Trigger == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Trigger); err != nil {
			return err
		}
	}
	if present.RedirectUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RedirectUrl); err != nil {
			return err
		}
	}
	if present.RedirectStatusCode == nil {
		if err := json.Unmarshal([]byte("301"), &decoded.RedirectStatusCode); err != nil {
			return err
		}
	}
	if present.DeploymentId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeploymentId); err != nil {
			return err
		}
	}
	if present.DeploymentResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeploymentResourceId); err != nil {
			return err
		}
	}
	if present.DeploymentVcsProviderBranch == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeploymentVcsProviderBranch); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"unverified\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Logs == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Logs); err != nil {
			return err
		}
	}
	if present.Protocol == nil {
		if err := json.Unmarshal([]byte("\"http\""), &decoded.Protocol); err != nil {
			return err
		}
	}

	*model = ProxyRule(decoded)
	return nil
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
