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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DeploymentList) UnmarshalJSON(data []byte) error {
	type alias DeploymentList
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Total json.RawMessage `json:"total"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Total == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Total); err != nil {
			return err
		}
	}

	*model = DeploymentList(decoded)
	return nil
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
