package models

import (
	"encoding/json"
	"errors"
)

// DNSRecordsList Model
type DnsRecordsList struct {
	// Total number of dnsRecords that matched your query.
	Total int `json:"total"`
	// List of dnsRecords.
	DnsRecords []DnsRecord `json:"dnsRecords"`

	// Used by Decode() method
	data []byte
}

func (model DnsRecordsList) New(data []byte) *DnsRecordsList {
	model.data = data
	return &model
}

func (model *DnsRecordsList) UnmarshalJSON(data []byte) error {
	type alias DnsRecordsList
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"total\":0}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
}

func (model *DnsRecordsList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
