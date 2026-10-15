package models

import (
	"encoding/json"
	"errors"
)

// Execution Model
type Execution struct {
	// Execution ID.
	Id string `json:"$id"`
	// Execution creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Execution update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Execution roles.
	Permissions []string `json:"$permissions"`
	// Function or site ID.
	ResourceId string `json:"resourceId"`
	// Execution resource type.
	ResourceType string `json:"resourceType"`
	// Deployment ID used to create the execution.
	DeploymentId string `json:"deploymentId"`
	// The trigger that caused the resource to execute. Possible values can be:
	// `http`, `schedule`, `event`, or `email`.
	Trigger string `json:"trigger"`
	// The status of the resource execution. Possible values can be: `waiting`,
	// `processing`, `completed`, `failed`, or `scheduled`.
	Status string `json:"status"`
	// HTTP request method type.
	RequestMethod string `json:"requestMethod"`
	// HTTP request path and query.
	RequestPath string `json:"requestPath"`
	// HTTP request headers as a key-value object. This will return only
	// whitelisted headers. All headers are returned if execution is created as
	// synchronous.
	RequestHeaders []Headers `json:"requestHeaders"`
	// HTTP response status code.
	ResponseStatusCode int `json:"responseStatusCode"`
	// HTTP response body. This will return empty unless execution is created as
	// synchronous.
	ResponseBody string `json:"responseBody"`
	// HTTP response headers as a key-value object. This will return only
	// whitelisted headers. All headers are returned if execution is created as
	// synchronous.
	ResponseHeaders []Headers `json:"responseHeaders"`
	// Resource logs. Includes the last 4,000 characters. This will return an
	// empty string unless the response is returned using an API key or as part of
	// a webhook payload.
	Logs string `json:"logs"`
	// Resource errors. Includes the last 4,000 characters. This will return an
	// empty string unless the response is returned using an API key or as part of
	// a webhook payload.
	Errors string `json:"errors"`
	// Total time the resource(function/site) took to respond, in seconds.
	Duration float64 `json:"duration"`
	// The scheduled time for execution. If left empty, execution will be queued
	// immediately.
	ScheduledAt *string `json:"scheduledAt"`

	// Used by Decode() method
	data []byte
}

func (model Execution) New(data []byte) *Execution {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Execution) UnmarshalJSON(data []byte) error {
	type alias Execution
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                 json.RawMessage `json:"$id"`
		CreatedAt          json.RawMessage `json:"$createdAt"`
		UpdatedAt          json.RawMessage `json:"$updatedAt"`
		ResourceId         json.RawMessage `json:"resourceId"`
		ResourceType       json.RawMessage `json:"resourceType"`
		DeploymentId       json.RawMessage `json:"deploymentId"`
		RequestMethod      json.RawMessage `json:"requestMethod"`
		RequestPath        json.RawMessage `json:"requestPath"`
		ResponseStatusCode json.RawMessage `json:"responseStatusCode"`
		ResponseBody       json.RawMessage `json:"responseBody"`
		Logs               json.RawMessage `json:"logs"`
		Errors             json.RawMessage `json:"errors"`
		Duration           json.RawMessage `json:"duration"`
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
	if present.ResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceId); err != nil {
			return err
		}
	}
	if present.ResourceType == nil {
		if err := json.Unmarshal([]byte("\"functions\""), &decoded.ResourceType); err != nil {
			return err
		}
	}
	if present.DeploymentId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeploymentId); err != nil {
			return err
		}
	}
	if present.RequestMethod == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RequestMethod); err != nil {
			return err
		}
	}
	if present.RequestPath == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RequestPath); err != nil {
			return err
		}
	}
	if present.ResponseStatusCode == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ResponseStatusCode); err != nil {
			return err
		}
	}
	if present.ResponseBody == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResponseBody); err != nil {
			return err
		}
	}
	if present.Logs == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Logs); err != nil {
			return err
		}
	}
	if present.Errors == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Errors); err != nil {
			return err
		}
	}
	if present.Duration == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Duration); err != nil {
			return err
		}
	}

	*model = Execution(decoded)
	return nil
}

func (model *Execution) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
