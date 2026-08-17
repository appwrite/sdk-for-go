package models

import (
	"encoding/json"
	"errors"
)

// Branch Model
type DedicatedDatabaseBranch struct {
	// Branch identifier.
	BranchId string `json:"branchId"`
	// Branch name.
	BranchName string `json:"branchName"`
	// Kubernetes namespace where the branch is deployed.
	Namespace string `json:"namespace"`
	// Unix timestamp when the branch expires.
	ExpiresAt int `json:"expiresAt"`
	// Branch hostname for direct connections.
	Host string `json:"host"`
	// Branch port. Null until the backing reports one.
	Port int `json:"port"`
	// Database name the client sends for routing to the branch.
	Database string `json:"database"`
	// Database username. Shared with the parent database.
	Username string `json:"username"`
	// Database password. Shared with the parent database.
	Password string `json:"password"`
	// Whether SSL is required.
	Ssl bool `json:"ssl"`
	// Database engine. Possible values: postgresql, mysql, mongodb.
	Engine string `json:"engine"`
	// Full connection string for the branch.
	ConnectionString string `json:"connectionString"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseBranch) New(data []byte) *DedicatedDatabaseBranch {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseBranch) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
