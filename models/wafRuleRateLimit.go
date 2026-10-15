package models

import (
	"encoding/json"
	"errors"
)

// WafRuleRateLimit Model
type WafRuleRateLimit struct {
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
	// Maximum number of matching requests allowed for the given interval.
	Limit int `json:"limit"`
	// Interval in seconds for the rate limit window.
	Interval int `json:"interval"`
	// Rate limit key: `ip` limits per client IP, `userId` limits per
	// authenticated user.
	Key string `json:"key"`
	// Rate limit strategy: `fixedWindow`, `slidingWindow`, or `tokenBucket`.
	Strategy string `json:"strategy"`
	// Maximum burst capacity in tokens, applied by the `tokenBucket` strategy.
	// The sustained refill rate is `limit / interval`. Defaults to `limit`.
	MaxBucketSize int `json:"maxBucketSize"`

	// Used by Decode() method
	data []byte
}

func (model WafRuleRateLimit) New(data []byte) *WafRuleRateLimit {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *WafRuleRateLimit) UnmarshalJSON(data []byte) error {
	type alias WafRuleRateLimit
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id            json.RawMessage `json:"$id"`
		CreatedAt     json.RawMessage `json:"$createdAt"`
		UpdatedAt     json.RawMessage `json:"$updatedAt"`
		Name          json.RawMessage `json:"name"`
		Description   json.RawMessage `json:"description"`
		TeamId        json.RawMessage `json:"teamId"`
		ProjectId     json.RawMessage `json:"projectId"`
		ResourceType  json.RawMessage `json:"resourceType"`
		ResourceId    json.RawMessage `json:"resourceId"`
		Priority      json.RawMessage `json:"priority"`
		Enabled       json.RawMessage `json:"enabled"`
		Conditions    json.RawMessage `json:"conditions"`
		Config        json.RawMessage `json:"config"`
		Limit         json.RawMessage `json:"limit"`
		Interval      json.RawMessage `json:"interval"`
		Key           json.RawMessage `json:"key"`
		Strategy      json.RawMessage `json:"strategy"`
		MaxBucketSize json.RawMessage `json:"maxBucketSize"`
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
	if present.Limit == nil {
		if err := json.Unmarshal([]byte("100"), &decoded.Limit); err != nil {
			return err
		}
	}
	if present.Interval == nil {
		if err := json.Unmarshal([]byte("3600"), &decoded.Interval); err != nil {
			return err
		}
	}
	if present.Key == nil {
		if err := json.Unmarshal([]byte("\"ip\""), &decoded.Key); err != nil {
			return err
		}
	}
	if present.Strategy == nil {
		if err := json.Unmarshal([]byte("\"fixedWindow\""), &decoded.Strategy); err != nil {
			return err
		}
	}
	if present.MaxBucketSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.MaxBucketSize); err != nil {
			return err
		}
	}

	*model = WafRuleRateLimit(decoded)
	return nil
}

func (model *WafRuleRateLimit) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
