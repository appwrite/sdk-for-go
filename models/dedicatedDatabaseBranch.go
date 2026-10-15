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
	Port *int `json:"port"`
	// Advertised catalog the client connects to. MySQL/MariaDB use default;
	// Postgres uses the routing label.
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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabaseBranch) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseBranch
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		BranchId         json.RawMessage `json:"branchId"`
		BranchName       json.RawMessage `json:"branchName"`
		Namespace        json.RawMessage `json:"namespace"`
		ExpiresAt        json.RawMessage `json:"expiresAt"`
		Host             json.RawMessage `json:"host"`
		Database         json.RawMessage `json:"database"`
		Username         json.RawMessage `json:"username"`
		Password         json.RawMessage `json:"password"`
		Ssl              json.RawMessage `json:"ssl"`
		Engine           json.RawMessage `json:"engine"`
		ConnectionString json.RawMessage `json:"connectionString"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.BranchId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BranchId); err != nil {
			return err
		}
	}
	if present.BranchName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BranchName); err != nil {
			return err
		}
	}
	if present.Namespace == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Namespace); err != nil {
			return err
		}
	}
	if present.ExpiresAt == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ExpiresAt); err != nil {
			return err
		}
	}
	if present.Host == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Host); err != nil {
			return err
		}
	}
	if present.Database == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Database); err != nil {
			return err
		}
	}
	if present.Username == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Username); err != nil {
			return err
		}
	}
	if present.Password == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Password); err != nil {
			return err
		}
	}
	if present.Ssl == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Ssl); err != nil {
			return err
		}
	}
	if present.Engine == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Engine); err != nil {
			return err
		}
	}
	if present.ConnectionString == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ConnectionString); err != nil {
			return err
		}
	}

	*model = DedicatedDatabaseBranch(decoded)
	return nil
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
