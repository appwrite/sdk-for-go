package models

import (
	"encoding/json"
	"errors"
)

// Topic Model
type Topic struct {
	// Topic ID.
	Id string `json:"$id"`
	// Topic creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Topic update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// The name of the topic.
	Name string `json:"name"`
	// Total count of email subscribers subscribed to the topic.
	EmailTotal int `json:"emailTotal"`
	// Total count of SMS subscribers subscribed to the topic.
	SmsTotal int `json:"smsTotal"`
	// Total count of push subscribers subscribed to the topic.
	PushTotal int `json:"pushTotal"`
	// Subscribe permissions.
	Subscribe []string `json:"subscribe"`
	// MQTT QoS for delivery on this topic. Null lets the subscriber choose their
	// level.
	Qos *int `json:"qos"`
	// Message retention in seconds for offline delivery. Null inherits the lowest
	// expiry of the project's enabled Appwrite providers, or the 7-day ledger
	// window if none sets one.
	Expiry *int `json:"expiry"`

	// Used by Decode() method
	data []byte
}

func (model Topic) New(data []byte) *Topic {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Topic) UnmarshalJSON(data []byte) error {
	type alias Topic
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id         json.RawMessage `json:"$id"`
		CreatedAt  json.RawMessage `json:"$createdAt"`
		UpdatedAt  json.RawMessage `json:"$updatedAt"`
		Name       json.RawMessage `json:"name"`
		EmailTotal json.RawMessage `json:"emailTotal"`
		SmsTotal   json.RawMessage `json:"smsTotal"`
		PushTotal  json.RawMessage `json:"pushTotal"`
		Subscribe  json.RawMessage `json:"subscribe"`
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
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.EmailTotal == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.EmailTotal); err != nil {
			return err
		}
	}
	if present.SmsTotal == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.SmsTotal); err != nil {
			return err
		}
	}
	if present.PushTotal == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.PushTotal); err != nil {
			return err
		}
	}
	if present.Subscribe == nil {
		if err := json.Unmarshal([]byte("[\"users\"]"), &decoded.Subscribe); err != nil {
			return err
		}
	}

	*model = Topic(decoded)
	return nil
}

func (model *Topic) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
