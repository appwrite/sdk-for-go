package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseBranchModel(t *testing.T) {
	model := DedicatedDatabaseBranch{BranchId: "a1b2c3d4", BranchName: "branch-a1b2c3d4", Namespace: "db-myproject-mydb-branch-a1b2c3d4", ExpiresAt: 1711411200, Host: "db-myproject-mydb-a1b2c3d4.fra.appwrite.center", Port: 5432, Database: "db-myproject-mydb-a1b2c3d4", Username: "appwrite", Password: "********", Ssl: true, Engine: "postgresql", ConnectionString: "postgresql://appwrite:****@db-myproject-mydb-a1b2c3d4.fra.appwrite.center:5432/db-myproject-mydb-a1b2c3d4?sslmode=disable"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseBranch
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.BranchId != model.BranchId {
		t.Errorf("Expected BranchId %v, got %v", model.BranchId, result.BranchId)
	}
	if result.BranchName != model.BranchName {
		t.Errorf("Expected BranchName %v, got %v", model.BranchName, result.BranchName)
	}
	if result.Namespace != model.Namespace {
		t.Errorf("Expected Namespace %v, got %v", model.Namespace, result.Namespace)
	}
	if result.ExpiresAt != model.ExpiresAt {
		t.Errorf("Expected ExpiresAt %v, got %v", model.ExpiresAt, result.ExpiresAt)
	}
	if result.Host != model.Host {
		t.Errorf("Expected Host %v, got %v", model.Host, result.Host)
	}
	if result.Port != model.Port {
		t.Errorf("Expected Port %v, got %v", model.Port, result.Port)
	}
	if result.Database != model.Database {
		t.Errorf("Expected Database %v, got %v", model.Database, result.Database)
	}
	if result.Username != model.Username {
		t.Errorf("Expected Username %v, got %v", model.Username, result.Username)
	}
	if result.Password != model.Password {
		t.Errorf("Expected Password %v, got %v", model.Password, result.Password)
	}
	if result.Ssl != model.Ssl {
		t.Errorf("Expected Ssl %v, got %v", model.Ssl, result.Ssl)
	}
	if result.Engine != model.Engine {
		t.Errorf("Expected Engine %v, got %v", model.Engine, result.Engine)
	}
	if result.ConnectionString != model.ConnectionString {
		t.Errorf("Expected ConnectionString %v, got %v", model.ConnectionString, result.ConnectionString)
	}
}
