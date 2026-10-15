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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabaseSpecification) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseSpecification
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Slug                 json.RawMessage `json:"slug"`
		Name                 json.RawMessage `json:"name"`
		Price                json.RawMessage `json:"price"`
		StorageOverageRate   json.RawMessage `json:"storageOverageRate"`
		BandwidthOverageRate json.RawMessage `json:"bandwidthOverageRate"`
		ReplicaRate          json.RawMessage `json:"replicaRate"`
		PitrRate             json.RawMessage `json:"pitrRate"`
		Cpu                  json.RawMessage `json:"cpu"`
		Memory               json.RawMessage `json:"memory"`
		MaxConnections       json.RawMessage `json:"maxConnections"`
		IncludedStorage      json.RawMessage `json:"includedStorage"`
		IncludedBandwidth    json.RawMessage `json:"includedBandwidth"`
		Enabled              json.RawMessage `json:"enabled"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Slug == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Slug); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Price == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Price); err != nil {
			return err
		}
	}
	if present.StorageOverageRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.StorageOverageRate); err != nil {
			return err
		}
	}
	if present.BandwidthOverageRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.BandwidthOverageRate); err != nil {
			return err
		}
	}
	if present.ReplicaRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ReplicaRate); err != nil {
			return err
		}
	}
	if present.PitrRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.PitrRate); err != nil {
			return err
		}
	}
	if present.Cpu == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Cpu); err != nil {
			return err
		}
	}
	if present.Memory == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Memory); err != nil {
			return err
		}
	}
	if present.MaxConnections == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.MaxConnections); err != nil {
			return err
		}
	}
	if present.IncludedStorage == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.IncludedStorage); err != nil {
			return err
		}
	}
	if present.IncludedBandwidth == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.IncludedBandwidth); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Enabled); err != nil {
			return err
		}
	}

	*model = DedicatedDatabaseSpecification(decoded)
	return nil
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
