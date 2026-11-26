package models

import (
	"encoding/json"
	"errors"
)

// Alert Model
type Alert struct {
	// Alert ID.
	Id string `json:"$id"`
	// Date the alert started firing, in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Alert update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// ID of the monitor that raised the alert.
	MonitorId string `json:"monitorId"`
	// ID of the project the alert is about. Empty for an alert of the whole
	// organization.
	ProjectId string `json:"projectId"`
	// Name of the monitor when the alert fired.
	Name string `json:"name"`
	// Whether the alert is still firing or has resolved.
	Status string `json:"status"`
	// Date the alert resolved, in ISO 8601 format. Empty while it fires.
	ResolvedAt *string `json:"resolvedAt"`
	// Date a member acknowledged the alert, in ISO 8601 format. Empty while
	// nobody has.
	AcknowledgedAt *string `json:"acknowledgedAt"`
	// ID of the user who acknowledged the alert. Empty while nobody has, or when
	// it was acknowledged with an API key.
	AcknowledgedBy string `json:"acknowledgedBy"`
	// Name of the user who acknowledged the alert, when they did.
	AcknowledgedByName string `json:"acknowledgedByName"`
	// ID of the escalation notifying the team of this alert. Empty when the
	// monitor notified its `notify` roles instead.
	EscalationId string `json:"escalationId"`
	// Index of the escalation's next step to notify, counting from 0. Equal to
	// the number of steps once the last round has gone out.
	EscalationStep int `json:"escalationStep"`
	// How many times the escalation's steps have been repeated for this alert.
	EscalationRound int `json:"escalationRound"`
	// When the next step is notified, in ISO 8601 format. Null when nothing is
	// due: the alert was acknowledged or resolved, or every step has gone out.
	EscalatesAt *string `json:"escalatesAt"`
	// The monitor's rules when the alert fired, as query strings.
	Rules []string `json:"rules"`
	// The monitor's severity when the alert fired: `info`, `warning`, `error` or
	// `critical`.
	Severity string `json:"severity"`
	// Metric the alert fired on. Null for a probe monitor's alert.
	Metric *string `json:"metric"`
	// How the metric was compared with the threshold. Null for a probe monitor's
	// alert.
	Comparator *string `json:"comparator"`
	// Threshold the metric crossed, in the metric's unit.
	Threshold float64 `json:"threshold"`
	// Latest reading while the alert fired: the metric's value for a usage
	// monitor, and for a probe monitor its `latency`.
	Value float64 `json:"value"`
	// Reading furthest past the threshold while the alert fired.
	Peak float64 `json:"peak"`
	// Type of the monitor that raised the alert.
	Type string `json:"type"`
	// Type of the resource the alert fired on: the project, or one database,
	// bucket, function or site.
	ResourceType string `json:"resourceType"`
	// ID of the resource the alert fired on. The project ID for a project-wide
	// monitor.
	ResourceId string `json:"resourceId"`
	// How many times the alert was announced as firing: once when it fired, and
	// again after each cooldown while it kept firing.
	Notifications int `json:"notifications"`
	// Probe monitors: response time in milliseconds of the latest check recorded
	// on the alert.
	Latency *int `json:"latency"`
	// What the failing check of a probe monitor found when the alert opened, such
	// as `connect_refused`. Null for a usage monitor's alert.
	Result *string `json:"result"`

	// Used by Decode() method
	data []byte
}

func (model Alert) New(data []byte) *Alert {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Alert) UnmarshalJSON(data []byte) error {
	type alias Alert
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                 json.RawMessage `json:"$id"`
		CreatedAt          json.RawMessage `json:"$createdAt"`
		UpdatedAt          json.RawMessage `json:"$updatedAt"`
		MonitorId          json.RawMessage `json:"monitorId"`
		ProjectId          json.RawMessage `json:"projectId"`
		Name               json.RawMessage `json:"name"`
		Status             json.RawMessage `json:"status"`
		AcknowledgedBy     json.RawMessage `json:"acknowledgedBy"`
		AcknowledgedByName json.RawMessage `json:"acknowledgedByName"`
		EscalationId       json.RawMessage `json:"escalationId"`
		EscalationStep     json.RawMessage `json:"escalationStep"`
		EscalationRound    json.RawMessage `json:"escalationRound"`
		Rules              json.RawMessage `json:"rules"`
		Severity           json.RawMessage `json:"severity"`
		Threshold          json.RawMessage `json:"threshold"`
		Value              json.RawMessage `json:"value"`
		Peak               json.RawMessage `json:"peak"`
		Type               json.RawMessage `json:"type"`
		ResourceType       json.RawMessage `json:"resourceType"`
		ResourceId         json.RawMessage `json:"resourceId"`
		Notifications      json.RawMessage `json:"notifications"`
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
	if present.MonitorId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.MonitorId); err != nil {
			return err
		}
	}
	if present.ProjectId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProjectId); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"firing\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.AcknowledgedBy == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AcknowledgedBy); err != nil {
			return err
		}
	}
	if present.AcknowledgedByName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AcknowledgedByName); err != nil {
			return err
		}
	}
	if present.EscalationId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.EscalationId); err != nil {
			return err
		}
	}
	if present.EscalationStep == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.EscalationStep); err != nil {
			return err
		}
	}
	if present.EscalationRound == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.EscalationRound); err != nil {
			return err
		}
	}
	if present.Rules == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Rules); err != nil {
			return err
		}
	}
	if present.Severity == nil {
		if err := json.Unmarshal([]byte("\"warning\""), &decoded.Severity); err != nil {
			return err
		}
	}
	if present.Threshold == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Threshold); err != nil {
			return err
		}
	}
	if present.Value == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Value); err != nil {
			return err
		}
	}
	if present.Peak == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Peak); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"usage\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.ResourceType == nil {
		if err := json.Unmarshal([]byte("\"project\""), &decoded.ResourceType); err != nil {
			return err
		}
	}
	if present.ResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceId); err != nil {
			return err
		}
	}
	if present.Notifications == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Notifications); err != nil {
			return err
		}
	}

	*model = Alert(decoded)
	return nil
}

func (model *Alert) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
