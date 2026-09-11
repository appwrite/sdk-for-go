package models

import (
	"encoding/json"
	"errors"
)

// Specification Model
type DedicatedDatabaseSpecification struct {
	// Specification slug. Use this value when creating a dedicated database.
	Slug string `json:"slug"`
	// Human readable specification name.
	Name string `json:"name"`
	// Monthly price of the specification in USD.
	Price float64 `json:"price"`
	// Price per GB of storage above the included amount, per month, in USD.
	StorageOverageRate float64 `json:"storageOverageRate"`
	// Price per GB of bandwidth above the included amount, per month, in USD.
	BandwidthOverageRate float64 `json:"bandwidthOverageRate"`
	// High availability replica price as a fraction of the specification price.
	ReplicaRate float64 `json:"replicaRate"`
	// Point-in-time recovery price as a fraction of the specification price.
	PitrRate float64 `json:"pitrRate"`
	// Allocated CPU in millicores.
	Cpu int `json:"cpu"`
	// Allocated memory in MB.
	Memory int `json:"memory"`
	// Maximum number of concurrent connections.
	MaxConnections int `json:"maxConnections"`
	// Included storage in GB before overage charges apply.
	IncludedStorage int `json:"includedStorage"`
	// Included bandwidth in GB before overage charges apply.
	IncludedBandwidth int `json:"includedBandwidth"`
	// Whether the specification is available on the current plan.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseSpecification) New(data []byte) *DedicatedDatabaseSpecification {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseSpecification) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
