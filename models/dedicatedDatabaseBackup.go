package models

import (
	"encoding/json"
	"errors"
)

// Backup Model
type DedicatedDatabaseBackup struct {
	// Backup ID.
	Id string `json:"$id"`
	// Backup creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Database ID this backup belongs to.
	DatabaseId string `json:"databaseId"`
	// Project ID.
	ProjectId string `json:"projectId"`
	// Backup policy ID when the backup was created by a schedule.
	PolicyId string `json:"policyId"`
	// Backup trigger. Possible values: manual, schedule.
	Trigger string `json:"trigger"`
	// Backup type. Possible values: full (complete database snapshot),
	// incremental (changes since last backup), wal (write-ahead log continuous
	// archival).
	Type string `json:"type"`
	// Backup type that was requested. Differs from `type` when the backend could
	// not run the requested type and took a different one instead, in which case
	// `fallbackReason` explains why. Empty for backups taken before the requested
	// type was recorded.
	RequestedType string `json:"requestedType"`
	// Why the backend ran a different backup type than the one requested. Empty
	// when the backup ran as requested.
	FallbackReason string `json:"fallbackReason"`
	// Backup status. Possible values: pending (queued for processing), running
	// (currently in progress), completed (successfully finished), failed
	// (encountered an error), verified (integrity check passed).
	Status string `json:"status"`
	// Backup size in bytes.
	SizeBytes int `json:"sizeBytes"`
	// Backup start time in ISO 8601 format.
	StartedAt string `json:"startedAt"`
	// Backup completion time in ISO 8601 format.
	CompletedAt string `json:"completedAt"`
	// Backup verification time in ISO 8601 format.
	VerifiedAt string `json:"verifiedAt"`
	// Backup expiration time in ISO 8601 format.
	ExpiresAt string `json:"expiresAt"`
	// Transaction-log position the backup anchors at, in the engine's own
	// notation: PostgreSQL `{walSegment}|{lsn}`, MySQL and MariaDB
	// `{binlogFile}|{offset}`, MongoDB `{seconds}|{increment}`. Empty when the
	// backup recorded no position, which is the case for backup types that carry
	// none.
	LogPosition string `json:"logPosition"`
	// Error message if backup failed.
	Error string `json:"error"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseBackup) New(data []byte) *DedicatedDatabaseBackup {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseBackup) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
