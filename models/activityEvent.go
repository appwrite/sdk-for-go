package models

import (
	"encoding/json"
	"errors"
)

// ActivityEvent Model
type ActivityEvent struct {
	// Event ID.
	Id string `json:"$id"`
	// Actor type.
	ActorType string `json:"actorType"`
	// Actor ID.
	ActorId string `json:"actorId"`
	// Actor Email.
	ActorEmail string `json:"actorEmail"`
	// Actor Name.
	ActorName string `json:"actorName"`
	// Resource parent.
	ResourceParent string `json:"resourceParent"`
	// Resource type.
	ResourceType string `json:"resourceType"`
	// Resource ID.
	ResourceId string `json:"resourceId"`
	// Resource.
	Resource string `json:"resource"`
	// Event name.
	Event string `json:"event"`
	// User agent.
	UserAgent string `json:"userAgent"`
	// IP address.
	Ip string `json:"ip"`
	// API mode when event triggered.
	Mode string `json:"mode"`
	// Location.
	Country string `json:"country"`
	// Continent code.
	ContinentCode string `json:"continentCode"`
	// City name.
	City string `json:"city"`
	// Region/state chain.
	Subdivisions string `json:"subdivisions"`
	// Internet service provider.
	Isp string `json:"isp"`
	// Autonomous System Number (ASN).
	AutonomousSystemNumber string `json:"autonomousSystemNumber"`
	// Organization that owns the ASN.
	AutonomousSystemOrganization string `json:"autonomousSystemOrganization"`
	// Connection type (e.g. cable, cellular, corporate).
	ConnectionType string `json:"connectionType"`
	// User type (e.g. residential, business, hosting).
	ConnectionUsageType string `json:"connectionUsageType"`
	// Registered organization of the IP.
	ConnectionOrganization string `json:"connectionOrganization"`
	// Log creation date in ISO 8601 format.
	Time string `json:"time"`
	// Project ID.
	ProjectId string `json:"projectId"`
	// Team ID.
	TeamId string `json:"teamId"`
	// Hostname.
	Hostname string `json:"hostname"`
	// Name of the SDK that triggered the event.
	Sdk string `json:"sdk"`
	// Version of the SDK that triggered the event.
	SdkVersion string `json:"sdkVersion"`

	// Used by Decode() method
	data []byte
}

func (model ActivityEvent) New(data []byte) *ActivityEvent {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *ActivityEvent) UnmarshalJSON(data []byte) error {
	type alias ActivityEvent
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                           json.RawMessage `json:"$id"`
		ActorType                    json.RawMessage `json:"actorType"`
		ActorId                      json.RawMessage `json:"actorId"`
		ActorEmail                   json.RawMessage `json:"actorEmail"`
		ActorName                    json.RawMessage `json:"actorName"`
		ResourceParent               json.RawMessage `json:"resourceParent"`
		ResourceType                 json.RawMessage `json:"resourceType"`
		ResourceId                   json.RawMessage `json:"resourceId"`
		Resource                     json.RawMessage `json:"resource"`
		Event                        json.RawMessage `json:"event"`
		UserAgent                    json.RawMessage `json:"userAgent"`
		Ip                           json.RawMessage `json:"ip"`
		Mode                         json.RawMessage `json:"mode"`
		Country                      json.RawMessage `json:"country"`
		ContinentCode                json.RawMessage `json:"continentCode"`
		City                         json.RawMessage `json:"city"`
		Subdivisions                 json.RawMessage `json:"subdivisions"`
		Isp                          json.RawMessage `json:"isp"`
		AutonomousSystemNumber       json.RawMessage `json:"autonomousSystemNumber"`
		AutonomousSystemOrganization json.RawMessage `json:"autonomousSystemOrganization"`
		ConnectionType               json.RawMessage `json:"connectionType"`
		ConnectionUsageType          json.RawMessage `json:"connectionUsageType"`
		ConnectionOrganization       json.RawMessage `json:"connectionOrganization"`
		Time                         json.RawMessage `json:"time"`
		ProjectId                    json.RawMessage `json:"projectId"`
		TeamId                       json.RawMessage `json:"teamId"`
		Hostname                     json.RawMessage `json:"hostname"`
		Sdk                          json.RawMessage `json:"sdk"`
		SdkVersion                   json.RawMessage `json:"sdkVersion"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.ActorType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ActorType); err != nil {
			return err
		}
	}
	if present.ActorId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ActorId); err != nil {
			return err
		}
	}
	if present.ActorEmail == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ActorEmail); err != nil {
			return err
		}
	}
	if present.ActorName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ActorName); err != nil {
			return err
		}
	}
	if present.ResourceParent == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceParent); err != nil {
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
	if present.Resource == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Resource); err != nil {
			return err
		}
	}
	if present.Event == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Event); err != nil {
			return err
		}
	}
	if present.UserAgent == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserAgent); err != nil {
			return err
		}
	}
	if present.Ip == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Ip); err != nil {
			return err
		}
	}
	if present.Mode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Mode); err != nil {
			return err
		}
	}
	if present.Country == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Country); err != nil {
			return err
		}
	}
	if present.ContinentCode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ContinentCode); err != nil {
			return err
		}
	}
	if present.City == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.City); err != nil {
			return err
		}
	}
	if present.Subdivisions == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Subdivisions); err != nil {
			return err
		}
	}
	if present.Isp == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Isp); err != nil {
			return err
		}
	}
	if present.AutonomousSystemNumber == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AutonomousSystemNumber); err != nil {
			return err
		}
	}
	if present.AutonomousSystemOrganization == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AutonomousSystemOrganization); err != nil {
			return err
		}
	}
	if present.ConnectionType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ConnectionType); err != nil {
			return err
		}
	}
	if present.ConnectionUsageType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ConnectionUsageType); err != nil {
			return err
		}
	}
	if present.ConnectionOrganization == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ConnectionOrganization); err != nil {
			return err
		}
	}
	if present.Time == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Time); err != nil {
			return err
		}
	}
	if present.ProjectId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProjectId); err != nil {
			return err
		}
	}
	if present.TeamId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TeamId); err != nil {
			return err
		}
	}
	if present.Hostname == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Hostname); err != nil {
			return err
		}
	}
	if present.Sdk == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Sdk); err != nil {
			return err
		}
	}
	if present.SdkVersion == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SdkVersion); err != nil {
			return err
		}
	}

	*model = ActivityEvent(decoded)
	return nil
}

func (model *ActivityEvent) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
