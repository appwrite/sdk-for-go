package models

import (
	"encoding/json"
	"errors"
)

// Subscriber Model
type Subscriber struct {
	// Subscriber ID.
	Id string `json:"$id"`
	// Subscriber creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Subscriber update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Target ID.
	TargetId string `json:"targetId"`
	// Target.
	Target Target `json:"target"`
	// Topic ID.
	UserId string `json:"userId"`
	// User Name.
	UserName string `json:"userName"`
	// Topic ID.
	TopicId string `json:"topicId"`
	// The target provider type. Can be one of the following: `email`, `sms` or
	// `push`.
	ProviderType string `json:"providerType"`

	// Used by Decode() method
	data []byte
}

func (model Subscriber) New(data []byte) *Subscriber {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Subscriber) UnmarshalJSON(data []byte) error {
	type alias Subscriber
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
		TargetId     json.RawMessage `json:"targetId"`
		UserId       json.RawMessage `json:"userId"`
		UserName     json.RawMessage `json:"userName"`
		TopicId      json.RawMessage `json:"topicId"`
		ProviderType json.RawMessage `json:"providerType"`
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
	if present.TargetId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TargetId); err != nil {
			return err
		}
	}
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.UserName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserName); err != nil {
			return err
		}
	}
	if present.TopicId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TopicId); err != nil {
			return err
		}
	}
	if present.ProviderType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderType); err != nil {
			return err
		}
	}

	*model = Subscriber(decoded)
	return nil
}

func (model *Subscriber) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
