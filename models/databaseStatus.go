package models

import (
	"encoding/json"
	"errors"
)

// Status Model
type DatabaseStatus struct {
	// Overall health status: healthy, degraded, unhealthy, or unknown when
	// nothing could be measured.
	Health string `json:"health"`
	// Whether the database is ready to accept connections.
	Ready bool `json:"ready"`
	// Database engine: postgresql, mysql, or mongodb.
	Engine string `json:"engine"`
	// Database engine version.
	Version string `json:"version"`
	// Database uptime in seconds.
	Uptime int `json:"uptime"`
	// Connection statistics. Null when the database has not been probed.
	Connections DatabaseStatusConnections `json:"connections"`
	// Requested replication sync mode. Possible values: async, sync, quorum.
	// Compare with effectiveSyncMode for what the primary is enforcing.
	SyncMode string `json:"syncMode"`
	// Replication sync mode the primary is actually enforcing. Null when high
	// availability is disabled or the state could not be read.
	EffectiveSyncMode *string `json:"effectiveSyncMode"`
	// Whether the enforced replication is weaker than the requested syncMode.
	SyncDegraded bool `json:"syncDegraded"`
	// Number of standby acknowledgements the primary waits for before a write is
	// committed.
	SyncAcknowledgements int `json:"syncAcknowledgements"`
	// Number of standbys registered with the primary for synchronous replication.
	SyncStandbyCount int `json:"syncStandbyCount"`
	// Whether the other sync fields are an engine reading rather than a recorded
	// estimate. True when the primary answered what it is enforcing, including
	// when that answer contradicted the record, in which case the contradicted
	// values are replaced by the ones the engine reports. False when the reading
	// could not be taken: the probe did not answer, there was no engine to ask,
	// or the values describe a configuration change just applied rather than
	// anything measured. Absent when no engine was asked at all, so an unprobed
	// database is distinguishable from an unconfirmed one. False never means a
	// standby was found lagging, because it is the absence of a reading rather
	// than a negative one, so draw no conclusion about replication health from it
	// or from a response that omits it.
	SyncStateConfirmed *bool `json:"syncStateConfirmed"`
	// List of database replicas and their status. Every configured member
	// appears, including one the backend has not brought up, which is reported as
	// not healthy.
	Replicas []DatabaseStatusReplica `json:"replicas"`
	// Storage volume information.
	Volumes []DatabaseStatusVolume `json:"volumes"`

	// Used by Decode() method
	data []byte
}

func (model DatabaseStatus) New(data []byte) *DatabaseStatus {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DatabaseStatus) UnmarshalJSON(data []byte) error {
	type alias DatabaseStatus
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Health               json.RawMessage `json:"health"`
		Ready                json.RawMessage `json:"ready"`
		Engine               json.RawMessage `json:"engine"`
		Version              json.RawMessage `json:"version"`
		Uptime               json.RawMessage `json:"uptime"`
		SyncMode             json.RawMessage `json:"syncMode"`
		SyncDegraded         json.RawMessage `json:"syncDegraded"`
		SyncAcknowledgements json.RawMessage `json:"syncAcknowledgements"`
		SyncStandbyCount     json.RawMessage `json:"syncStandbyCount"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Health == nil {
		if err := json.Unmarshal([]byte("\"unknown\""), &decoded.Health); err != nil {
			return err
		}
	}
	if present.Ready == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Ready); err != nil {
			return err
		}
	}
	if present.Engine == nil {
		if err := json.Unmarshal([]byte("\"postgresql\""), &decoded.Engine); err != nil {
			return err
		}
	}
	if present.Version == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Version); err != nil {
			return err
		}
	}
	if present.Uptime == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Uptime); err != nil {
			return err
		}
	}
	if present.SyncMode == nil {
		if err := json.Unmarshal([]byte("\"async\""), &decoded.SyncMode); err != nil {
			return err
		}
	}
	if present.SyncDegraded == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.SyncDegraded); err != nil {
			return err
		}
	}
	if present.SyncAcknowledgements == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.SyncAcknowledgements); err != nil {
			return err
		}
	}
	if present.SyncStandbyCount == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.SyncStandbyCount); err != nil {
			return err
		}
	}

	*model = DatabaseStatus(decoded)
	return nil
}

func (model *DatabaseStatus) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
