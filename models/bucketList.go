package models

import (
	"encoding/json"
	"errors"
)

// BucketsList Model
type BucketList struct {
	// Total number of buckets that matched your query.
	Total int `json:"total"`
	// List of buckets.
	Buckets []Bucket `json:"buckets"`

	// Used by Decode() method
	data []byte
}

func (model BucketList) New(data []byte) *BucketList {
	model.data = data
	return &model
}

func (model *BucketList) UnmarshalJSON(data []byte) error {
	type alias BucketList
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

func (model *BucketList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
