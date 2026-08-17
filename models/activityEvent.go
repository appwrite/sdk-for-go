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
