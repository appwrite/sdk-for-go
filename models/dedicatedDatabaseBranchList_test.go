package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabaseBranchListModel(t *testing.T) {
	model := DedicatedDatabaseBranchList{Total: 2, Branches: []DedicatedDatabaseBranch{DedicatedDatabaseBranch{BranchId: "a1b2c3d4", BranchName: "branch-a1b2c3d4", Namespace: "db-myproject-mydb-branch-a1b2c3d4", ExpiresAt: 1711411200, Host: "db-myproject-mydb-a1b2c3d4.fra.appwrite.center", Database: "db-myproject-mydb-a1b2c3d4", Username: "appwrite", Password: "********", Ssl: true, Engine: "postgresql", ConnectionString: "postgresql://appwrite:****@db-myproject-mydb-a1b2c3d4.fra.appwrite.center:5432/db-myproject-mydb-a1b2c3d4?sslmode=require"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseBranchList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestDedicatedDatabaseBranchListModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseBranchList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := DedicatedDatabaseBranchList{Total: 2, Branches: []DedicatedDatabaseBranch{DedicatedDatabaseBranch{BranchId: "a1b2c3d4", BranchName: "branch-a1b2c3d4", Namespace: "db-myproject-mydb-branch-a1b2c3d4", ExpiresAt: 1711411200, Host: "db-myproject-mydb-a1b2c3d4.fra.appwrite.center", Database: "db-myproject-mydb-a1b2c3d4", Username: "appwrite", Password: "********", Ssl: true, Engine: "postgresql", ConnectionString: "postgresql://appwrite:****@db-myproject-mydb-a1b2c3d4.fra.appwrite.center:5432/db-myproject-mydb-a1b2c3d4?sslmode=require"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
