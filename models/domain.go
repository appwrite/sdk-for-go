package models

import (
	"encoding/json"
	"errors"
)

// Domain Model
type Domain struct {
	// Domain ID.
	Id string `json:"$id"`
	// Domain creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Domain update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Domain name.
	Domain string `json:"domain"`
	// Domain registrar (e.g. "appwrite" or "third_party").
	Registrar string `json:"registrar"`
	// Nameservers setting. "Appwrite" or empty string.
	Nameservers string `json:"nameservers"`
	// Domain expiry date in ISO 8601 format.
	Expire string `json:"expire"`
	// Domain renewal date in ISO 8601 format.
	Renewal string `json:"renewal"`
	// If set to true, the domain will automatically renew.
	AutoRenewal bool `json:"autoRenewal"`
	// Renewal price (in cents).
	RenewalPrice int `json:"renewalPrice"`
	// Transfer status for domains being transferred in. Null when the domain is
	// not being transferred.
	TransferStatus string `json:"transferStatus"`
	// Team ID.
	TeamId string `json:"teamId"`
	// Dns records
	DnsRecords []DnsRecord `json:"dnsRecords"`

	// Used by Decode() method
	data []byte
}

func (model Domain) New(data []byte) *Domain {
	model.data = data
	return &model
}

func (model *Domain) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
