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
	TransferStatus *string `json:"transferStatus"`
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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Domain) UnmarshalJSON(data []byte) error {
	type alias Domain
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
		Domain       json.RawMessage `json:"domain"`
		Registrar    json.RawMessage `json:"registrar"`
		Nameservers  json.RawMessage `json:"nameservers"`
		Expire       json.RawMessage `json:"expire"`
		Renewal      json.RawMessage `json:"renewal"`
		AutoRenewal  json.RawMessage `json:"autoRenewal"`
		RenewalPrice json.RawMessage `json:"renewalPrice"`
		TeamId       json.RawMessage `json:"teamId"`
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
	if present.Domain == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Domain); err != nil {
			return err
		}
	}
	if present.Registrar == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Registrar); err != nil {
			return err
		}
	}
	if present.Nameservers == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Nameservers); err != nil {
			return err
		}
	}
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}
	if present.Renewal == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Renewal); err != nil {
			return err
		}
	}
	if present.AutoRenewal == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.AutoRenewal); err != nil {
			return err
		}
	}
	if present.RenewalPrice == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.RenewalPrice); err != nil {
			return err
		}
	}
	if present.TeamId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TeamId); err != nil {
			return err
		}
	}

	*model = Domain(decoded)
	return nil
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
