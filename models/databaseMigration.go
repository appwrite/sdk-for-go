package models

import (
	"encoding/json"
	"errors"
)

// DatabaseMigration Model
type DatabaseMigration struct {
	// Database migration ID.
	Id string `json:"$id"`
	// Migration creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Migration update time in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Project ID that owns the migrating database.
	ProjectId string `json:"projectId"`
	// Logical database ID being migrated.
	DatabaseId string `json:"databaseId"`
	// Dedicated compute specification provisioned for the migration target.
	Specification string `json:"specification"`
	// Migration phase. Possible values: pending, provisioned, capturing,
	// backfilling, catching_up, verifying, ready_to_cutover, cutover, soaking,
	// done, failed, rolled_back.
	Phase string `json:"phase"`
	// Number of times a migration step has failed and been recorded.
	Attempt int `json:"attempt"`
	// Reason the most recent migration step failed, empty while none has.
	LastError string `json:"lastError"`
	// Number of documents still pending replication to the target.
	LagDocuments int `json:"lagDocuments"`
	// Highest source changelog sequence applied to the target so far.
	ChangelogWatermark int `json:"changelogWatermark"`
	// Time the migrated data was verified against the source in ISO 8601 format.
	VerifiedAt string `json:"verifiedAt"`
	// Time routing was flipped to the target in ISO 8601 format.
	CutoverAt string `json:"cutoverAt"`
	// Time the post-cutover soak window ends in ISO 8601 format.
	SoakUntil string `json:"soakUntil"`
	// Whether the migration cuts over automatically once ready. Set when the
	// migration is created and never changed afterwards, so it always reports
	// what was asked for.
	AutoCutover bool `json:"autoCutover"`
	// Whether a cutover has been requested and not yet attempted. Set by the
	// cutover endpoint and cleared when the attempt is made, so a cutover that
	// fails a check parks the migration again rather than retrying on its own.
	CutoverRequested bool `json:"cutoverRequested"`
	// Whether the migration is paused.
	Paused bool `json:"paused"`

	// Used by Decode() method
	data []byte
}

func (model DatabaseMigration) New(data []byte) *DatabaseMigration {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DatabaseMigration) UnmarshalJSON(data []byte) error {
	type alias DatabaseMigration
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                 json.RawMessage `json:"$id"`
		CreatedAt          json.RawMessage `json:"$createdAt"`
		UpdatedAt          json.RawMessage `json:"$updatedAt"`
		ProjectId          json.RawMessage `json:"projectId"`
		DatabaseId         json.RawMessage `json:"databaseId"`
		Specification      json.RawMessage `json:"specification"`
		Phase              json.RawMessage `json:"phase"`
		Attempt            json.RawMessage `json:"attempt"`
		LastError          json.RawMessage `json:"lastError"`
		LagDocuments       json.RawMessage `json:"lagDocuments"`
		ChangelogWatermark json.RawMessage `json:"changelogWatermark"`
		VerifiedAt         json.RawMessage `json:"verifiedAt"`
		CutoverAt          json.RawMessage `json:"cutoverAt"`
		SoakUntil          json.RawMessage `json:"soakUntil"`
		AutoCutover        json.RawMessage `json:"autoCutover"`
		CutoverRequested   json.RawMessage `json:"cutoverRequested"`
		Paused             json.RawMessage `json:"paused"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.ProjectId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProjectId); err != nil {
			return err
		}
	}
	if present.DatabaseId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DatabaseId); err != nil {
			return err
		}
	}
	if present.Specification == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Specification); err != nil {
			return err
		}
	}
	if present.Phase == nil {
		if err := json.Unmarshal([]byte("\"pending\""), &decoded.Phase); err != nil {
			return err
		}
	}
	if present.Attempt == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Attempt); err != nil {
			return err
		}
	}
	if present.LastError == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.LastError); err != nil {
			return err
		}
	}
	if present.LagDocuments == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.LagDocuments); err != nil {
			return err
		}
	}
	if present.ChangelogWatermark == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ChangelogWatermark); err != nil {
			return err
		}
	}
	if present.VerifiedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.VerifiedAt); err != nil {
			return err
		}
	}
	if present.CutoverAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CutoverAt); err != nil {
			return err
		}
	}
	if present.SoakUntil == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SoakUntil); err != nil {
			return err
		}
	}
	if present.AutoCutover == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.AutoCutover); err != nil {
			return err
		}
	}
	if present.CutoverRequested == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.CutoverRequested); err != nil {
			return err
		}
	}
	if present.Paused == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Paused); err != nil {
			return err
		}
	}

	*model = DatabaseMigration(decoded)
	return nil
}

func (model *DatabaseMigration) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
