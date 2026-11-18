package models

import (
	"encoding/json"
	"errors"
)

// AnalyticsAnnotation Model
type AnalyticsAnnotation struct {
	// Annotation ID.
	Id string `json:"$id"`
	// Annotation creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Annotation update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// ID of the analytics property the annotation belongs to.
	PropertyId string `json:"propertyId"`
	// Point in time the annotation marks, in ISO 8601 format.
	Date string `json:"date"`
	// Annotation text.
	Content string `json:"content"`
	// ID of the user who created the annotation. Empty when it was created with
	// an API key.
	UserId string `json:"userId"`

	// Used by Decode() method
	data []byte
}

func (model AnalyticsAnnotation) New(data []byte) *AnalyticsAnnotation {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AnalyticsAnnotation) UnmarshalJSON(data []byte) error {
	type alias AnalyticsAnnotation
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id         json.RawMessage `json:"$id"`
		CreatedAt  json.RawMessage `json:"$createdAt"`
		UpdatedAt  json.RawMessage `json:"$updatedAt"`
		PropertyId json.RawMessage `json:"propertyId"`
		Date       json.RawMessage `json:"date"`
		Content    json.RawMessage `json:"content"`
		UserId     json.RawMessage `json:"userId"`
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
	if present.PropertyId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PropertyId); err != nil {
			return err
		}
	}
	if present.Date == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Date); err != nil {
			return err
		}
	}
	if present.Content == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Content); err != nil {
			return err
		}
	}
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}

	*model = AnalyticsAnnotation(decoded)
	return nil
}

func (model *AnalyticsAnnotation) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
