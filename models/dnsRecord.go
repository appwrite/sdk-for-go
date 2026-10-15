package models

import (
	"encoding/json"
	"errors"
)

// DNSRecord Model
type DnsRecord struct {
	// DNS Record ID.
	Id string `json:"$id"`
	// DNS Record creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// DNS Record update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// DNS record type (e.g. A, CNAME, MX).
	Type string `json:"type"`
	// Record name or subdomain.
	Name string `json:"name"`
	// Value of the record (IP address, domain, etc.).
	Value string `json:"value"`
	// Time to live (in seconds).
	Ttl int `json:"ttl"`
	// Record priority (commonly used for MX).
	Priority int `json:"priority"`
	// Whether this record is locked (read-only).
	Lock bool `json:"lock"`
	// Record weight (used for SRV records).
	Weight int `json:"weight"`
	// Target port (used for SRV records).
	Port int `json:"port"`
	// Comment for the DNS record.
	Comment string `json:"comment"`

	// Used by Decode() method
	data []byte
}

func (model DnsRecord) New(data []byte) *DnsRecord {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DnsRecord) UnmarshalJSON(data []byte) error {
	type alias DnsRecord
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
		Value     json.RawMessage `json:"value"`
		Ttl       json.RawMessage `json:"ttl"`
		Priority  json.RawMessage `json:"priority"`
		Lock      json.RawMessage `json:"lock"`
		Weight    json.RawMessage `json:"weight"`
		Port      json.RawMessage `json:"port"`
		Comment   json.RawMessage `json:"comment"`
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
	if present.Value == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Value); err != nil {
			return err
		}
	}
	if present.Ttl == nil {
		if err := json.Unmarshal([]byte("3600"), &decoded.Ttl); err != nil {
			return err
		}
	}
	if present.Priority == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Priority); err != nil {
			return err
		}
	}
	if present.Lock == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Lock); err != nil {
			return err
		}
	}
	if present.Weight == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Weight); err != nil {
			return err
		}
	}
	if present.Port == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Port); err != nil {
			return err
		}
	}
	if present.Comment == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Comment); err != nil {
			return err
		}
	}

	*model = DnsRecord(decoded)
	return nil
}

func (model *DnsRecord) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
