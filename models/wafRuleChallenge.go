package models

import (
	"encoding/json"
	"errors"
)

// WafRuleChallenge Model
type WafRuleChallenge struct {
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
	// Challenge type enforced when the rule matches.
	ChallengeType string `json:"challengeType"`
	// Challenge difficulty from 1 (easiest) to 5 (hardest) enforced when the rule
	// matches.
	Difficulty int `json:"difficulty"`
	// Seconds a visitor stays cleared after passing the challenge before being
	// challenged again.
	Ttl int `json:"ttl"`

	// Used by Decode() method
	data []byte
}

func (model WafRuleChallenge) New(data []byte) *WafRuleChallenge {
	model.data = data
	return &model
}

func (model *WafRuleChallenge) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
