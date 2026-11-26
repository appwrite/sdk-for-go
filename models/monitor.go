package models

import (
	"encoding/json"
	"errors"
)

// Monitor Model
type Monitor struct {
	// Monitor ID.
	Id string `json:"$id"`
	// Monitor creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Monitor update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Monitor name.
	Name string `json:"name"`
	// Project ID.
	ProjectId string `json:"projectId"`
	// Rules that must all hold for the monitor to fire, as query strings. A usage
	// monitor has one comparing a metric with a threshold, optionally scoped with
	// `equal('resourceType')` and `equal('resourceId')`. A probe monitor fires
	// when its rules match the check, over these variables: `http` `statusCode`,
	// `latency`, `certificateDays` and `certificateLifetime`; `dns` `answer`,
	// `answers` and `latency`; `tcp` `latency`; `smtp` `code`, `latency`,
	// `certificateDays` and `certificateLifetime`. A failed connection fires
	// whatever the rules.
	Rules []string `json:"rules"`
	// What the monitor checks: `usage` compares a usage metric with a threshold;
	// `http`, `dns`, `tcp` and `smtp` probe a target on their own interval.
	Type string `json:"type"`
	// Type of the resource the monitor watches, from its rules: the project, or
	// one database, bucket, function or site.
	ResourceType string `json:"resourceType"`
	// ID of the resource the monitor watches, from its rules. The project ID for
	// a project-wide monitor.
	ResourceId string `json:"resourceId"`
	// Seconds a counter metric is summed over, and a growth metric is measured
	// across. Gauges compare their latest value. Usage is read a short lag behind
	// the evaluation (`_APP_ALERTS_USAGE_LAG`, 120 seconds by default) so events
	// still being written are counted by the next evaluation.
	Window int `json:"window"`
	// Minimum seconds between two team notifications of this monitor firing.
	// Events fire on every change regardless.
	Cooldown int `json:"cooldown"`
	// Whether the monitor is evaluated.
	Enabled bool `json:"enabled"`
	// Organization roles whose confirmed members get an email and a console
	// notification when the monitor fires and resolves. Empty when the alert
	// events are the only delivery. Not used while `escalationId` is set.
	Notify []string `json:"notify"`
	// ID of the organization's escalation that notifies the team of this
	// monitor's alerts, step by step until one is acknowledged. When set,
	// `notify` is not used. Empty to notify the `notify` roles once.
	EscalationId string `json:"escalationId"`
	// How urgent the monitor's alerts are: `info`, `warning`, `error` or
	// `critical`. Critical alerts may be delivered with higher urgency by the
	// notification channels that support it.
	Severity string `json:"severity"`
	// Whether the monitor is firing, as of its last evaluation.
	Status string `json:"status"`
	// Metric value at the last evaluation. Null before the first evaluation, and
	// when the last evaluation (`evaluatedAt`) found no usage recorded to read:
	// without data a monitor is not compared, so it keeps its status and open
	// alert.
	Value *float64 `json:"value"`
	// Last evaluation date in ISO 8601 format.
	EvaluatedAt *string `json:"evaluatedAt"`
	// ID of the alert the monitor is firing for. Empty while it is not firing.
	AlertId string `json:"alertId"`
	// Date the monitor last announced an alert, in ISO 8601 format.
	TriggeredAt *string `json:"triggeredAt"`
	// Probe monitors: seconds between two checks, read back raised to the plan's
	// minimum. 0 for usage monitors, which are evaluated on the server's
	// interval.
	Interval int `json:"interval"`
	// Probe monitors: seconds a check may take before it fails.
	Timeout int `json:"timeout"`
	// HTTP monitors: URL requested.
	Url string `json:"url"`
	// HTTP monitors: method the URL is requested with. Null for other types.
	Method *string `json:"method"`
	// HTTP monitors: headers sent, as an object of header names and values.
	Headers interface{} `json:"headers"`
	// HTTP monitors: request body sent.
	Body string `json:"body"`
	// DNS monitors: domain queried.
	Domain string `json:"domain"`
	// DNS monitors: record type queried. Null for other types.
	Record *string `json:"record"`
	// DNS monitors: resolver queried. Empty for the server's DNS resolver.
	Resolver string `json:"resolver"`
	// TCP and SMTP monitors: hostname or IP address connected to.
	Host string `json:"host"`
	// TCP and SMTP monitors: port connected to.
	Port int `json:"port"`
	// SMTP monitors: how the connection is secured. `auto` picks implicit TLS on
	// port 465, STARTTLS on 587 and none otherwise. Null for other types.
	TlsMode *string `json:"tlsMode"`
	// Probe monitors: what the last check found. `ok` passes; anything else
	// opened the monitor's alert. Null before the first check.
	Result *string `json:"result"`
	// Probe monitors: milliseconds the last check took.
	Latency *int `json:"latency"`
	// Probe monitors: last status code an HTTP target answered with, reply code
	// an SMTP server sent, or response code a DNS resolver returned.
	Code *int `json:"code"`
	// Probe monitors: expiry date of the TLS certificate the last check saw, in
	// ISO 8601 format. Null without TLS.
	CertificateExpiresAt *string `json:"certificateExpiresAt"`

	// Used by Decode() method
	data []byte
}

func (model Monitor) New(data []byte) *Monitor {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Monitor) UnmarshalJSON(data []byte) error {
	type alias Monitor
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
		Name         json.RawMessage `json:"name"`
		ProjectId    json.RawMessage `json:"projectId"`
		Rules        json.RawMessage `json:"rules"`
		Type         json.RawMessage `json:"type"`
		ResourceType json.RawMessage `json:"resourceType"`
		ResourceId   json.RawMessage `json:"resourceId"`
		Window       json.RawMessage `json:"window"`
		Cooldown     json.RawMessage `json:"cooldown"`
		Enabled      json.RawMessage `json:"enabled"`
		Notify       json.RawMessage `json:"notify"`
		EscalationId json.RawMessage `json:"escalationId"`
		Severity     json.RawMessage `json:"severity"`
		Status       json.RawMessage `json:"status"`
		AlertId      json.RawMessage `json:"alertId"`
		Interval     json.RawMessage `json:"interval"`
		Timeout      json.RawMessage `json:"timeout"`
		Url          json.RawMessage `json:"url"`
		Body         json.RawMessage `json:"body"`
		Domain       json.RawMessage `json:"domain"`
		Resolver     json.RawMessage `json:"resolver"`
		Host         json.RawMessage `json:"host"`
		Port         json.RawMessage `json:"port"`
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
	if present.ProjectId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProjectId); err != nil {
			return err
		}
	}
	if present.Rules == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Rules); err != nil {
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
	if present.Window == nil {
		if err := json.Unmarshal([]byte("3600"), &decoded.Window); err != nil {
			return err
		}
	}
	if present.Cooldown == nil {
		if err := json.Unmarshal([]byte("3600"), &decoded.Cooldown); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Notify == nil {
		if err := json.Unmarshal([]byte("[\"owner\"]"), &decoded.Notify); err != nil {
			return err
		}
	}
	if present.EscalationId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.EscalationId); err != nil {
			return err
		}
	}
	if present.Severity == nil {
		if err := json.Unmarshal([]byte("\"warning\""), &decoded.Severity); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"ok\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.AlertId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AlertId); err != nil {
			return err
		}
	}
	if present.Interval == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Interval); err != nil {
			return err
		}
	}
	if present.Timeout == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Timeout); err != nil {
			return err
		}
	}
	if present.Url == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Url); err != nil {
			return err
		}
	}
	if present.Body == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Body); err != nil {
			return err
		}
	}
	if present.Domain == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Domain); err != nil {
			return err
		}
	}
	if present.Resolver == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Resolver); err != nil {
			return err
		}
	}
	if present.Host == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Host); err != nil {
			return err
		}
	}
	if present.Port == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Port); err != nil {
			return err
		}
	}

	*model = Monitor(decoded)
	return nil
}

func (model *Monitor) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
