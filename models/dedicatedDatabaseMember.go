package models

import (
	"encoding/json"
	"errors"
)

// Member Model
type DedicatedDatabaseMember struct {
	// Member identifier.
	Id string `json:"$id"`
	// Member role. Possible values: primary (accepts reads and writes), replica
	// (read-only follower), unknown (placement not established; reported while a
	// transition is moving or restarting the topology and this member has not
	// been probed, so no member can be named the write target).
	Role string `json:"role"`
	// Member pod status. Possible values: pending (configured but absent from the
	// backend topology, so nothing is bringing it up), provisioning (pod missing
	// or Pending), starting (Running but not Ready), active (Running and Ready),
	// failed (Failed phase or CrashLoopBackOff container), or the lowercased pod
	// phase reported by the cluster.
	Status string `json:"status"`
	// Whether the member is streaming from the primary. True when the engine
	// reports the replication link up, false when it reports the link down, and
	// null when no reading was taken: a primary has no stream to report, an
	// inactive member is not probed, and neither is any member while no primary
	// is established.
	Replicating bool `json:"replicating"`
	// Replication lag in seconds. Null when the lag is not known: a primary has
	// none to report, and a member the backend has not probed has none yet. Also
	// null against `replicating: true`, for a member that is streaming but whose
	// engine printed no numeric lag.
	LagSeconds float64 `json:"lagSeconds"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseMember) New(data []byte) *DedicatedDatabaseMember {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseMember) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
