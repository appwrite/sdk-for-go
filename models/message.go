package models

import (
	"encoding/json"
	"errors"
)

// Message Model
type Message struct {
	// Message ID.
	Id string `json:"$id"`
	// Message creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Message update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Message provider type.
	ProviderType string `json:"providerType"`
	// Topic IDs set as recipients.
	Topics []string `json:"topics"`
	// User IDs set as recipients.
	Users []string `json:"users"`
	// Target IDs set as recipients.
	Targets []string `json:"targets"`
	// The scheduled time for message.
	ScheduledAt *string `json:"scheduledAt"`
	// The time when the message was delivered.
	DeliveredAt *string `json:"deliveredAt"`
	// Delivery errors if any.
	DeliveryErrors []string `json:"deliveryErrors"`
	// Number of recipients the message was delivered to.
	DeliveredTotal int `json:"deliveredTotal"`
	// Data of the message.
	Data interface{} `json:"data"`
	// Status of delivery.
	Status string `json:"status"`

	// Used by Decode() method
	data []byte
}

func (model Message) New(data []byte) *Message {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Message) UnmarshalJSON(data []byte) error {
	type alias Message
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id             json.RawMessage `json:"$id"`
		CreatedAt      json.RawMessage `json:"$createdAt"`
		UpdatedAt      json.RawMessage `json:"$updatedAt"`
		ProviderType   json.RawMessage `json:"providerType"`
		DeliveredTotal json.RawMessage `json:"deliveredTotal"`
		Data           json.RawMessage `json:"data"`
		Status         json.RawMessage `json:"status"`
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
	if present.ProviderType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderType); err != nil {
			return err
		}
	}
	if present.DeliveredTotal == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.DeliveredTotal); err != nil {
			return err
		}
	}
	if present.Data == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Data); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"draft\""), &decoded.Status); err != nil {
			return err
		}
	}

	*model = Message(decoded)
	return nil
}

func (model *Message) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
