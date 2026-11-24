package models

import (
	"encoding/json"
	"errors"
)

// RuntimeLog Model
type RuntimeLog struct {
	// Log ID.
	Id string `json:"$id"`
	// Resource type: functions, sites or databases.
	ResourceType string `json:"resourceType"`
	// Resource ID.
	ResourceId string `json:"resourceId"`
	// Resource internal ID.
	ResourceInternalId string `json:"resourceInternalId"`
	// Deployment ID. Empty for databases, which have no deployments.
	DeploymentId string `json:"deploymentId"`
	// Region the runtime ran in when it wrote the log.
	Region string `json:"region"`
	// Log timestamp.
	Timestamp string `json:"timestamp"`
	// Log stream.
	Stream string `json:"stream"`
	// Log message.
	Message string `json:"message"`

	// Used by Decode() method
	data []byte
}

func (model RuntimeLog) New(data []byte) *RuntimeLog {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *RuntimeLog) UnmarshalJSON(data []byte) error {
	type alias RuntimeLog
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                 json.RawMessage `json:"$id"`
		ResourceType       json.RawMessage `json:"resourceType"`
		ResourceId         json.RawMessage `json:"resourceId"`
		ResourceInternalId json.RawMessage `json:"resourceInternalId"`
		DeploymentId       json.RawMessage `json:"deploymentId"`
		Region             json.RawMessage `json:"region"`
		Timestamp          json.RawMessage `json:"timestamp"`
		Stream             json.RawMessage `json:"stream"`
		Message            json.RawMessage `json:"message"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.ResourceType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceType); err != nil {
			return err
		}
	}
	if present.ResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceId); err != nil {
			return err
		}
	}
	if present.ResourceInternalId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceInternalId); err != nil {
			return err
		}
	}
	if present.DeploymentId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeploymentId); err != nil {
			return err
		}
	}
	if present.Region == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Region); err != nil {
			return err
		}
	}
	if present.Timestamp == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Timestamp); err != nil {
			return err
		}
	}
	if present.Stream == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Stream); err != nil {
			return err
		}
	}
	if present.Message == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Message); err != nil {
			return err
		}
	}

	*model = RuntimeLog(decoded)
	return nil
}

func (model *RuntimeLog) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
