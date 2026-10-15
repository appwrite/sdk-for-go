package models

import (
	"encoding/json"
	"errors"
)

// WafRuleRedirect Model
type WafRuleRedirect struct {
	// Rule ID.
	Id string `json:"$id"`
	// WAF rule creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// WAF rule last update time in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Human friendly rule name.
	Name string `json:"name"`
	// Optional description for the rule.
	Description string `json:"description"`
	// Team ID.
	TeamId string `json:"teamId"`
	// Project ID.
	ProjectId string `json:"projectId"`
	// Resource type the rule is scoped to.
	ResourceType string `json:"resourceType"`
	// Resource identifier. Empty for API-wide rules.
	ResourceId string `json:"resourceId"`
	// Action performed when the rule matches.
	Action string `json:"action"`
	// Evaluation priority. Lower values execute earlier.
	Priority int `json:"priority"`
	// Whether the rule is active.
	Enabled bool `json:"enabled"`
	// List of conditions evaluated for this rule.
	Conditions []interface{} `json:"conditions"`
	// Action specific configuration.
	Config interface{} `json:"config"`
	// Target location for the redirect.
	Location string `json:"location"`
	// HTTP status code used for the redirect.
	StatusCode int `json:"statusCode"`

	// Used by Decode() method
	data []byte
}

func (model WafRuleRedirect) New(data []byte) *WafRuleRedirect {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *WafRuleRedirect) UnmarshalJSON(data []byte) error {
	type alias WafRuleRedirect
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
		Name         json.RawMessage `json:"name"`
		Description  json.RawMessage `json:"description"`
		TeamId       json.RawMessage `json:"teamId"`
		ProjectId    json.RawMessage `json:"projectId"`
		ResourceType json.RawMessage `json:"resourceType"`
		ResourceId   json.RawMessage `json:"resourceId"`
		Priority     json.RawMessage `json:"priority"`
		Enabled      json.RawMessage `json:"enabled"`
		Conditions   json.RawMessage `json:"conditions"`
		Config       json.RawMessage `json:"config"`
		Location     json.RawMessage `json:"location"`
		StatusCode   json.RawMessage `json:"statusCode"`
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
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.TeamId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TeamId); err != nil {
			return err
		}
	}
	if present.ProjectId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProjectId); err != nil {
			return err
		}
	}
	if present.ResourceType == nil {
		if err := json.Unmarshal([]byte("\"api\""), &decoded.ResourceType); err != nil {
			return err
		}
	}
	if present.ResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceId); err != nil {
			return err
		}
	}
	if present.Priority == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Priority); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Conditions == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Conditions); err != nil {
			return err
		}
	}
	if present.Config == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Config); err != nil {
			return err
		}
	}
	if present.Location == nil {
		if err := json.Unmarshal([]byte("\"/\""), &decoded.Location); err != nil {
			return err
		}
	}
	if present.StatusCode == nil {
		if err := json.Unmarshal([]byte("302"), &decoded.StatusCode); err != nil {
			return err
		}
	}

	*model = WafRuleRedirect(decoded)
	return nil
}

func (model *WafRuleRedirect) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
