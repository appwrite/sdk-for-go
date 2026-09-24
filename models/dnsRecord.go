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
