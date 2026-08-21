package models

import (
	"encoding/json"
	"errors"
)

// Replica Model
type DatabaseStatusReplica struct {
	// Member index within the database. Read `role` for which member accepts
	// writes: a failover moves the primary without renumbering the indexes.
	Index int `json:"index"`
	// Member role. Possible values: primary (accepts reads and writes), replica
	// (read-only follower), unknown (placement not established; reported while a
	// transition is moving or restarting the topology, so no member can be named
	// the write target).
	Role string `json:"role"`
	// Whether the replica is healthy.
	Healthy bool `json:"healthy"`
	// Whether the engine reports this member's replication stream as up. Null
	// when no reading was taken: a primary has no stream to report, and a member
	// that is not healthy, or whose probe did not answer, has none yet. `healthy`
	// is a reachability probe of the member itself and says nothing about
	// replication, so a healthy member may still not be replicating.
	Replicating bool `json:"replicating"`
	// Replication lag in seconds (null for primary). Also null against
	// `replicating: true`, for a member that is streaming but whose engine
	// printed no numeric lag.
	LagSeconds float64 `json:"lagSeconds"`

	// Used by Decode() method
	data []byte
}

func (model DatabaseStatusReplica) New(data []byte) *DatabaseStatusReplica {
	model.data = data
	return &model
}

func (model *DatabaseStatusReplica) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
