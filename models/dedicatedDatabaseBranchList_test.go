package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseBranchListModel(t *testing.T) {
	model := DedicatedDatabaseBranchList{Branches: []DedicatedDatabaseBranch{DedicatedDatabaseBranch{BranchId: "a1b2c3d4", BranchName: "branch-a1b2c3d4", Namespace: "db-myproject-mydb-branch-a1b2c3d4", ExpiresAt: 1711411200, Host: "db-myproject-mydb-a1b2c3d4.fra.appwrite.center", Port: 5432, Database: "db-myproject-mydb-a1b2c3d4", Username: "appwrite", Password: "********", Ssl: true, Engine: "postgresql", ConnectionString: "postgresql://appwrite:****@db-myproject-mydb-a1b2c3d4.fra.appwrite.center:5432/db-myproject-mydb-a1b2c3d4?sslmode=disable"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseBranchList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
}
