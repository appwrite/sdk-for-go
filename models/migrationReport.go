package models

import (
	"encoding/json"
	"errors"
)

// MigrationReport Model
type MigrationReport struct {
	// Number of users to be migrated.
	User int `json:"user"`
	// Number of teams to be migrated.
	Team int `json:"team"`
	// Number of databases to be migrated.
	Database int `json:"database"`
	// Number of rows to be migrated.
	Row int `json:"row"`
	// Number of files to be migrated.
	File int `json:"file"`
	// Number of buckets to be migrated.
	Bucket int `json:"bucket"`
	// Number of functions to be migrated.
	Function int `json:"function"`
	// Number of platforms to be migrated.
	Platform int `json:"platform"`
	// Number of API keys to be migrated.
	ApiKey int `json:"api-key"`
	// Number of project variables to be migrated.
	ProjectVariable int `json:"project-variable"`
	// Number of webhooks to be migrated.
	Webhook int `json:"webhook"`
	// Number of auth-method configs to be migrated (always 0 or 1 — the
	// project-level flag bundle).
	AuthMethods int `json:"auth-methods"`
	// Number of protocol configs to be migrated (always 0 or 1 — the
	// project-level REST/GraphQL/WebSocket flags).
	ProjectProtocols int `json:"project-protocols"`
	// Number of label sets to be migrated (always 0 or 1 — the project-level
	// RBAC label array).
	ProjectLabels int `json:"project-labels"`
	// Number of service configs to be migrated (always 0 or 1 — the
	// project-level enable/disable flags for all 17 services).
	ProjectServices int `json:"project-services"`
	// Number of policy bundles to be migrated (always 0 or 1 — the
	// project-level security policies covering password rules, session behavior,
	// user limits, and membership privacy).
	Policies int `json:"policies"`
	// Number of SMTP configurations to be migrated (always 0 or 1 — the
	// project-level custom SMTP settings; password is not exposed by the source
	// API).
	Smtp int `json:"smtp"`
	// Number of custom-domain proxy rules to be migrated. Auto-generated
	// `.appwrite.network` rules are skipped — they are recreated by parent
	// Function/Site migration.
	Rule int `json:"rule"`
	// Number of custom email templates to be migrated (one per templateId ×
	// locale pair).
	ProjectEmailTemplate int `json:"project-email-template"`
	// Number of project environments to be migrated. Each is recreated under a
	// new ID, empty: its own resources belong to a separate project.
	ProjectEnvironment int `json:"project-environment"`
	// Number of sites to be migrated.
	Site int `json:"site"`
	// Number of providers to be migrated.
	Provider int `json:"provider"`
	// Number of topics to be migrated.
	Topic int `json:"topic"`
	// Number of subscribers to be migrated.
	Subscriber int `json:"subscriber"`
	// Number of messages to be migrated.
	Message int `json:"message"`
	// Size of files to be migrated in mb.
	Size int `json:"size"`
	// Version of the Appwrite instance to be migrated.
	Version string `json:"version"`
	// Number of OAuth2 provider configurations to be migrated. Secrets
	// (clientSecret, p8File) are never migrated — destination admin must
	// re-enter them per provider.
	Oauth2Provider int `json:"oauth2-provider"`
	// Number of backup policies to be migrated.
	BackupPolicy int `json:"backup-policy"`

	// Used by Decode() method
	data []byte
}

func (model MigrationReport) New(data []byte) *MigrationReport {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *MigrationReport) UnmarshalJSON(data []byte) error {
	type alias MigrationReport
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		User                 json.RawMessage `json:"user"`
		Team                 json.RawMessage `json:"team"`
		Database             json.RawMessage `json:"database"`
		Row                  json.RawMessage `json:"row"`
		File                 json.RawMessage `json:"file"`
		Bucket               json.RawMessage `json:"bucket"`
		Function             json.RawMessage `json:"function"`
		Platform             json.RawMessage `json:"platform"`
		ApiKey               json.RawMessage `json:"api-key"`
		ProjectVariable      json.RawMessage `json:"project-variable"`
		Webhook              json.RawMessage `json:"webhook"`
		AuthMethods          json.RawMessage `json:"auth-methods"`
		ProjectProtocols     json.RawMessage `json:"project-protocols"`
		ProjectLabels        json.RawMessage `json:"project-labels"`
		ProjectServices      json.RawMessage `json:"project-services"`
		Policies             json.RawMessage `json:"policies"`
		Smtp                 json.RawMessage `json:"smtp"`
		Rule                 json.RawMessage `json:"rule"`
		ProjectEmailTemplate json.RawMessage `json:"project-email-template"`
		ProjectEnvironment   json.RawMessage `json:"project-environment"`
		Site                 json.RawMessage `json:"site"`
		Provider             json.RawMessage `json:"provider"`
		Topic                json.RawMessage `json:"topic"`
		Subscriber           json.RawMessage `json:"subscriber"`
		Message              json.RawMessage `json:"message"`
		Size                 json.RawMessage `json:"size"`
		Version              json.RawMessage `json:"version"`
		Oauth2Provider       json.RawMessage `json:"oauth2-provider"`
		BackupPolicy         json.RawMessage `json:"backup-policy"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.User == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.User); err != nil {
			return err
		}
	}
	if present.Team == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Team); err != nil {
			return err
		}
	}
	if present.Database == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Database); err != nil {
			return err
		}
	}
	if present.Row == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Row); err != nil {
			return err
		}
	}
	if present.File == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.File); err != nil {
			return err
		}
	}
	if present.Bucket == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Bucket); err != nil {
			return err
		}
	}
	if present.Function == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Function); err != nil {
			return err
		}
	}
	if present.Platform == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Platform); err != nil {
			return err
		}
	}
	if present.ApiKey == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ApiKey); err != nil {
			return err
		}
	}
	if present.ProjectVariable == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ProjectVariable); err != nil {
			return err
		}
	}
	if present.Webhook == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Webhook); err != nil {
			return err
		}
	}
	if present.AuthMethods == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.AuthMethods); err != nil {
			return err
		}
	}
	if present.ProjectProtocols == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ProjectProtocols); err != nil {
			return err
		}
	}
	if present.ProjectLabels == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ProjectLabels); err != nil {
			return err
		}
	}
	if present.ProjectServices == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ProjectServices); err != nil {
			return err
		}
	}
	if present.Policies == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Policies); err != nil {
			return err
		}
	}
	if present.Smtp == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Smtp); err != nil {
			return err
		}
	}
	if present.Rule == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Rule); err != nil {
			return err
		}
	}
	if present.ProjectEmailTemplate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ProjectEmailTemplate); err != nil {
			return err
		}
	}
	if present.ProjectEnvironment == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ProjectEnvironment); err != nil {
			return err
		}
	}
	if present.Site == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Site); err != nil {
			return err
		}
	}
	if present.Provider == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Provider); err != nil {
			return err
		}
	}
	if present.Topic == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Topic); err != nil {
			return err
		}
	}
	if present.Subscriber == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Subscriber); err != nil {
			return err
		}
	}
	if present.Message == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Message); err != nil {
			return err
		}
	}
	if present.Size == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Size); err != nil {
			return err
		}
	}
	if present.Version == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Version); err != nil {
			return err
		}
	}
	if present.Oauth2Provider == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Oauth2Provider); err != nil {
			return err
		}
	}
	if present.BackupPolicy == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.BackupPolicy); err != nil {
			return err
		}
	}

	*model = MigrationReport(decoded)
	return nil
}

func (model *MigrationReport) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
