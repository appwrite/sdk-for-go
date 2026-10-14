package models

import (
	"encoding/json"
	"errors"
)

// ColumnURL Model
type ColumnUrl struct {
	// Column Key.
	Key string `json:"key"`
	// Column type.
	Type string `json:"type"`
	// Column status. Possible values: `available`, `processing`, `deleting`,
	// `stuck`, or `failed`
	Status string `json:"status"`
	// Error message. Displays error generated on failure of creating or deleting
	// an column.
	Error string `json:"error"`
	// Is column required?
	Required bool `json:"required"`
	// Is column an array?
	Array *bool `json:"array"`
	// Column creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Column update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// String format.
	Format string `json:"format"`
	// Default value for column when not provided. Cannot be set when column is
	// required.
	Default *string `json:"default"`

	// Used by Decode() method
	data []byte
}

func (model ColumnUrl) New(data []byte) *ColumnUrl {
	model.data = data
	return &model
}

func (model *ColumnUrl) UnmarshalJSON(data []byte) error {
	type alias ColumnUrl
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"key\":\"\",\"type\":\"\",\"error\":\"\",\"required\":false,\"$createdAt\":\"\",\"$updatedAt\":\"\",\"format\":\"url\"}"), &fields); err != nil {
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

func (model *ColumnUrl) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
