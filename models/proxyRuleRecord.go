package models

import (
	"encoding/json"
	"errors"
)

// RuleRecord Model
type ProxyRuleRecord struct {
	// Record check ID.
	Id string `json:"$id"`
	// Record check creation date in ISO 8601 format. The date the record type was
	// first checked for the rule.
	CreatedAt string `json:"$createdAt"`
	// Record check update date in ISO 8601 format. The date DNS was last checked
	// for this record.
	UpdatedAt string `json:"$updatedAt"`
	// DNS record type. Possible values are "CNAME", "A", "AAAA", or "CAA".
	Type string `json:"type"`
	// Hostname the record was looked up for. A CAA record that blocks the
	// certificate names the parent it was found on.
	Name string `json:"name"`
	// Value the record is expected to have.
	Expected string `json:"expected"`
	// Values DNS returned for the record type when last checked, at most 10.
	// Empty when none were found.
	Current []string `json:"current"`
	// Check status. Possible values are "verified" when the expected value was
	// found, "missing" when no record of the type exists yet, "mismatch" when
	// records exist but none has the expected value, "misplaced" when no record
	// exists but one was found at the domain followed by its zone, as DNS
	// providers do when the full domain is entered as the host, or "failed" when
	// the lookup itself failed.
	Status string `json:"status"`
	// Why the check did not pass. Empty when verified.
	Message string `json:"message"`
	// Whether the record must be verified on its own. Records that are not
	// required are alternatives: the domain passes when any one of them is
	// verified.
	Required bool `json:"required"`

	// Used by Decode() method
	data []byte
}

func (model ProxyRuleRecord) New(data []byte) *ProxyRuleRecord {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *ProxyRuleRecord) UnmarshalJSON(data []byte) error {
	type alias ProxyRuleRecord
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UpdatedAt json.RawMessage `json:"$updatedAt"`
		Type      json.RawMessage `json:"type"`
		Name      json.RawMessage `json:"name"`
		Expected  json.RawMessage `json:"expected"`
		Current   json.RawMessage `json:"current"`
		Status    json.RawMessage `json:"status"`
		Message   json.RawMessage `json:"message"`
		Required  json.RawMessage `json:"required"`
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
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Expected == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expected); err != nil {
			return err
		}
	}
	if present.Current == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Current); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Message == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Message); err != nil {
			return err
		}
	}
	if present.Required == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Required); err != nil {
			return err
		}
	}

	*model = ProxyRuleRecord(decoded)
	return nil
}

func (model *ProxyRuleRecord) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
