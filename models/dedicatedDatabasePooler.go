package models

import (
	"encoding/json"
	"errors"
)

// PoolerConfig Model
type DedicatedDatabasePooler struct {
	// Whether connection pooling is enabled.
	Enabled bool `json:"enabled"`
	// Connection pool mode. Possible values: transaction (releases connections
	// back to pool after each transaction), session (holds connections for the
	// entire client session).
	Mode string `json:"mode"`
	// Client-connection ceiling the pooler accepts. Enforced on MySQL and
	// MariaDB; on PostgreSQL the pooler has no client cap, so this reports the
	// database's advertised networkMaxConnections and cannot be set here.
	MaxConnections int `json:"maxConnections"`
	// Default pool size per user.
	DefaultPoolSize int `json:"defaultPoolSize"`
	// Pooler listening port.
	Port int `json:"port"`
	// Whether SELECTs are routed to HA replicas while writes and locked reads
	// stay on the primary. Active only when HA is enabled.
	ReadWriteSplitting bool `json:"readWriteSplitting"`
	// Effective CPU request applied to the pooler sidecar container (Kubernetes
	// quantity). Returns the proportional default (5% of DB CPU, floor 100m)
	// unless overridden.
	PoolerCpuRequest string `json:"poolerCpuRequest"`
	// Effective CPU limit applied to the pooler sidecar container (Kubernetes
	// quantity). Returns the proportional default (10% of DB CPU, floor 200m)
	// unless overridden.
	PoolerCpuLimit string `json:"poolerCpuLimit"`
	// Effective memory request applied to the pooler sidecar container
	// (Kubernetes quantity). Returns the proportional default (7.5% of DB memory,
	// floor 64Mi) unless overridden.
	PoolerMemoryRequest string `json:"poolerMemoryRequest"`
	// Effective memory limit applied to the pooler sidecar container (Kubernetes
	// quantity). Returns the proportional default (15% of DB memory, floor 128Mi)
	// unless overridden.
	PoolerMemoryLimit string `json:"poolerMemoryLimit"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabasePooler) New(data []byte) *DedicatedDatabasePooler {
	model.data = data
	return &model
}

func (model *DedicatedDatabasePooler) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
