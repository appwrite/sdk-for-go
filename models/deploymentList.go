package models

import (
	"encoding/json"
	"errors"
)

// DeploymentsList Model
type DeploymentList struct {
	// Total number of deployments that matched your query.
	Total int `json:"total"`
	// List of deployments.
	Deployments []Deployment `json:"deployments"`

	// Used by Decode() method
	data []byte
}

func (model DeploymentList) New(data []byte) *DeploymentList {
	model.data = data
	return &model
}

func (model *DeploymentList) UnmarshalJSON(data []byte) error {
	type alias DeploymentList
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"total\":0}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
}

func (model *DeploymentList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
