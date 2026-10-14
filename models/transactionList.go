package models

import (
	"encoding/json"
	"errors"
)

// TransactionList Model
type TransactionList struct {
	// Total number of transactions that matched your query.
	Total int `json:"total"`
	// List of transactions.
	Transactions []Transaction `json:"transactions"`

	// Used by Decode() method
	data []byte
}

func (model TransactionList) New(data []byte) *TransactionList {
	model.data = data
	return &model
}

func (model *TransactionList) UnmarshalJSON(data []byte) error {
	type alias TransactionList
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

func (model *TransactionList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
