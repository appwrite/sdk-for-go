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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabasePooler) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabasePooler
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Enabled             json.RawMessage `json:"enabled"`
		Mode                json.RawMessage `json:"mode"`
		MaxConnections      json.RawMessage `json:"maxConnections"`
		DefaultPoolSize     json.RawMessage `json:"defaultPoolSize"`
		Port                json.RawMessage `json:"port"`
		ReadWriteSplitting  json.RawMessage `json:"readWriteSplitting"`
		PoolerCpuRequest    json.RawMessage `json:"poolerCpuRequest"`
		PoolerCpuLimit      json.RawMessage `json:"poolerCpuLimit"`
		PoolerMemoryRequest json.RawMessage `json:"poolerMemoryRequest"`
		PoolerMemoryLimit   json.RawMessage `json:"poolerMemoryLimit"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Mode == nil {
		if err := json.Unmarshal([]byte("\"transaction\""), &decoded.Mode); err != nil {
			return err
		}
	}
	if present.MaxConnections == nil {
		if err := json.Unmarshal([]byte("100"), &decoded.MaxConnections); err != nil {
			return err
		}
	}
	if present.DefaultPoolSize == nil {
		if err := json.Unmarshal([]byte("20"), &decoded.DefaultPoolSize); err != nil {
			return err
		}
	}
	if present.Port == nil {
		if err := json.Unmarshal([]byte("6432"), &decoded.Port); err != nil {
			return err
		}
	}
	if present.ReadWriteSplitting == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.ReadWriteSplitting); err != nil {
			return err
		}
	}
	if present.PoolerCpuRequest == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PoolerCpuRequest); err != nil {
			return err
		}
	}
	if present.PoolerCpuLimit == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PoolerCpuLimit); err != nil {
			return err
		}
	}
	if present.PoolerMemoryRequest == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PoolerMemoryRequest); err != nil {
			return err
		}
	}
	if present.PoolerMemoryLimit == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PoolerMemoryLimit); err != nil {
			return err
		}
	}

	*model = DedicatedDatabasePooler(decoded)
	return nil
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
